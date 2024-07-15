package message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelby/bot/helpers"
	"log"
	"os"
)

func SendRecap(s *discordgo.Session, channelID string, metrics helpers.Metrics) {
	// Create new embed
	embed := NewEmbed().
		SetTitle("Stats du jour")

	if helpers.IsSuccess(metrics.Metrics) {
		embed.SetThumbnail(os.Getenv("SUCCESS_PICTURE"))
	} else {
		embed.SetThumbnail(os.Getenv("FAILURE_PICTURE"))
	}

	// Add metrics Fields
	for _, metric := range metrics.Metrics {
		embed.AddField(metric.Name, formatField(metric.Success, metric.DisplayValue, metric.Threshold))
	}

	embed.InlineAllFields()

	sendEmbedMessage(s, channelID, embed.MessageEmbed)

}

func SendWorkoutsDetails(s *discordgo.Session, channelID string, metrics helpers.Metrics) {
	embed := NewEmbed().
		SetTitle("Séances").
		SetThumbnail(os.Getenv("WORKOUTS_PICTURE"))

	// Add workouts Fields
	for _, workout := range metrics.Workouts {
		embed.AddField(workout.Name, workout.Duration)
	}

	embed.InlineAllFields()

	sendEmbedMessage(s, channelID, embed.MessageEmbed)
}

func SendResults(s *discordgo.Session, channelID string, success bool, winner *discordgo.User) {
	var message string
	if success {
		message = "Pas de gagnant aujourd'hui, mais ça aurait dû être " + winner.GlobalName
	} else {
		message = "Gagnant du jour: " + winner.Mention() + ", bien joué pour tes 5€ chacal"
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
