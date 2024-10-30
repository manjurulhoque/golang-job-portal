package testing

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// BaseTestCase struct to encapsulate common setup and teardown logic
type BaseTestCase struct {
	Router *gin.Engine
}

// Setup method to initialize common dependencies
func (btc *BaseTestCase) Setup(t *testing.T, modelList ...interface{}) {
	// Initialize test database
	gin.SetMode(gin.TestMode)
	var err error
	config.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	// Run any necessary migrations
	config.DB.AutoMigrate(modelList...)

	// Set up the router
	//btc.Router = routes.SetupRouter()
}

// TearDown method to clean up after tests
func (btc *BaseTestCase) TearDown() {
	// Optional: Close the database connection or clean up resources
	sqlDb, err := config.DB.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDb.Close()
	if err != nil {
		panic(err)
	}
}
