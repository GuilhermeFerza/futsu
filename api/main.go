package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joho/godotenv"
)

func enviarEmailConfirmacao(destinatario string) {

	remetente := os.Getenv("REMETENTE")
	senha := os.Getenv("SENHA")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	mensagem := []byte(
		"Subject: Subscription Confirmed!\r\n" +
			"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
			"\r\n" +
			"Hey!\n\nYour registration has been successfully confirmed. Stay tuned for updates on upcoming releases!",
	)

	auth := smtp.PlainAuth("", remetente, senha, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, remetente, []string{destinatario}, mensagem)

	if err != nil {
		fmt.Println("Erro ao enviar e-mail para", destinatario, ":", err)
		return
	}
	fmt.Println("E-mail enviado com sucesso para", destinatario)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	r.POST("/api/subscribers", func(c *gin.Context) {
		var json struct {
			Email string `json:"email" binding:"required,email"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": "Email invalido"})
			return
		}

		url := os.Getenv("DB_CONN_ACS")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err := pgx.Connect(ctx, url)
		if err != nil {
			fmt.Printf("ERRO DETALHADO: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha na conexão com banco"})
			return
		}
		defer conn.Close(ctx)

		_, err = conn.Exec(ctx, "insert into subscribers (email) values ($1)", json.Email)

		if err != nil {

			if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
				c.JSON(http.StatusConflict, gin.H{"error": "Este e-mail já está inscrito!"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco"})
			return
		}

		go enviarEmailConfirmacao(json.Email)

		c.JSON(http.StatusOK, gin.H{"message": "Inscrito com sucesso!"})

	})

	r.Run("0.0.0.0:8081")
}
