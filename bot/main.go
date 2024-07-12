package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"time"
)

// Bot parameters
var (
	Token     string
	AppID     string
	GuildID   string
	ChannelID string
)

// Parse command line arguments
func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&AppID, "a", "", "Application ID")
	flag.StringVar(&GuildID, "g", "", "Guild ID")
	flag.StringVar(&ChannelID, "c", "", "Channel ID")

	flag.Parse()

	if Token == "" {
		flag.Usage()
		os.Exit(1)
	}
}

// Register commands

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	// Instantiate new Discord session
	dg, err := discordgo.New("Bot " + Token)
	checkErr(err)

	// Register commands
	_, err = dg.ApplicationCommandBulkOverwrite(AppID, GuildID, []*discordgo.ApplicationCommand{
		{
			Name:        "test-1",
			Description: "Just a test",
		},
	})
	checkErr(err)

	// Add event handler
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()

		switch data.Name {
		case "test-1":
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Bien joué",
				},
			})
			checkErr(err)
		}
	})

	// Open session
	err = dg.Open()
	checkErr(err)

	// Close session once main function ends
	defer func(dg *discordgo.Session) {
		err := dg.Close()
		checkErr(err)
	}(dg)

	// Start a ticket to send a message every minute
	ticket := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticket.C {
			sendScheduleMessage(dg, ChannelID, "Super 123")
		}
	}()

	defer ticket.Stop()

	// Keep the bot running
	fmt.Println("Running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func sendScheduleMessage(s *discordgo.Session, channelID, message string) {
	_, err := s.ChannelMessageSend(channelID, message)
	checkErr(err)

}
