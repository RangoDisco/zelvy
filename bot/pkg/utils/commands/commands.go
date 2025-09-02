package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelvy/bot/pkg/services"
	"github.com/rangodisco/zelvy/bot/pkg/utils"
	"log"
	"os"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name: "set",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "offday",
					Description: "Disable one or multiple goal(s) for the day",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "paypal",
					Description: "Link your PayPal account",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "email",
							Description: "The email address linked to your PayPal account",
						},
					},
				},
			},
			Description: "All set commands",
		},
		{
			Name: "get",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "summary",
					Description: "Fetch summary for a given day",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "date",
							Description: "Wanted date formatted as YYYY-MM-DD (e.g 2023-07-31)",
						},
					},
				},
			},
			Description: "All get commands",
		},
	}
	ComponentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// Send all values to the backend to disable it for today
		"metrics_to_disable": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.MessageComponentData()

			resp, _ := services.DisableGoals(data.Values)

			var content string
			if len(resp.DisabledGoals) > 0 {
				content = "Failed " + string(rune(len(resp.DisabledGoals))) + "times"
			} else {
				content = "Success"
			}

			// Send a message to the user
			response := &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			}
			err := s.InteractionRespond(i.Interaction, response)
			if err != nil {
				log.Println(err)
			}
		},
	}
	Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"set": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse

			switch i.ApplicationCommandData().Options[0].Name {
			case "offday":
				// Prevent other users to use the command
				if i.Member.User.ID != os.Getenv("MAIN_USER_ID") {
					response = &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Bah oui mais non du coup",
							Flags:   discordgo.MessageFlagsEphemeral,
						},
					}
					err := s.InteractionRespond(i.Interaction, response)
					if err != nil {
						panic(err)
					}
					return
				}

				minValues := 1
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Which goal(s) would you like to disable today ?",
						Flags:   discordgo.MessageFlagsEphemeral,
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.SelectMenu{
										CustomID:    "metrics_to_disable",
										Placeholder: "Please select one or multiple goal(s)",
										MinValues:   &minValues,
										MaxValues:   5,
										Options:     services.GetGoalOptions(),
									},
								},
							},
						},
					},
				}
				err := s.InteractionRespond(i.Interaction, response)
				if err != nil {
					log.Fatal(err)
				}
			case "paypal":
				sRes := services.HandlePaypalCommand(i)
				response = returnEphemeralInteraction(sRes)
				err := s.InteractionRespond(i.Interaction, response)
				if err != nil {
					log.Fatal(err)
				}
			}
		},
		"get": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse
			switch i.ApplicationCommandData().Options[0].Name {
			case "summary":
				// Prevent other users to use the command
				if i.Member.User.ID != os.Getenv("MAIN_USER_ID") {
					response = &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Bah oui mais non du coup",
							Flags:   discordgo.MessageFlagsEphemeral,
						},
					}
					err := s.InteractionRespond(i.Interaction, response)
					if err != nil {
						log.Fatal(err)
					}
					return
				}
				// TODO check len
				// FOR NOW ONLY TODAYS SUMMARY CAN BE FETCHED
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Done",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
				err := s.InteractionRespond(i.Interaction, response)
				if err != nil {
					log.Fatal(err)
				}
				utils.SendScheduleMessage(s)
				return
			}
		},
	}
)

func returnEphemeralInteraction(c string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
