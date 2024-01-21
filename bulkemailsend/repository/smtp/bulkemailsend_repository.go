package smtp

import (
	"fmt"
	"net/smtp"

	"github.com/reinhardjs/sayakaya/domain"
)

type bulkEmailSendRepository struct {
	Username string
	Password string
	Host     string
	Port     string
	Sender   string
}

func NewMysqlUserRepository(username string, password string, host string, port string, sender string) domain.BulkEmailSendRepository {
	return &bulkEmailSendRepository{}
}

func (r *bulkEmailSendRepository) BulkSend(bulkEmailSend *domain.BulkEmailSend) (err error) {
	username := r.Username
	password := r.Password
	host := r.Host
	port := r.Port

	// Subject and body
	subject := bulkEmailSend.Subject
	body := bulkEmailSend.Message

	// Sender and receiver
	from := r.Sender
	to := bulkEmailSend.Recipients

	// Build the message
	message := fmt.Sprintf("From: %s\r\n", from)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	// Authentication.
	auth := smtp.PlainAuth("", username, password, host)

	// Send email
	err = smtp.SendMail(host+":"+port, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
