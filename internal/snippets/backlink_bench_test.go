package snippets_test

import (
	"strings"
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func BenchmarkBacklink(b *testing.B) {
	sampleSmall := "This is a sample text with keyword. Another keyword here."
	sampleMedium := strings.Repeat(sampleSmall+" Some more text with keyword. ", 100)
	sampleLarge := strings.Repeat(sampleSmall+" Keyword appears multiple times. ", 5000)

	tests := []struct {
		name  string
		input string
	}{
		{"Small input", sampleSmall},
		{"Medium input", sampleMedium},
		{"Large input", sampleLarge},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for b.Loop() {
				_ = snippets.Backlink(tt.input, "keyword", 2, "https://example.com")
			}
		})
	}
}
