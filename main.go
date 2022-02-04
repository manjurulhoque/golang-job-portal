package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/manjurulhoque/golang-job-portal/config"
	"github.com/manjurulhoque/golang-job-portal/models"
	"github.com/manjurulhoque/golang-job-portal/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	logrus.SetReportCaller(true)

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Job{})

	r := routes.SetupRouter()
	// running
	r.Run()
}
