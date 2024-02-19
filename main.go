package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	userRoutes "github.com/vishalsharma/api/routes/usersRoutes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}


	r := userRoutes.UserRoutes()
	color.Blue("SERVER IS RUNNING PORT :: "+os.Getenv(os.Getenv("RUN_MODE")+"_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv(os.Getenv("RUN_MODE")+"_PORT"), r))
}
