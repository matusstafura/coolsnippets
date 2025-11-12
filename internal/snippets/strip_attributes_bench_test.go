package snippets_test

import (
	"strings"
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func BenchmarkStripAttributes(b *testing.B) {
	sampleSmall := `<a href="url" title="link">Click Here</a><img src="image.jpg" alt="An image" width="500" height="600"/>`
	sampleMedium := strings.Repeat(sampleSmall, 100)
	sampleLarge := strings.Repeat(sampleSmall, 5000)

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
				_ = snippets.StripAttributes(tt.input, []string{})
			}
		})
	}
}
