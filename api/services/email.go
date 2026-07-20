package services

import (
	"fmt"
	"net/smtp"
	"os"
)

func EnviarEmailConfirmacao(destinatario string) {

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

func EnviarEmailLancamento(destinatario, nomeMusica, link string) {
	remetente := os.Getenv("REMETENTE")
	senha := os.Getenv("SENHA")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	mensagem := []byte(
		"Subject: New Song is OUT!\r\n" +
			"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
			"\r\n" +
			"Sup!\n\n" +
			"Just released a new music called \"" + nomeMusica + "\".\n\n" +
			"Hurry to listen it: " + link + "\n\n" +
			"Thanks for keep up with my work!",
	)

	auth := smtp.PlainAuth("", remetente, senha, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, remetente, []string{destinatario}, mensagem)

	if err != nil {
		fmt.Println("Erro ao enviar musica para", destinatario, ":", err)
		return
	}
	fmt.Println("Novidade enviada com sucesso para", destinatario)
}
