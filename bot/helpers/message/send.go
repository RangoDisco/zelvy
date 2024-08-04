package message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelby/bot/helpers"
	"log"
	"os"
)

/**
 * Send the metrics recap on the previously created thread (by CreateThread)
 */
func SendRecap(s *discordgo.Session, channelID string, summary helpers.Summary) {
	// Create new embed
	embed := NewEmbed().
		SetTitle("Stats du jour")

	if helpers.IsSuccess(summary.Metrics) {
		embed.SetThumbnail(os.Getenv("SUCCESS_PICTURE"))
	} else {
		embed.SetThumbnail(os.Getenv("FAILURE_PICTURE"))
	}

	// Add metrics Fields
	for _, metric := range summary.Metrics {
		embed.AddField(formatFieldTitle(metric.Name, metric.Success), metric.DisplayValue+"/"+metric.Threshold)
	}

	sendEmbedMessage(s, channelID, embed.MessageEmbed)

}

/**
 * Send the workouts on the previously created thread (by CreateThread)
 */
func SendWorkoutsDetails(s *discordgo.Session, channelID string, summary helpers.Summary) {
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
 * Send the winner on the previously created thread (by CreateThread)
 * In case I failed, mentioned the winner, otherwise only send its name
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
