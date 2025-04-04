package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelvy/bot/utils"
)

var (
	Token     string
	AppID     string
	GuildID   string
	ChannelID string
)

// Parse command line arguments
func init() {
	// utils.ParseCommandLine()
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
		panic(e)
	}
}

func main() {

	// Init env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Init new Discord session
	dg, err := discordgo.New("Bot " + utils.Token)
	checkErr(err)

	// Register commands
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	cmdIds := make(map[string]string, len(Commands))

	for _, cmd := range Commands {
		rcmd, err := dg.ApplicationCommandCreate(utils.AppID, utils.GuildID, cmd)
		checkErr(err)
		cmdIds[rcmd.Name] = rcmd.ID
	}

	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	// Open session
	err = dg.Open()
	checkErr(err)

	// Close session once main function ends
	defer func(dg *discordgo.Session) {
		err := dg.Close()
		checkErr(err)
	}(dg)

	// Register scheduler
	s := utils.StartScheduler()
	defer func() {
		err := s.Shutdown()
		if err != nil {
			return
		}
	}()

	// Keep the bot running
	fmt.Println("Running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
}
