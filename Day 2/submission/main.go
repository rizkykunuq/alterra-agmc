package main

import (
	"fmt"
	"log"
	"os"

	"submission/day2/config"
	"submission/day2/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server " + os.Getenv("APP_NAME") + " running in port " + os.Getenv("PORT"))
	config.InitDB()
	e := routes.AllRoutes()
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
