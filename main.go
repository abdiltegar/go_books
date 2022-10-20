package main

import (
	"learn_orm/config"
	"learn_orm/database"
	"learn_orm/routes"
)

func main() {
	database.InitDB()
	config.InitConfig()

	e := routes.New(database.DB)
	e.Logger.Fatal(e.Start(config.Cfg.API_PORT))
}
