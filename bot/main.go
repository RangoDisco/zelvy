package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelby/bot/helpers"
	"github.com/rangodisco/zelby/bot/helpers/message"
	"io"
	"log"
	"net/http"
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

	// Start a ticket to send a message every minute
	//ticket := time.NewTicker(1 * time.Minute)
	//go func() {
	//	for range ticket.C {
	//		sendScheduleMessage(dg, ChannelID)
	//	}
	//}()
	//
	//defer ticket.Stop()

	sendScheduleMessage(dg)

	// Keep the bot running
	fmt.Println("Running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func sendScheduleMessage(s *discordgo.Session) {
	// First fetch today's routes
	res, err := http.Get("http://localhost:8080/api/metrics/today")
	checkErr(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatal("Failed to get routes")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	// Unmarshal response body to Metrics struct
	var metrics message.Metrics
	if err := json.Unmarshal(body, &metrics); err != nil {
		log.Fatalf("error unmarshalling response body: %v", err)
	}

	// Pick a winner
	winner := helpers.PickWinner(s)

	// Calculate results
	//isSuccess := helpers.IsSuccess(metrics.Metrics)
	isSuccess := true

	// Create thread
	thread := message.CreateThread(s, ChannelID, isSuccess)

	// Send first stats message
	message.SendRecap(s, thread.ID, metrics)

	// Send workout details
	message.SendWorkoutsDetails(s, thread.ID, metrics)

	// Send results
	message.SendResults(s, thread.ID, isSuccess, winner)

}
