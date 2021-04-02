package email

import (
	"bytes"
	"fmt"
	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/entities"
	"github.com/ariel17/food/internal/services"
	"net/smtp"
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
	buffer := bytes.NewBufferString("")
	s.printer.PrintPlan(buffer, plan)
	_, err := fmt.Fprintln(buffer, "")
	if err != nil {
		return err
	}
	s.printer.PrintShopList(buffer, items)
	return smtp.SendMail(config.String(), config.Auth, config.Account, config.Recipients, buffer.Bytes())
}