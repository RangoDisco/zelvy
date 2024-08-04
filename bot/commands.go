package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelby/bot/helpers"
)

var (
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
			// Prevent other users to use the command TODO: change placeholder user id
			if i.Member.User.ID != "YOUR_USER_ID" {
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

			var response *discordgo.InteractionResponse
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
			if err != nil {
				panic(err)
			}
		},
	}
)
