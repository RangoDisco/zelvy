package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron/v2"
	"log"
)

// StartScheduler starts the scheduler to check for summary each day
func StartScheduler(dg *discordgo.Session) gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("Error creating scheduler: %v", err)
	}

	// Schedule the job to run every day at 20h00
	_, err = s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(22, 0, 0),
			),
		),
		// Assign the task to the job
		gocron.NewTask(SendScheduleMessage, dg),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)
	if err != nil {
		log.Fatalf("Failed to schedule order sync: %v", err)
	}
	// Start the scheduler
	s.Start()
	return s
}
