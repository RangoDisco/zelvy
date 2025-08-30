package utils

import (
	"github.com/rangodisco/zelvy/bot/pkg/config"
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// SendRecap adds the metrics recap to the previously created thread (by CreateThread)
func SendRecap(s *discordgo.Session, channelID string, summary *pb_sum.GetSummaryResponse) {
	// Create new embed
	embed := NewEmbed().
		SetTitle("Today's stats")

	if IsSuccessful(summary.Goals) {
		embed.SetThumbnail(os.Getenv("SUCCESS_PICTURE"))
	} else {
		embed.SetThumbnail(os.Getenv("FAILURE_PICTURE"))
	}

	// Add metrics Fields
	for _, g := range summary.Goals {

		embed.AddField(formatFieldTitle(g.Name, g.IsSuccessful, g.IsOff), g.DisplayValue+"/"+g.DisplayThreshold)
	}

	sendEmbedMessage(s, channelID, embed.MessageEmbed)

}

/**
 * SendWorkoutsDetails adds workouts to the previously created thread (by CreateThread)
 */
func SendWorkoutsDetails(s *discordgo.Session, channelID string, summary *pb_sum.GetSummaryResponse) {
	embed := NewEmbed().
		SetTitle("Workouts").
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
		_, err = s.ChannelMessageSend(channelID, "No winner today, but it should have been "+winner.GlobalName)
	} else {
		_, err = s.ChannelMessageSend(channelID, "Today's winner: "+winner.Mention())
	}
	if err != nil {
		log.Fatal(err)
	}
}

func formatFieldTitle(name string, success bool, isOff bool) string {
	if isOff {
		return name + " :pause_button"
	}

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
	isSuccessful := IsSuccessful(summary.Goals)

	// Create thread
	thread := CreateThread(s, config.ChannelID, isSuccessful)

	// Send the first stats message
	SendRecap(s, thread.ID, summary)

	// Send workout details
	SendWorkoutsDetails(s, thread.ID, summary)

	// Get Discord profile of winner
	winner, _ := s.User(summary.Winner.DiscordId)

	// Send results
	SendResults(s, thread.ID, isSuccessful, winner)
}
