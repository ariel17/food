package main

import (
	"fmt"
	"os"

	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/clients/email"
	"github.com/ariel17/food/internal/repositories"
	"github.com/ariel17/food/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	flags := configs.Flag{}
	flags.Parse()

	if flags.Help {
		flags.ShowHelp()
		return
	}

	printer := services.NewPrinter()
	if flags.ShowConfig {
		printer.PrintConfiguration(os.Stdout, flags)
		return
	}

	repository := repositories.New(flags.Source)
	defer repository.Close()

	if flags.ShowPlates {
		plates, err := repository.GetAllPlates()
		if err != nil {
			panic(err)
		}
		printer.PrintPlates(os.Stdout, plates)
		return
	}

	if flags.CreatePlan {
		planner := services.NewPlannerService(repository)
		plan, err := planner.CreatePlan()
		if err != nil {
			panic(err)
		}
		printer.PrintPlan(os.Stdout, plan)

		items, err := planner.CreateShopList(plan)
		if err != nil {
			panic(err)
		}
		printer.PrintShopList(os.Stdout, items)

		if flags.EnableEmail {
			emailSender := email.NewSender(printer)
			if err := emailSender.Send(plan, items); err != nil {
				panic(err)
			}
			fmt.Println("Email sent.")
		}
	}
}