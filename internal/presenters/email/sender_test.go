package email

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/food/internal/entities"
	"github.com/ariel17/food/internal/services"
)

func TestSender_Render(t *testing.T) {
	printer := services.NewPrinter()
	s := sender{
		printer: printer,
	}
	b, _ := s.render("holis")
	html := string(b)
	expected := strings.Join([]string{
		"Subject: This week's food plan",
		"MIME-Version: 1.0",
		"Content-Type: text/html",
		"Content-Disposition: inline",
		"",
		"<html>",
		"  <body>",
		"    <pre style=\"font: monospace\">",
		"      holis",
		"    </pre>",
		"  </body>",
		"</html>",
	}, "\n")
	assert.Equal(t, expected, html)
}

func TestSender_Content(t *testing.T) {
	printer := services.NewPrinter()
	s := sender{
		printer: printer,
	}
	plan := []entities.Plate{{Name: "milanesa"}}
	items := []entities.Step{{Ingredient: entities.Ingredient{Name: "carne"}, Amount: 1000, Unit: "g"}}
	content := s.content(plan, items)
	expected := strings.Join([]string{
		"",
		"Day\t\tName",
		"---\t\t----",
		"1\t\tmilanesa",
		"",
		"",
		"Item\t\t\tAmount\t\tUnit",
		"----\t\t\t------\t\t----",
		"carne\t\t\t1.00\t\tKg",
		"",
	}, "\n")
	assert.Equal(t, expected, content)
}