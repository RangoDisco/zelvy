package database

import (
	"fmt"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	"github.com/rangodisco/zelvy/server/internal/models"
	"os"

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
	var ginMode = os.Getenv("APP_ENV")
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

	// Auto-migrate models
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
	}
	res := db.Create(&testUser)

	if res.Error != nil {
		return res.Error
	}

	goals := []models.Goal{
		{
			ID:         uuid.New(),
			Type:       pb_goa.GoalType_KCAL_BURNED.String(),
			Value:      1000,
			Name:       "Burned kcal",
			Unit:       "kcal",
			Comparison: "greater",
			Active:     true,
		},
		{
			ID:         uuid.New(),
			Type:       pb_goa.GoalType_KCAL_CONSUMED.String(),
			Value:      2000,
			Name:       "Eaten kcal",
			Unit:       "kcal",
			Comparison: "less",
			Active:     true,
		},
		{
			ID:         uuid.New(),
			Type:       pb_goa.GoalType_MILLILITER_DRANK.String(),
			Value:      2000,
			Name:       "Water drank",
			Unit:       "kcal",
			Comparison: "greater",
			Active:     true,
		},
		{
			ID:         uuid.New(),
			Type:       pb_goa.GoalType_MAIN_WORKOUT_DURATION.String(),
			Value:      3600,
			Name:       "Gym duration",
			Unit:       "mn",
			Comparison: "greater",
			Active:     true,
		},
		{
			ID:         uuid.New(),
			Type:       pb_goa.GoalType_EXTRA_WORKOUT_DURATION.String(),
			Value:      3600,
			Name:       "Cardio duration",
			Unit:       "mn",
			Comparison: "greater",
			Active:     true,
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
