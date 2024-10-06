package utils

import (
	"flag"
	"os"
)

// Bot parameters
var (
	Token     string
	AppID     string
	GuildID   string
	ChannelID string
)

func ParseCommandLine() {
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
