package tests

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/rangodisco/zelby/server/components"
	"github.com/rangodisco/zelby/server/tests/factories"
)

func TestHomeDisplay(t *testing.T) {
	// Create a summary view model
	summaryViewModel := factories.CreateSummaryViewModel()
	// Pipe the rendered template into goquery
	r, w := io.Pipe()
	go func() {
		_ = components.Summary(summaryViewModel).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("Failed to read template: %v", err)
	}

	// Expect the main section to be rendered
	if doc.Find(`[data-testid="summaryTemplate"]`).Length() == 0 {
		t.Fatalf("Expected main section attribute to be rendered, but it wasn't")
	}

	// Expect date to be rendered
	if actualDate := doc.Find(`[data-testid="summaryTemplateDate"]`).Text(); actualDate != summaryViewModel.Date {
		t.Fatalf("Expected date to be %q, but got %q", summaryViewModel.Date, actualDate)
	}
}
