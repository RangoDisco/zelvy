package main

import (
	"bytes"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelby/bot/helpers"
	"io"
	"log"
	"net/http"
	"os"
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
					Description: "Désactive un ou plusieurs objectifs pour aujourd'hui",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "paypal",
					Description: "Link ton PayPal",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "email",
							Description: "L'email que tu utilises pour ton compte Paypal",
						},
					},
				},
			},
			Description: "All commands",
		},
	}
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// Send all values to the backend to disable it for today
		"metrics_to_disable": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.MessageComponentData()

			// Send the values to the backend
			helpers.SetOffDay(data.Values)
		},
	}
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"set": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse

			switch i.ApplicationCommandData().Options[0].Name {
			case "offday":
				// Prevent other users to use the command
				if i.Member.User.ID != os.Getenv("MAIN_USER_ID") {
					var response *discordgo.InteractionResponse
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
						Content: "Ok super, tu veux désactiver quel(s) objectif(s) pour aujourd'hui ?",
						Flags:   discordgo.MessageFlagsEphemeral,
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.SelectMenu{
										CustomID:    "metrics_to_disable",
										Placeholder: "Sélectionne un ou plusieurs objectifs",
										MinValues:   &minValues,
										MaxValues:   5,
										Options: []discordgo.SelectMenuOption{
											{
												Label:       "Séance principale",
												Description: "Pas de salle",
												Value:       "MAIN_WORKOUT_DURATION",
												Default:     false,
												Emoji: &discordgo.ComponentEmoji{
													Name: "🏋️",
												},
											},
											{
												Label:       "Sport additionnel",
												Description: "Pas de cardio",
												Value:       "EXTRA_WORKOUT_DURATION",
												Emoji: &discordgo.ComponentEmoji{
													Name: "👟",
												},
											},
											{
												Label:       "Calories consommées",
												Description: "Mange à balle",
												Value:       "KCAL_CONSUMED",
												Emoji: &discordgo.ComponentEmoji{
													Name: "🍛",
												},
											},
											{
												Label:       "Calories brulées",
												Description: "Pas bouger",
												Value:       "KCAL_BURNED",
												Emoji: &discordgo.ComponentEmoji{
													Name: "🔥",
												},
											},
											{
												Label:       "Eau",
												Description: "Pas d'eau",
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
				checkErr(err)
			case "paypal":
				response = handlePaypalCommand(i)
				err := s.InteractionRespond(i.Interaction, response)
				checkErr(err)
			}
		},
	}
)

func handlePaypalCommand(m *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	var c string
	o := m.ApplicationCommandData().Options[0].Options

	// Send error to user in case no email was provided
	if len(o) == 0 {
		c = "Merci de rentrer l'email lié à ton compte Paypal"
		return returnEphemeralInteraction(c)
	}

	email := o[0].StringValue()
	// Create user in database
	status := createUser(m.Member.User, email)

	// Let the user know if the email was added successfully
	switch status {
	case 200:
		c = "Email modifiée avec succès"
	case 201:
		c = "Paypal ajouté avec succès"
	default:
		c = "Erreur lors de l'ajout de ton compte Paypal"
	}

	return returnEphemeralInteraction(c)
}

/**
 * createUser Send user's info to backend and create a new user or update the existing email
 */
func createUser(u *discordgo.User, email string) int {
	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")

	b := CreateUserBody{
		Username:    u.GlobalName,
		DiscordID:   u.ID,
		PaypalEmail: email,
	}

	j, err := json.Marshal(b)
	checkErr(err)

	r, err := http.NewRequest("POST", baseUrl+"/api/users", bytes.NewBuffer(j))
	checkErr(err)

	r.Header.Add("X-API-KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(r)
	checkErr(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	return resp.StatusCode
}

func returnEphemeralInteraction(c string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
