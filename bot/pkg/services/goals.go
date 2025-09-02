package services

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/rangodisco/zelvy/bot/pkg/config"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	"google.golang.org/grpc/metadata"
	"time"
)

// DisableGoals disables any given goal for today
func DisableGoals(goals []string) (*pb_goa.DisableGoalsResponse, error) {

	client := pb_goa.NewGoalServiceClient(config.Conn)
	body := pb_goa.DisableGoalsRequest{}

	for _, goal := range goals {
		body.Goals = append(body.Goals, pb_goa.GoalType(pb_goa.GoalType_value[goal]))
	}

	ctx, cancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authorization": config.ApiKey})), 10*time.Second)
	defer cancel()

	resp, err := client.DisableGoals(ctx, &body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetGoalOptions() []discordgo.SelectMenuOption {
	return []discordgo.SelectMenuOption{
		{
			Label:       "Gym",
			Description: "No gym",
			Value:       pb_goa.GoalType_MAIN_WORKOUT_DURATION.String(),
			Default:     false,
			Emoji: &discordgo.ComponentEmoji{
				Name: "🏋️",
			},
		},
		{
			Label:       "Cardio",
			Description: "No cardio",
			Value:       pb_goa.GoalType_EXTRA_WORKOUT_DURATION.String(),
			Emoji: &discordgo.ComponentEmoji{
				Name: "👟",
			},
		},
		{
			Label:       "Eaten Kcal",
			Description: "No limit",
			Value:       pb_goa.GoalType_KCAL_CONSUMED.String(),
			Emoji: &discordgo.ComponentEmoji{
				Name: "🍛",
			},
		},
		{
			Label:       "Burned kcal",
			Description: "Lazy",
			Value:       pb_goa.GoalType_KCAL_BURNED.String(),
			Emoji: &discordgo.ComponentEmoji{
				Name: "🔥",
			},
		},
		{
			Label:       "Water",
			Description: "",
			Value:       pb_goa.GoalType_MILLILITER_DRANK.String(),
			Emoji: &discordgo.ComponentEmoji{
				Name: "🍶",
			},
		},
	}
}
