package main

import (
	"mhub/app/config"
	"mhub/app/database"
	"mhub/routes"
	"os"
)

func main() {
	// Inisialisasi konfigurasi
	cfg := config.InitConfig()
	db := database.InitDBMysql(cfg)
	database.InitMigrationMysql(db)
	e := routes.InitmyRoutes()

	// Mulai server
	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVERPORT")))
}
