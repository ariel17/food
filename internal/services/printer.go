package services

import (
	"fmt"
	"github.com/ariel17/food/internal/entities"
	"strings"
)

const (
	tabSize = 8
	kilo = 1000
)

type Printer interface {
	PrintPlates(plates []entities.Plate)
	PrintPlan(plan []entities.Plate)
	PrintShopList(items []entities.Step)
}

func NewPrinter() Printer {
	return &printer{}
}

type printer struct {
}

func (p *printer) PrintPlates(plates []entities.Plate) {
	fmt.Println("")
	fmt.Println("Name")
	fmt.Println("----")
	for _, plate := range plates {
		fmt.Println(fmt.Sprintf("* %s", plate.Name))
	}
}

func (p *printer) PrintPlan(plan []entities.Plate) {
	fmt.Println("")
	fmt.Println("Day\t\tName")
	fmt.Println("---\t\t----")
	for day, plate := range plan {
		fmt.Println(fmt.Sprintf("%d\t\t%s", day+1, plate.Name))
	}
}

func (p *printer) PrintShopList(items []entities.Step) {
	fmt.Println("")
	fmt.Println("Item\t\t\tAmount\t\tUnit")
	fmt.Println("----\t\t\t------\t\t----")
	for _, item := range items {
		tabs := generateTabs(item.Ingredient.Name, 3)
		amount, unit := formatAmountAndUnit(item.Amount, item.Unit)
		fmt.Println(fmt.Sprintf("%s"+tabs+"%.2f\t\t%s", item.Ingredient.Name, amount, unit))
	}
}

func generateTabs(s string, expected int) string {
	consumed := len(s) / tabSize
	return strings.Repeat("\t", expected-consumed)
}

func formatAmountAndUnit(amount float64, unit string) (float64, string) {
	newAmount := amount / kilo
	if newAmount >= 1 {
		return newAmount, "K"+unit
	}
	return amount, unit
}
