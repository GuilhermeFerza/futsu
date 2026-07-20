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
)

func PostEmail(c *gin.Context) {
	var json struct {
		NomeMusica string `json:"nome_musica" binding:"required"`
		Link       string `json:"link" binding:"required"`
		SenhaAdmin string `json:"senha_admin" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Preencha todos os dados"})
		return
	}

	if json.SenhaAdmin != os.Getenv("SENHA_ADM") {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "nao autorizado"})
		return
	}

	url := os.Getenv("DB_CONN_ACS")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha no banco"})
		return
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, "SELECT email from subscribers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar inscritos"})
		return
	}

	defer rows.Close()

	contador := 0
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err == nil {
			go services.EnviarEmailLancamento(email, json.NomeMusica, json.Link)
			contador++
		}
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": fmt.Sprintf("Iniciando envio para %d pessoas!", contador)})

}
