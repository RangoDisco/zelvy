package services

import (
	"fmt"

	"server/internal/enums"
	"server/internal/models"
	"server/pkg/types"
)

// TODO: rework whole system
func GetWorkoutTypeChart() (types.Chart, error) {
	var chart types.Chart
	chart.Type = "radar"

	// Get workouts for this week and last week
	thisWeek, lastWeek, err := fetchChartWorkouts()
	if err != nil {
		return types.Chart{}, err
	}

	// Get labels for all datasests
	labels := getWorkoutTypeLabels()
	fmt.Println(labels)

	chart.Labels = labels

	// Build datasets with workout data
	var datasets []types.Dataset

	datasets = append(datasets, BuildWorkoutTypeDataset(thisWeek, chart, 0, "Cette semaine"))
	datasets = append(datasets, BuildWorkoutTypeDataset(lastWeek, chart, 1, "La semaine dernière"))

	// Add sets to chart
	chart.Datasets = datasets

	return chart, nil

}

// Create dataset for workout type count chart (radar)
func BuildWorkoutTypeDataset(workouts []models.Workout, chart types.Chart, i int, label string) types.Dataset {
	data := []int{}

	// Get data for each label
	for _, l := range chart.Labels {
		count := 0
		// Iterate over workouts and add data to the corresponding index
		for _, w := range workouts {
			// TODO: absolutely find a fix to this monstrosity
			switch w.ActivityType {
			case enums.WorkoutTypeStrength:
				if l == "Salle" {
					count++
				}
			case enums.WorkoutTypeWalking:
				if l == "Marche" {
					count++
				}
			case enums.WorkoutTypeRunning:
				if l == "Footing" {
					count++
				}
			case enums.WorkoutTypeCycling:
				if l == "Vélo" {
					count++
				}
			}
		}

		// Add data to dataset
		data = append(data, count)
	}

	// Get colors for dataset
	colors := getColorByIndex(i)

	// Create dataset
	dataset := types.Dataset{
		Label:           label,
		Data:            data,
		BackgroundColor: colors.Background,
		BorderColor:     colors.Border,
	}

	// Add datasets to chart
	return dataset

}

// Get labels based on workout's activity type
func getWorkoutTypeLabels() []string {
	return []string{"Salle", "Marche", "Footing", "Vélo", "Autre"}
}

func getColorByIndex(i int) types.ChartColors {
	// Define background and border colors
	colors := []types.ChartColors{
		{Background: "rgba(8, 92, 167, 0.2)", Border: "rgba(8, 92, 167, 1)"},
		{Background: "rgba(191, 174, 113, 0.2)", Border: "rgba(191, 174, 113, 1)"},
	}

	// Return color based on index
	return colors[i]

}
