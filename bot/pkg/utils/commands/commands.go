package commands

import (
	"github.com/rangodisco/zelvy/bot/pkg/services"
	"github.com/rangodisco/zelvy/bot/pkg/utils"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type CreateUserBody struct {
	Username    string `json:"username"`
	DiscordID   string `json:"discordId"`
	PaypalEmail string `json:"paypalEmail"`
}

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
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// Send all values to the backend to disable it for today
		"metrics_to_disable": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.MessageComponentData()

			// Send the values to the backend
			services.SetOffDay(data.Values)

			// Send a message to the user
			response := &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Done",
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
										Options: []discordgo.SelectMenuOption{
											{
												Label:       "Gym",
												Description: "No gym",
												Value:       "MAIN_WORKOUT_DURATION",
												Default:     false,
												Emoji: &discordgo.ComponentEmoji{
													Name: "🏋️",
												},
											},
											{
												Label:       "Cardio",
												Description: "No cardio",
												Value:       "EXTRA_WORKOUT_DURATION",
												Emoji: &discordgo.ComponentEmoji{
													Name: "👟",
												},
											},
											{
												Label:       "Eaten Kcal",
												Description: "No limit",
												Value:       "KCAL_CONSUMED",
												Emoji: &discordgo.ComponentEmoji{
													Name: "🍛",
												},
											},
											{
												Label:       "Burned kcal",
												Description: "Lazy",
												Value:       "KCAL_BURNED",
												Emoji: &discordgo.ComponentEmoji{
													Name: "🔥",
												},
											},
											{
												Label:       "Water",
												Description: "",
												Value:       "MILILITER_CONSUMED",
												Emoji: &discordgo.ComponentEmoji{
													Name: "🍶",
												},
											},
										},
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
