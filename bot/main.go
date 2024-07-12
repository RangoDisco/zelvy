package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

type Metrics struct {
	ID           string    `json:"id"`
	Date         string    `json:"date"`
	Steps        int       `json:"steps"`
	KcalBurned   int       `json:"kcalBurned"`
	KcalConsumed int       `json:"kcalConsumed"`
	Workouts     []Workout `json:"workouts"`
}

type Workout struct {
	ID           string `json:"id"`
	MetricsID    string `json:"metricsId"`
	Name         string `json:"name"`
	Duration     int    `json:"duration"`
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
}

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
	//ticket := time.NewTicker(1 * time.Minute)
	//go func() {
	//	for range ticket.C {
	//		sendScheduleMessage(dg, ChannelID)
	//	}
	//}()
	//
	//defer ticket.Stop()

	sendScheduleMessage(dg, ChannelID)

	// Keep the bot running
	fmt.Println("Running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func sendScheduleMessage(s *discordgo.Session, channelID string) {
	// First fetch today's metrics
	res, err := http.Get("http://localhost:8080/api/metrics/today")
	checkErr(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatal("Failed to get metrics")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	// Unmarshal response body to Metrics struct
	var metrics Metrics
	if err := json.Unmarshal(body, &metrics); err != nil {
		log.Fatalf("error unmarshalling response body: %v", err)
	}

	// Build message
	embed := createEmbedMessage(metrics)

	_, err = s.ChannelMessageSendEmbed(channelID, embed)
	checkErr(err)

}

func createEmbedMessage(metrics Metrics) *discordgo.MessageEmbed {

	embed := NewEmbed().
		SetTitle("Est-ce que c'est ok aujourd'hui ?").
		SetDescription("Voici les métriques du jour").
		AddField("Nombre de pas", strconv.Itoa(metrics.Steps)).
		AddField("Calories consomées", appendStatus(metrics.KcalConsumed, 2100)).
		AddField("Calories brulées", appendStatus(metrics.KcalBurned, 1000))

	for _, workout := range metrics.Workouts {
		embed.AddField(workout.Name, strconv.Itoa(workout.Duration)+" min")
	}

	return embed.MessageEmbed
}

func appendStatus(value int, threshold int) string {
	if value >= threshold {
		return strconv.Itoa(value) + "/" + strconv.Itoa(threshold) + " :white_check_mark:"
	}
	return strconv.Itoa(value) + "/" + strconv.Itoa(threshold) + " :x:"
}
