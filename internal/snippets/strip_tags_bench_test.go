package snippets_test

import (
	"strings"
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func BenchmarkStripTags(b *testing.B) {
	sampleSmall := `<div class="container" id="main"><a href="https://example.com" style="color:red;">Link</a></div>`
	sampleMedium := strings.Repeat(sampleSmall+" Some text without attributes. ", 100)
	sampleLarge := strings.Repeat(sampleSmall+" <img src=\"image.jpg\" alt=\"Image\">", 5000)

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
				_ = snippets.StripAttributes(tt.input, nil)
			}
		})
	}
}
