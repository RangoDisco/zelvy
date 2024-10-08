package tests

import (
	"context"
	"io"
	"testing"

	"server/components"

	"github.com/PuerkitoBio/goquery"
)

func TestNav(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = components.Navbar().Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("Failed to read template: %v", err)
	}

	// Expecting 1 navbar
	if doc.Find(`[data-testid="navTemplate"]`).Length() == 0 {
		t.Error("Navbar not found")
	}
}
