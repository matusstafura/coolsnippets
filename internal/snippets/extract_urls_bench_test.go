package snippets_test

import (
	"strings"
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func BenchmarkExtractURLs(b *testing.B) {
	sampleSmall := "Check https://example.com and http://test.com/page for links."
	sampleMedium := strings.Repeat(sampleSmall+" Some text without URLs. ", 100)
	sampleLarge := strings.Repeat(sampleSmall+" <a href=\"https://www.site.com\">link</a>", 5000)

	tests := []struct {
		name  string
		input string
	}{
		{"Small input (2 URLs)", sampleSmall},
		{"Medium input (~10k chars)", sampleMedium},
		{"Large input (~1M chars)", sampleLarge},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for b.Loop() {
				_ = snippets.ExtractURLs(tt.input)
			}
		})
	}
}

func BenchmarkExtractTagValues(b *testing.B) {
	sampleSmall := `<a href="https://example.com">Example</a> <a href="http://test.com/page">Test</a>`
	sampleMedium := strings.Repeat(sampleSmall+" Some text without tags. ", 100)
	sampleLarge := strings.Repeat(sampleSmall+" <div><a href=\"https://www.site.com\">Site</a></div>", 5000)

	tests := []struct {
		name  string
		input string
		tag   string
	}{
		{"Small input (2 tags)", sampleSmall, "href"},
		{"Medium input (~10k chars)", sampleMedium, "href"},
		{"Large input (~1M chars)", sampleLarge, "href"},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for b.Loop() {
				_ = snippets.ExtractTagValues(tt.input, tt.tag)
			}
		})
	}
}
