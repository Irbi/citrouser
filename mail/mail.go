package mail

import (
	"github.com/Irbi/citrouser/services"
	log "github.com/sirupsen/logrus"
	"os"
)

var Mail *mailService

type mailService struct {
	webUrl string
}

func (m *mailService) ResetPassword(name, email, password string) error {

	log.Debug("ResetPassword email, code: ", email, password)

	return services.Mail.SendEmail("Reset password", email, os.Getenv("APP_FROM_EMAIL_ADDRESS"),
		map[string]string{"action": "Reset password",
			"password": password,
			"name":     name,
		})
}
