package services

import (
	"fmt"
	"io"
	"strings"

	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/entities"
)

const (
	tabSize = 8
	kilo    = 1000
)

type Printer interface {
	PrintPlates(w io.Writer, plates []entities.Plate)
	PrintPlan(w io.Writer, plan []entities.Plate)
	PrintShopList(w io.Writer, items []entities.Step)
	PrintConfiguration(w io.Writer, flags configs.Flag)
}

func NewPrinter() Printer {
	return &printer{}
}

type printer struct {
}

func (p *printer) PrintPlates(w io.Writer, plates []entities.Plate) {
	_, _ = fmt.Fprintln(w, "")
	_, _ = fmt.Fprintln(w, "Name")
	_, _ = fmt.Fprintln(w, "----")
	for _, plate := range plates {
		_, _ = fmt.Fprintf(w, "* %s\n", plate.Name)
	}
}

func (p *printer) PrintPlan(w io.Writer, plan []entities.Plate) {
	_, _ = fmt.Fprintln(w, "")
	_, _ = fmt.Fprintln(w, "Day\t\tName")
	_, _ = fmt.Fprintln(w, "---\t\t----")
	for day, plate := range plan {
		_, _ = fmt.Fprintf(w, "%d\t\t%s\n", day+1, plate.Name)
	}
}

func (p *printer) PrintShopList(w io.Writer, items []entities.Step) {
	_, _ = fmt.Fprintln(w, "")
	_, _ = fmt.Fprintln(w, "Item\t\t\tAmount\t\tUnit")
	_, _ = fmt.Fprintln(w, "----\t\t\t------\t\t----")
	for _, item := range items {
		tabs := generateTabs(item.Ingredient.Name, 3)
		amount, unit := formatAmountAndUnit(item.Amount, item.Unit)
		_, _ = fmt.Fprintf(w, "%s"+tabs+"%.2f\t\t%s\n", item.Ingredient.Name, amount, unit)
	}
}

func (p *printer) PrintConfiguration(w io.Writer, flags configs.Flag) {
	_, _ = fmt.Fprintln(w, "")
	_, _ = fmt.Fprintln(w, "Database connection string:", configs.GetDatabaseConfig().String())
	ec := configs.GetEmailConfig()
	_, _ = fmt.Fprintln(w, "Email host:", ec.String())
	_, _ = fmt.Fprintln(w, "Email recipients:", ec.Recipients)
	_, _ = fmt.Fprintln(w, "Recipe source:", flags.Source)
}

func generateTabs(s string, expected int) string {
	consumed := len(s) / tabSize
	return strings.Repeat("\t", expected-consumed)
}

func formatAmountAndUnit(amount float64, unit string) (float64, string) {
	newAmount := amount / kilo
	if newAmount >= 1 {
		return newAmount, "K" + unit
	}
	return amount, unit
}