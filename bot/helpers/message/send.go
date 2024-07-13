package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
)

type WorkoutData struct {
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     int    `json:"duration"`
}

type Metric struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Value        int    `json:"value"`
	DisplayValue string `json:"displayValue"`
	Threshold    string `json:"threshold"`
	Success      bool   `json:"success"`
}

type Metrics struct {
	ID       string        `json:"id"`
	Date     string        `json:"date"`
	Steps    int           `json:"steps"`
	Metrics  []Metric      `json:"metrics"`
	Workouts []WorkoutData `json:"workouts"`
}

func SendRecap(s *discordgo.Session, channelID string, metrics Metrics) {
	// Create new embed
	embed := NewEmbed().
		SetTitle("Stats du jour")

	// Add metrics Fields
	for _, metric := range metrics.Metrics {
		embed.AddField(metric.Name, formatField(metric.Success, metric.DisplayValue, metric.Threshold))
	}

	embed.InlineAllFields()

	sendEmbedMessage(s, channelID, embed.MessageEmbed)

}

func SendWorkoutsDetails(s *discordgo.Session, channelID string, metrics Metrics) {
	embed := NewEmbed().
		SetTitle("Séances")

	// Add workouts Fields
	for _, workout := range metrics.Workouts {
		embed.AddField(workout.Name, strconv.Itoa(workout.Duration)+" min")
	}

	sendEmbedMessage(s, channelID, embed.MessageEmbed)
}

func SendResults(s *discordgo.Session, channelID string, success bool, winner *discordgo.User) {
	var message string
	if success {
		message = "Pas de gagnant aujourd'hui, mais ça aurait dû être " + winner.Username
	} else {
		message = "Gagnant du jour: " + fmt.Sprintf("<@%s>", winner.ID) + ", bien joué chacal"
	}

	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Fatal(err)
	}
}

func formatField(success bool, value string, threshold string) string {

	if success {
		return value + "/" + threshold + " :white_check_mark:"
	}
	return value + "/" + threshold + " :x:"
}

func sendEmbedMessage(s *discordgo.Session, channelID string, embed *discordgo.MessageEmbed) {
	// Send message
	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Fatal(err)
	}
}
