package email

import (
	"bytes"
	"fmt"
	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/entities"
	"github.com/ariel17/food/internal/services"
	"html/template"
	"net/smtp"
	"os"
)

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
	message := s.render(s.content(plan, items))
	return smtp.SendMail(config.String(), config.Auth, config.Account, config.Recipients, message)
}

func (s *sender) content(plan []entities.Plate, items []entities.Step) string {
	buffer := bytes.NewBufferString("")
	s.printer.PrintPlan(buffer, plan)
	_, _ = fmt.Fprintln(buffer, "")
	s.printer.PrintShopList(buffer, items)
	return buffer.String()
}

func (s *sender) render(content string) []byte {
	workingDir, _ := os.Getwd()
	template := template.Must(template.ParseFiles(workingDir + "/template.html"))
	buffer := bytes.NewBufferString("")
	template.Execute(buffer, content)
	return buffer.Bytes()
}