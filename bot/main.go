package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelvy/bot/utils"
)

// Parse command line arguments
func init() {
	utils.ParseCommandLine()
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

	// Close the session once the main function ends
	defer func(dg *discordgo.Session) {
		err := dg.Close()
		checkErr(err)
	}(dg)

	utils.SetupClient()

	// Register scheduler
	s := utils.StartScheduler(dg)

	defer func() {
		err := s.Shutdown()
		if err != nil {
			return
		}
	}()

	// Keep the program running
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-done

	// Keep the bot running
	fmt.Println("Running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
}
