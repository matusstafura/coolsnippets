package snippets_test

import (
	"strings"
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func BenchmarkUnescapeHtml(b *testing.B) {
	sampleSmall := "This &amp; that &lt; those &gt; these."
	sampleMedium := strings.Repeat(sampleSmall+" More &quot;text&quot; with &apos;entities&apos;. ", 100)
	sampleLarge := strings.Repeat(sampleSmall+" HTML &amp; entities appear multiple times. ", 5000)

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
				_ = snippets.UnescapeHtml(tt.input)
			}
		})
	}
}
