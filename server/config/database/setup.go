package database

import (
	"fmt"
	"os"
	"server/internal/enums"
	"time"

	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func SetupDatabase() error {
	var err error
	var ginMode = os.Getenv("GIN_MODE")
	switch ginMode {
	case "test":
		err = InitTestDatabase()
	default:
		err = InitDatabase()
	}

	return err
}

func InitDatabase() error {
	// Open a database connection
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return err
}

func InitTestDatabase() error {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
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
		return err
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
		return res.Error
	}

	goals := []models.Goal{
		{
			ID:         uuid.New(),
			Type:       enums.KcalBurned,
			Value:      1000,
			Name:       "",
			Unit:       "kcal",
			Comparison: "greater",
		},
		{
			ID:         uuid.New(),
			Type:       enums.KcalConsumed,
			Value:      2000,
			Name:       "",
			Unit:       "kcal",
			Comparison: "less",
		},
		{
			ID:         uuid.New(),
			Type:       enums.MilliliterDrank,
			Value:      2000,
			Name:       "",
			Unit:       "kcal",
			Comparison: "greater",
		},
		{
			ID:         uuid.New(),
			Type:       enums.MainWorkoutDuration,
			Value:      3600,
			Name:       "",
			Unit:       "mn",
			Comparison: "greater",
		},
		{
			ID:         uuid.New(),
			Type:       enums.ExtraWorkoutDuration,
			Value:      3600,
			Name:       "",
			Unit:       "mn",
			Comparison: "greater",
		},
	}

	for _, g := range goals {
		res = db.Create(&g)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
