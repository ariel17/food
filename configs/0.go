package configs

import (
	"flag"
	"os"
)

type Flag struct {
	ShowPlates  bool
	CreatePlan  bool
	EnableEmail bool
	Help        bool
	ShowConfig  bool
	Source      string
}

func (f *Flag) Parse() {
	flag.BoolVar(&f.ShowPlates, "show-plates", false, "Shows available plates.")
	flag.BoolVar(&f.CreatePlan, "create-plan", false, "Creates a new plan.")
	flag.BoolVar(&f.EnableEmail, "enable-email", false, "Enables email sender.")
	flag.BoolVar(&f.Help, "help", false, "Shows help.")
	flag.BoolVar(&f.ShowConfig, "show-config", false, "Shows actual configuration.")
	flag.StringVar(&f.Source, "source", "yaml", "Source of plates: database, yaml.")
	flag.Parse()

	if f.Source != "yaml" && f.Source != "database" {
		panic("Invalid source: use yaml or database.")
	}
}

func (f *Flag) ShowHelp() {
	flag.PrintDefaults()
}

var environment string

// IsProduction returns true the current environment is production.
func IsProduction() bool {
	return environment == "production"
}

func init() {
	environment = os.Getenv("ENVIRONMENT")
}