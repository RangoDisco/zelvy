package utils

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

/**
 * Start the thread on the given channel with, as a title: whether the day was a success & the current date
 * Thread is later used by SendRecap to send the daily recap
 */
func CreateThread(s *discordgo.Session, channelID string, isSuccess bool) *discordgo.Channel {

	// Get current day
	date := time.Now().Format("2006-01-02")

	name := "Stats du " + date

	if isSuccess {
		name += " ✅"
	} else {
		name += " ❌"
	}

	// Thread config
	threadStart := &discordgo.ThreadStart{
		Name:                name,
		AutoArchiveDuration: 60,
		Type:                discordgo.ChannelTypeGuildPublicThread,
	}

	// Create the thread
	thread, err := s.ThreadStartComplex(channelID, threadStart)
	if err != nil {
		log.Fatal(err)
	}

	return thread
}
