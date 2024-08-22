package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rangodisco/zelvy/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func SetupDatabase() {

	var ginMode = os.Getenv("GIN_MODE")
	switch ginMode {
	case "test":
		InitTestDatabase()
	default:
		InitDatabase()
	}
}

func InitDatabase() {
	// Open a database connection
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

func InitTestDatabase() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto-migrate your models
	err = db.AutoMigrate(
		&models.Summary{},
		&models.Metric{},
		&models.Workout{},
		&models.Goal{},
		&models.Offday{},
		&models.User{},
	)

	if err != nil {
		log.Fatal("Failed to migrate models")
	}

	// Seed database
	testUser := models.User{
		ID:          uuid.New(),
		Username:    "test123",
		DiscordID:   "123456789",
		PaypalEmail: "testEmail@gmail.com",
		CreatedAt:   time.Now(),
	}
	res := db.Create(&testUser)

	if res.Error != nil {
		panic(res.Error)
	}
}

func GetDB() *gorm.DB {
	return db
}
