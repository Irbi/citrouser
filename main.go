package main

import (
	"fmt"
	"github.com/Irbi/citrouser/api"
	"github.com/Irbi/citrouser/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func main() {
	fmt.Println("####################### YEP, IT WORKS #######################")

	a := App{}
	a.Initialize()
	a.Run()
}

func (a *App) Initialize() {
	initLogger()
	initDb(a)
	initRouter(a)
}

func initLogger() {
	log.SetLevel(log.DebugLevel)

	formatter := &log.TextFormatter{
		TimestampFormat: "02-03-2021 10:39:40",
		FullTimestamp:   true,
	}
	formatter.ForceColors = true
	log.SetFormatter(formatter)
}

func initDb(a *App) {
	log.Info("Init DB...")

	a.DB = db.Init(os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"))

	log.Info("DB connected")
}

func initRouter(a *App) {
	log.Info("Init API Routing...")
	a.Router = gin.Default()
	api.New(a.Router)
	log.Info("Routing is running")
}

func (a *App) Run() {
	log.Info("Run Server...")

	err := a.Router.Run(":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Server is running")
}