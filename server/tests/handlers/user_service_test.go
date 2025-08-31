package tests

import (
	"context"
	pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"
	"github.com/rangodisco/zelvy/server/tests/factories"
	"github.com/rangodisco/zelvy/server/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	client := pb_usr.NewUserServiceClient(utils.Conn)
	body := factories.CreateAddUserRequest()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.AddUser(ctx, body)
	if err != nil {
		t.Fatal(err, resp)
	}

	assert.Equal(t, resp.Message, "Upsert successful")

}
