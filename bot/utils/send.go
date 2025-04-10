package utils

import (
	"github.com/rangodisco/zelvy/bot/types"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

/**
 * SendRecap adds the metrics recap to the previously created thread (by CreateThread)
 */
func SendRecap(s *discordgo.Session, channelID string, summary types.Summary) {
	// Create new embed
	embed := NewEmbed().
		SetTitle("Stats du jour")

	if IsSuccess(summary.Metrics) {
		embed.SetThumbnail(os.Getenv("SUCCESS_PICTURE"))
	} else {
		embed.SetThumbnail(os.Getenv("FAILURE_PICTURE"))
	}

	// Add metrics Fields
	for _, metric := range summary.Metrics {
		embed.AddField(formatFieldTitle(metric.Name, metric.Success), metric.DisplayValue+"/"+metric.DisplayThreshold)
	}

	sendEmbedMessage(s, channelID, embed.MessageEmbed)

}

/**
 * SendWorkoutsDetails adds workouts to the previously created thread (by CreateThread)
 */
func SendWorkoutsDetails(s *discordgo.Session, channelID string, summary types.Summary) {
	embed := NewEmbed().
		SetTitle("Séances").
		SetThumbnail(os.Getenv("WORKOUTS_PICTURE"))

	// Add workouts Fields
	for _, workout := range summary.Workouts {
		embed.AddField(workout.Name, workout.Duration)
	}

	embed.InlineAllFields()

	sendEmbedMessage(s, channelID, embed.MessageEmbed)
}

/**
 * SendResults adds the winner on the previously created thread (by CreateThread)
 * In case I failed, mention the winner, otherwise only send its name
 */
func SendResults(s *discordgo.Session, channelID string, success bool, winner *discordgo.User) {
	var err error

	if success {
		_, err = s.ChannelMessageSend(channelID, "Pas de gagnant aujourd'hui, mais ça aurait dû être "+winner.GlobalName)
	} else {
		_, err = s.ChannelMessageSend(channelID, "Gagnant du jour: "+winner.Mention()+", bien joué pour tes 5€ chacal")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func formatFieldTitle(name string, success bool) string {

	if success {
		return name + " :white_check_mark:"
	}
	return name + " :x:"
}

func sendEmbedMessage(s *discordgo.Session, channelID string, embed *discordgo.MessageEmbed) {
	// Send message
	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Fatal(err)
	}
}

func SendScheduleMessage(s *discordgo.Session) {
	// Fetch Summary
	summary, err := FetchSummary()
	if err != nil {
		log.Fatalf("Error fetching summary: %v", err)
	}

	// Calculate results
	isSuccess := IsSuccess(summary.Metrics)

	// Create thread
	thread := CreateThread(s, ChannelID, isSuccess)

	// Send first stats message
	SendRecap(s, thread.ID, summary)

	// Send workout details
	SendWorkoutsDetails(s, thread.ID, summary)

	// Get Discord profile of winner
	winner, _ := s.User(summary.Winner.DiscordID)

	// Send results
	SendResults(s, thread.ID, isSuccess, winner)
}
