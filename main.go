package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/presenters/email"
	"github.com/ariel17/food/internal/repositories"
	"github.com/ariel17/food/internal/services"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", configs.GetDatabaseConfig().String())
	if err != nil {
		panic(err)
	}

	platesFlag := flag.Bool("plates", false, "Shows plates available.")
	planFlag := flag.Bool("plan", false, "Creates a new plan.")
	helpFlag := flag.Bool("help", false, "Shows help.")
	flag.Parse()

	repository := repositories.NewRepositoryMySQL(db)
	planner := services.NewPlannerService(repository)
	printer := services.NewPrinter()
	emailSender := email.NewSender(printer)

	if *platesFlag {
		plates, err := repository.GetAllPlates()
		if err != nil {
			panic(err)
		}
		printer.PrintPlates(os.Stdout, plates)
	}

	if *planFlag {
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

		if err := emailSender.Send(plan, items); err != nil {
			panic(err)
		}
		fmt.Println("Email sent.")
	}

	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
