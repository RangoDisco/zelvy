package services

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelvy/bot/pkg/config"
	pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"
	"google.golang.org/grpc/metadata"
	"time"
)

// HandlePaypalCommand allows users to upsert themselves with their email
func HandlePaypalCommand(m *discordgo.InteractionCreate) string {
	var c string
	o := m.ApplicationCommandData().Options[0].Options

	// Send error to user in case no email was provided
	if len(o) == 0 {
		c = "Please enter a valid email"
		return c
	}

	email := o[0].StringValue()
	// Create user in database
	err := createUser(m.Member.User, email)
	if err != nil {
		c = "An error occurred while creating/updating the user"
	} else {
		c = "User created/updated successfully"
	}

	return c
}

// createUser sends the user's to create a new user or update the existing email
func createUser(u *discordgo.User, email string) error {
	body := pb_usr.AddUserRequest{
		Username:    u.GlobalName,
		DiscordId:   u.ID,
		PaypalEmail: email,
	}

	client := pb_usr.NewUserServiceClient(config.Conn)
	ctx, cancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authorization": config.ApiKey})), 10*time.Second)
	defer cancel()

	_, err := client.AddUser(ctx, &body)
	if err != nil {
		return err
	}

	return nil
}
