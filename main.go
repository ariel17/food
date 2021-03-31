package main

import (
	"database/sql"
	"flag"
	"github.com/ariel17/food/configs"
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

	platesFlag := flag.Bool("plates", false, "Shows plates availables.")
	planFlag := flag.Bool("plan", false, "Creates a new plan.")
	flag.Parse()

	repository := repositories.NewRepositoryMySQL(db)
	planner := services.NewPlannerService(repository)
	printer := services.NewPrinter()

	if *platesFlag {
		plates, err := repository.GetAllPlates()
		if err != nil {
			panic(err)
		}
		printer.PrintPlates(plates)

	} else if *planFlag {
		plan, err := planner.CreatePlan()
		if err != nil {
			panic(err)
		}
		printer.PrintPlan(plan)

		items, err := planner.CreateShopList(plan)
		if err != nil {
			panic(err)
		}
		printer.PrintShopList(items)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
