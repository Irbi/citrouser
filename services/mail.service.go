package services

var Mail *mailService

func init() {
	Mail = &mailService{}
}

type mailService struct {
}

type MailContext struct {
	Content string
	Footer  string
}

func (m *mailService) SendEmail(title, email, templConst string, data map[string]string) error {
	return nil
}

