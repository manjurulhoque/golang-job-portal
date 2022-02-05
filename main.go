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
	logrus.SetReportCaller(true) // to show filename and line number

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}(config.DB)

	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Job{})
	config.DB.AutoMigrate(&models.Applicant{})
	config.DB.AutoMigrate(&models.Tag{})

	r := routes.SetupRouter()
	// running
	err = r.Run()
	if err != nil {
		return
	}
}
