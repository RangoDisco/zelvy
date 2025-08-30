package setup

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelvy/bot/pkg/config"
	"github.com/rangodisco/zelvy/bot/pkg/utils"
	"github.com/rangodisco/zelvy/bot/pkg/utils/commands"
	"github.com/rangodisco/zelvy/bot/pkg/utils/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	dg  *discordgo.Session
	err error
)

// Setup setups everything needed for the discord bot to work
func Setup() error {
	config.SetGlobals()

	// Init new Discord session
	dg, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		return err
	}

	err = registerCommands()
	if err != nil {
		return err
	}

	err = setupSession()
	if err != nil {
		return err
	}

	grpc.SetupClient()

	//setupScheduler()

	// Keep the program running
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-done

	// Keep the bot running
	fmt.Println("Bot running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc

	return nil
}

// registerCommands registers each command contained in the commands util
func registerCommands() error {
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := commands.Handlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	cmdIds := make(map[string]string, len(commands.Commands))

	for _, cmd := range commands.Commands {
		rcmd, err := dg.ApplicationCommandCreate(config.AppID, config.GuildID, cmd)
		if err != nil {
			return err
		}
		cmdIds[rcmd.Name] = rcmd.ID
	}

	if err != nil {
		return err
	}

	return nil
}

// setupSession opens a new discordgo session and closes it when the main function ends
func setupSession() error {
	err = dg.Open()
	if err != nil {
		return err
	}

	defer func(dg *discordgo.Session) {
		err := dg.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dg)

	return nil
}

func setupScheduler() {
	s := utils.StartScheduler(dg)

	defer func() {
		err := s.Shutdown()
		if err != nil {
			return
		}
	}()
}
