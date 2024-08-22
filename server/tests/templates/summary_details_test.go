package tests

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/rangodisco/zelvy/server/components"
	"github.com/rangodisco/zelvy/server/tests/factories"
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

	// Expect the metrics to be rendered
	if actualMetricsCount := doc.Find(`[data-testid="summaryDetailsTemplateMetric"]`).Length(); actualMetricsCount != len(testSummary.Metrics) {
		t.Fatalf("Expected %d metrics to be rendered, but %d found", len(testSummary.Metrics), actualMetricsCount)
	}

	// Expect the metrics to contains the correct values
	doc.Find(`[data-testid="summaryDetailsTemplateMetric"]`).Each(func(i int, sel *goquery.Selection) {
		// Check name
		expectedName := testSummary.Metrics[i].Name
		if actualName := sel.Find(`[data-testid="summaryDetailsTemplateMetricName"]`).Text(); actualName != expectedName {
			t.Fatalf("Expected metric name to be %q, but got %q", expectedName, actualName)
		}

		// Check picto
		expectedPicto := testSummary.Metrics[i].Picto
		if actualPicto := sel.Find(`[data-testid="summaryDetailsTemplateMetricPicto"]`).Text(); actualPicto != expectedPicto {
			t.Fatalf("Expected metric picto to be %q, but got %q", expectedPicto, actualPicto)
		}

		// Check display value
		expectedDisplayValue := testSummary.Metrics[i].DisplayValue
		if actualDisplayValue := sel.Find(`[data-testid="summaryDetailsTemplateMetricValue"]`).Text(); actualDisplayValue != expectedDisplayValue {
			t.Fatalf("Expected metric display value to be %q, but got %q", expectedDisplayValue, actualDisplayValue)
		}

		// Check threshold
		expectedThreshold := testSummary.Metrics[i].DisplayThreshold
		if actualThreshold := sel.Find(`[data-testid="summaryDetailsTemplateMetricThreshold"]`).Text(); actualThreshold != expectedThreshold {
			t.Fatalf("Expected metric threshold to be %q, but got %q", expectedThreshold, actualThreshold)
		}
	})

}
