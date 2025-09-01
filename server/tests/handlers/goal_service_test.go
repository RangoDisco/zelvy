package tests

import (
	"context"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	"github.com/rangodisco/zelvy/server/tests/factories"
	"github.com/rangodisco/zelvy/server/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDisableGoals(t *testing.T) {
	client := pb_goa.NewGoalServiceClient(utils.Conn)
	body := factories.CreateDisableGoalRequest()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.DisableGoals(ctx, body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, body.Goals[0], resp.DisabledGoals[0])
	assert.Equal(t, len(body.Goals), len(resp.DisabledGoals)+len(resp.ErrorGoals))
	assert.Equal(t, body.Goals[len(body.Goals)-1], resp.ErrorGoals[len(resp.ErrorGoals)-1])
}
