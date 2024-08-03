package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelby/bot/helpers"
	"github.com/rangodisco/zelby/bot/helpers/message"
	"log"
	"os"
	"os/signal"
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

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	// Init env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Init new Discord session
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

	sendScheduleMessage(dg)

	// Keep the bot running
	fmt.Println("Running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
}

func sendScheduleMessage(s *discordgo.Session) {
	// Fetch Summary
	summary := helpers.FetchSummary()

	// Pick a winner
	winner := helpers.PickWinner(s)

	// Calculate results
	isSuccess := helpers.IsSuccess(summary.Metrics)

	// Create thread
	thread := message.CreateThread(s, ChannelID, isSuccess)

	// Send first stats message
	message.SendRecap(s, thread.ID, summary)

	// Send workout details
	message.SendWorkoutsDetails(s, thread.ID, summary)

	// Send results
	message.SendResults(s, thread.ID, isSuccess, winner)
}
