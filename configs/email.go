package configs

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type EmailConfig struct {
	Host       string
	Port       int
	Account    string
	Recipients []string
	Auth       smtp.Auth
}

func (e EmailConfig) String() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

var emailConfig EmailConfig

func GetEmailConfig() EmailConfig {
	return emailConfig
}

func init() {
	emailConfig = EmailConfig{}
	host := os.Getenv("EMAIL_HOST")
	account := os.Getenv("EMAIL_ACCOUNT")
	password := os.Getenv("EMAIL_PASS")

	emailConfig.Host = host
	emailConfig.Port = getInt("EMAIL_PORT")
	emailConfig.Account = account
	emailConfig.Recipients = strings.Split(os.Getenv("EMAIL_RECIPIENTS"), ",")
	emailConfig.Auth = smtp.PlainAuth("", account, password, host)

}
