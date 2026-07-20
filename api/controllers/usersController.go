package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GuilhermeFerza/futsu/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func PostUsers(c *gin.Context) {
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

	go services.EnviarEmailConfirmacao(json.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Inscrito com sucesso!"})

}
