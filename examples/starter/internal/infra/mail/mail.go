package mail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/9ssi7/gopre-starter/assets"
	"github.com/9ssi7/gopre-starter/config"
	smtp_mail "github.com/xhit/go-simple-mail/v2"
)

type SendConfig struct {
	To      []string
	Subject string
	Message string
}

type SendWithTemplateConfig struct {
	SendConfig
	Template string
	Data     any
}

type srv struct {
	cnf    config.Smtp
	server *smtp_mail.SMTPServer
}

var client *srv

func Init() {
	cnf := config.ReadValue().Adapters.Smtp
	server := smtp_mail.NewSMTPClient()
	server.Host = cnf.Host
	server.Port = cnf.Port
	server.Username = cnf.Sender
	server.Password = cnf.Password
	server.Encryption = smtp_mail.EncryptionSTARTTLS
	server.Authentication = smtp_mail.AuthLogin
	client = &srv{
		server: server,
		cnf:    cnf,
	}
}

func GetClient() *srv {
	if client == nil {
		Init()
	}
	return client
}

func GetField(str string) string {
	if str == "" {
		return "N/A"
	}
	return str
}

func (s *srv) createClient() (*smtp_mail.SMTPClient, error) {
	return s.server.Connect()
}

func (s *srv) SendText(cnf SendConfig) error {
	client, err := s.createClient()
	if err != nil {
		fmt.Println("Error creating client: ", err)
		return err
	}
	email := smtp_mail.NewMSG()
	email.SetFrom(s.cnf.From)
	email.AddTo(cnf.To...)
	email.SetSubject(cnf.Subject)
	email.SetSender(s.cnf.Sender)
	email.SetReplyTo(s.cnf.Reply)
	email.AddAlternative(smtp_mail.TextPlain, cnf.Message)
	err = email.Send(client)
	if err != nil {
		fmt.Println("Error sending email: ", err)
		return err
	}
	return nil
}

func (s *srv) SendWithTemplate(cnf SendWithTemplateConfig) error {
	client, err := s.createClient()
	if err != nil {
		fmt.Println("Error creating client: ", err)
		return err
	}
	dir := assets.EmbedMailTemplate()
	t := template.Must(template.ParseFS(dir, fmt.Sprintf("mail/%s.html", cnf.Template)))
	var tpl bytes.Buffer
	t.Execute(&tpl, cnf.Data)
	body := tpl.String()
	email := smtp_mail.NewMSG()
	email.SetFrom(s.cnf.From)
	email.AddTo(cnf.To...)
	email.SetSubject(cnf.Subject)
	email.SetSender(s.cnf.Sender)
	email.SetReplyTo(s.cnf.Reply)
	email.SetBody(smtp_mail.TextHTML, body)
	if err = email.Send(client); err != nil {
		fmt.Println("Error sending email: ", err)
		return err
	}
	return nil
}
