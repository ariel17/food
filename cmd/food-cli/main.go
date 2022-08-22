package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/clients/email"
	"github.com/ariel17/food/internal/repositories"
	"github.com/ariel17/food/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", configs.GetDatabaseConfig().String())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	platesFlag := flag.Bool("plates", false, "Shows plates available.")
	planFlag := flag.Bool("plan", false, "Creates a new plan.")
	emailFlag := flag.Bool("email", false, "Enables email sender.")
	helpFlag := flag.Bool("help", false, "Shows help.")
	configFlag := flag.Bool("config", false, "Shows actual configuration.")
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	printer := services.NewPrinter()
	if *configFlag{
		printer.PrintConfiguration(os.Stdout)
		return
	}

	repository := repositories.NewRepositoryMySQL(db)

	if *platesFlag {
		plates, err := repository.GetAllPlates()
		if err != nil {
			panic(err)
		}
		printer.PrintPlates(os.Stdout, plates)
		return
	}

	if *planFlag {
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

		if *emailFlag {
			emailSender := email.NewSender(printer)
			if err := emailSender.Send(plan, items); err != nil {
				panic(err)
			}
			fmt.Println("Email sent.")
		}
	}
}