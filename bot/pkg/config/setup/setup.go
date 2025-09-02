package setup

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelvy/bot/pkg/config"
	"github.com/rangodisco/zelvy/bot/pkg/utils"
	"github.com/rangodisco/zelvy/bot/pkg/utils/commands"
)

var (
	dg  *discordgo.Session
	err error
)

// Setup setups everything needed for the discord bot to work
func Setup(errChan chan<- error, stopChan <-chan struct{}) {
	config.SetGlobals()

	// Init new Discord session
	dg, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		errChan <- err
		return
	}

	err = registerCommands()
	if err != nil {
		errChan <- err
		return
	}

	setupSession(errChan, stopChan)

	config.SetupClient()

	//setupScheduler()
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
func setupSession(errChan chan<- error, stopChan <-chan struct{}) {
	go func() {
		if err := dg.Open(); err != nil {
			errChan <- err
			return
		}
		defer dg.Close()
		<-stopChan
	}()
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
