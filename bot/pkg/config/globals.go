package config

import (
	"flag"
	"os"
)

var (
	Token     string
	AppID     string
	GuildID   string
	ChannelID string
	ApiKey    string
)

func SetGlobals() {
	Token = os.Getenv("BOT_TOKEN")
	AppID = os.Getenv("BOT_APP_ID")
	GuildID = os.Getenv("BOT_GUILD_ID")
	ChannelID = os.Getenv("BOT_CHANNEL_ID")
	ApiKey = os.Getenv("API_KEY")

	if Token == "" {
		flag.Usage()
		os.Exit(1)
	}
}
