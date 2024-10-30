package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/manjurulhoque/golang-job-portal/docs"
	"github.com/manjurulhoque/golang-job-portal/internal/config"
	"github.com/manjurulhoque/golang-job-portal/internal/models"
	"github.com/manjurulhoque/golang-job-portal/internal/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
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
		slog.Error("Error loading .env file", "error", err.Error())
	}

	config.DB, err = gorm.Open(postgres.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		slog.Error("Error connecting to database", "error", err.Error())
		panic(err)
	}

	sqlDB, err := config.DB.DB()
	if err != nil {
		slog.Error("Error getting database connection", "error", err.Error())
		panic(err)
	}
	defer func(sqlDB *sql.DB) {
		err = sqlDB.Close()
		if err != nil {
			slog.Error("Error closing the database connection", "error", err.Error())
			panic(err)
		}
	}(sqlDB)

	err = config.DB.AutoMigrate(&models.User{}, &models.Job{}, &models.Applicant{}, &models.Tag{})
	if err != nil {
		slog.Error("Error migrating the schema", "error", err.Error())
		panic(err)
	}

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
		slog.Error("Error running the server", "error", err.Error())
		panic(err)
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
