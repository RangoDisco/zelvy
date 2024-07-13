package helpers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelby/bot/helpers/message"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func IsSuccess(metrics []message.Metric) bool {
	// For each metric, check if it's a success
	for _, metric := range metrics {
		if !metric.Success {
			return false
		}
	}
	return true
}

func PickWinner(s *discordgo.Session) *discordgo.User {
	// Get all users
	userIds := strings.Split(os.Getenv("USER_IDS"), ",")

	// Create a pseudo random number
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	winnerId := userIds[random.Intn(len(userIds))]

	winner, err := s.User(winnerId)
	if err != nil {
		log.Fatalln(err)
	}

	return winner

}
