package snippets_test

import (
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func TestUnescapeHtml(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Basic HTML unescape",
			input: "Hello &amp; welcome to &lt;Coding&gt;!",
			want:  "Hello & welcome to <Coding>!",
		},
		{
			name:  "Quotes and apostrophes",
			input: "materia&#322;u",
			want:  "materiału",
		},
		{
			name:  "Mixed entities",
			input: "Wz&oacute;r",
			want:  "Wzór",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.UnescapeHtml(tt.input)
			if tt.want != got {
				t.Errorf("%s: UnescapeHtml() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
