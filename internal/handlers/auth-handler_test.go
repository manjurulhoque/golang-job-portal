package handlers

import (
	"github.com/manjurulhoque/golang-job-portal/internal/models"
	tests "github.com/manjurulhoque/golang-job-portal/internal/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Set up a test database connection
//func SetupTestDB() (*gorm.DB, error) {
//	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}

func TestRegister(t *testing.T) {
	// Initialize BaseTestCase
	var btc tests.BaseTestCase
	btc.Setup(t, models.User{})
	//defer btc.TearDown() // Ensures TearDown is called after the test

	user := models.RegisterInput{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password",
		Role:     "user",
	}

	err := Register(&user)
	assert.NoError(t, err)
}

func TestLogin(t *testing.T) {
	// Initialize BaseTestCase
	var btc tests.BaseTestCase
	btc.Setup(t, models.User{})
	defer btc.TearDown() // Ensures TearDown is called after the test

	user := models.LoginData{
		Email:    "test@example.com",
		Password: "password",
	}

	err := Login(&user)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, "test@example.com", "Email should be the same")
}
