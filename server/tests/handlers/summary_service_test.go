package tests

import (
	"context"
	"github.com/joho/godotenv"
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	"github.com/rangodisco/zelvy/server/config"
	"github.com/rangodisco/zelvy/server/tests/utils"
	"os"
	"testing"
	"time"

	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/tests/factories"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Load main environment variables
	godotenv.Load(".env")

	// Ensure that the environment is set to test
	os.Setenv("GIN_MODE", "test")

	// Load environment variables
	config.LoadEnv()

	database.SetupDatabase()

	utils.SetupGrpc()

	m.Run()
}

func TestAddSummary(t *testing.T) {
	// Create an example input model
	body := factories.CreateSummaryInputModel()

	utils.Client = pb_sum.NewSummaryServiceClient(utils.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := utils.Client.AddSummary(ctx, body)
	if err != nil {
		t.Fatal(err, resp)
	}

	assert.NotNil(t, resp)
}

func TestGetSummary(t *testing.T) {

	utils.Client = pb_sum.NewSummaryServiceClient(utils.Conn)

	// Try fetching the latest summary
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := utils.Client.GetSummary(ctx, &pb_sum.GetSummaryResquest{})
	if err != nil {
		t.Fatal(err, resp)
	}

	assert.Greater(t, len(resp.Workouts), 0)
}
