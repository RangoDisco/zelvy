package tests

import (
	"context"
	"io"
	"testing"

	"server/components"
	"server/tests/factories"

	"github.com/PuerkitoBio/goquery"
)

func TestDisplay(t *testing.T) {
	testSummary := factories.CreateSummaryViewModel()

	r, w := io.Pipe()

	go func() {
		_ = components.SummaryDetails(testSummary).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("Failed to read template: %v", err)
	}

	// Expect the main section to be rendered
	if doc.Find(`[data-testid="summaryDetailsTemplate"]`).Length() == 0 {
		t.Fatalf("Expected data-testid attribute to be rendered, but it wasn't")
	}

	// Expect the goals to be rendered
	if actualMetricsCount := doc.Find(`[data-testid="summaryDetailsTemplateMetric"]`).Length(); actualMetricsCount != len(testSummary.Goals) {
		t.Fatalf("Expected %d metrics to be rendered, but %d found", len(testSummary.Goals), actualMetricsCount)
	}

	// Expect the goals to contain the correct values
	doc.Find(`[data-testid="summaryDetailsTemplateMetric"]`).Each(func(i int, sel *goquery.Selection) {
		// Check name
		expectedName := testSummary.Goals[i].Name
		if actualName := sel.Find(`[data-testid="summaryDetailsTemplateMetricName"]`).Text(); actualName != expectedName {
			t.Fatalf("Expected metric name to be %q, but got %q", expectedName, actualName)
		}

		// Check picto
		expectedPicto := testSummary.Goals[i].Picto
		if actualPicto := sel.Find(`[data-testid="summaryDetailsTemplateMetricPicto"]`).Text(); actualPicto != expectedPicto {
			t.Fatalf("Expected metric picto to be %q, but got %q", expectedPicto, actualPicto)
		}

		// Check display value
		expectedDisplayValue := testSummary.Goals[i].DisplayValue
		if actualDisplayValue := sel.Find(`[data-testid="summaryDetailsTemplateMetricValue"]`).Text(); actualDisplayValue != expectedDisplayValue {
			t.Fatalf("Expected metric display value to be %q, but got %q", expectedDisplayValue, actualDisplayValue)
		}

		// Check the threshold
		expectedThreshold := testSummary.Goals[i].DisplayThreshold
		if actualThreshold := sel.Find(`[data-testid="summaryDetailsTemplateMetricThreshold"]`).Text(); actualThreshold != expectedThreshold {
			t.Fatalf("Expected metric threshold to be %q, but got %q", expectedThreshold, actualThreshold)
		}
	})

}
