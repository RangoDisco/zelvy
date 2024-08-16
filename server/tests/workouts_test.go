package tests

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/rangodisco/zelby/server/components"
	"github.com/rangodisco/zelby/server/tests/factories"
)

func TestWorkoutTemplate(t *testing.T) {

	testWorkouts := factories.CreateWorkoutViewModels()

	r, w := io.Pipe()

	go func() {
		_ = components.Workouts(testWorkouts).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("Failed to read template: %v", err)
	}

	// Expect the main section to be rendered
	if doc.Find(`[data-testid="workoutsTemplate"]`).Length() == 0 {
		t.Fatalf("Expected data-testid attribute to be rendered, but it wasn't")
	}

	// Expect the workouts to be rendered
	if doc.Find(`[data-testid="workoutsTemplateWorkout"]`).Length() != len(testWorkouts) {
		t.Fatalf("Expected %d workouts to be rendered, but %d found", len(testWorkouts), doc.Find(`[data-testid="workoutsTemplateWorkout"]`).Length())
	}

	// Expect both workouts to contains the correct values
	if actualWorkoutsCount := doc.Find(`[data-testid="workoutsTemplateWorkout"]`).Length(); actualWorkoutsCount != len(testWorkouts) {
		t.Fatalf("Expected %d workouts to be rendered, but %d found", len(testWorkouts), actualWorkoutsCount)
	}

	// Expect the workout to conains the correct values
	doc.Find(`[data-testid="workoutsTemplateWorkout"]`).Each(func(i int, sel *goquery.Selection) {
		// Check name
		expectedName := testWorkouts[i].Name
		if actualName := sel.Find(`[data-testid="workoutsTemplateWorkoutName"]`).Text(); actualName != expectedName {
			t.Fatalf("Expected workout name to be %q, but got %q", expectedName, actualName)
		}

		// Check duration
		expectedDuration := testWorkouts[i].Duration
		if actualDuration := sel.Find(`[data-testid="workoutsTemplateWorkoutDuration"]`).Text(); actualDuration != expectedDuration {
			t.Fatalf("Expected workout duration to be %q, but got %q", expectedDuration, actualDuration)
		}

		// Check activity type
		expectedActivityType := testWorkouts[i].ActivityType
		if actualActivityType := sel.Find(`[data-testid="workoutsTemplateWorkoutActivityType"]`).Text(); actualActivityType != expectedActivityType {
			t.Fatalf("Expected workout activity type to be %q, but got %q", expectedActivityType, actualActivityType)
		}
	})

}
