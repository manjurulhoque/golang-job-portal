package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/manjurulhoque/golang-job-portal/config"
	docs "github.com/manjurulhoque/golang-job-portal/docs"
	"github.com/manjurulhoque/golang-job-portal/models"
	"github.com/manjurulhoque/golang-job-portal/routes"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
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
	docs.SwaggerInfo_swagger.BasePath = "/v1/api"
	docs.SwaggerInfo_swagger.Host = "localhost:8080"
	docs.SwaggerInfo_swagger.Title = "Golang job portal swagger API"
	docs.SwaggerInfo_swagger.Description = "Golang job portal swagger server"

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// Routes
	r.GET("/", HealthCheck)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// running
	err = r.Run()
	if err != nil {
		return
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
