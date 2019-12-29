package main

import (
	"remote-action/app/utils"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	utils.InitLogger()
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Error(err)
	}
	if os.Getenv("ENV") == "development" {
		utils.DebugLogging()
	}

	e := echo.New()
	Router(e)
	utils.Logger.Fatal(e.Start(":3000"))
}
