package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/entities"
	"github.com/ariel17/food/internal/services"
)

// Sender takes the food plan, items involved and sends it by email to
type Sender interface {
	Send(plan []entities.Plate, items []entities.Step) error
}

func NewSender(printer services.Printer) Sender {
	return &sender{
		printer: printer,
	}
}

type sender struct {
	printer services.Printer
}

func (s *sender) Send(plan []entities.Plate, items []entities.Step) error {
	config := configs.GetEmailConfig()
	message, err := s.render(s.content(plan, items))
	if err != nil {
		return err
	}
	return smtp.SendMail(config.String(), config.Auth, config.Account, config.Recipients, message)
}

func (s *sender) content(plan []entities.Plate, items []entities.Step) string {
	buffer := bytes.NewBufferString("")
	s.printer.PrintPlan(buffer, plan)
	_, _ = fmt.Fprintln(buffer, "")
	s.printer.PrintShopList(buffer, items)
	return buffer.String()
}

func (s *sender) render(content string) ([]byte, error) {
	workingDir, _ := os.Getwd()
	template := template.Must(template.ParseFiles(workingDir + "/template.html"))
	buffer := bytes.NewBufferString("")
	err := template.Execute(buffer, content)
	return buffer.Bytes(), err
}
