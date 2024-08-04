package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

var (
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"set": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse

			data := i.MessageComponentData()
			switch data.Values[0] {
			case "MAIN_WORKOUT_DURATION":
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "This is the way.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			default:
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "It is not the way to go.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			}
			err := s.InteractionRespond(i.Interaction, response)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second) // Doing that so user won't see instant response.
			_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Anyways, now when you know how to use single select menus, let's see how multi select menus work. " +
					"Try calling `/selects multi` command.",
				Flags: discordgo.MessageFlagsEphemeral,
			})
			if err != nil {
				panic(err)
			}
		},
		"MAIN_WORKOUT_DURATION": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.MessageComponentData()

			const stackoverflowFormat = `https://stackoverflow.com/questions/tagged/%s`

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Here is your stackoverflow URL: " + fmt.Sprintf(stackoverflowFormat, strings.Join(data.Values, "+")),
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second) // Doing that so user won't see instant response.
			_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "But wait, there is more! You can also auto populate the select menu. Try executing `/selects auto-populated`.",
				Flags:   discordgo.MessageFlagsEphemeral,
			})
			if err != nil {
				panic(err)
			}
		},
	}
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"set": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse
			minValues := 1
			response = &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Now let's see how the multi-item select menu works: " +
						"try generating your own stackoverflow search link",
					Flags: discordgo.MessageFlagsEphemeral,
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.SelectMenu{
									CustomID:    "metrics_to_disable",
									Placeholder: "Quel objectif désactiver ?",
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
