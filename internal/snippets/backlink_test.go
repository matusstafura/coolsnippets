package snippets_test

import (
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func TestBacklink(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		keyword string
		nth     int
		url     string
		want    string
		wantErr bool
	}{
		{
			name:    "Basic backlink test",
			source:  "<p>This is a sample text with keyword.</p>",
			keyword: "keyword",
			nth:     1,
			url:     "http://example.com",
			want:    "<p>This is a sample text with <a href=\"http://example.com\">keyword</a>.</p>",
			wantErr: false,
		},
		{
			name:    "Multiple occurrences",
			source:  "example keyword example text example",
			keyword: "example",
			nth:     2,
			url:     "http://example.com",
			want:    "example keyword <a href=\"http://example.com\">example</a> text example",
			wantErr: false,
		},
		{
			name:    "Nth occurrence not found",
			source:  "keyword example text",
			keyword: "example",
			nth:     2,
			url:     "http://example.com",
			want:    "keyword example text",
			wantErr: false,
		},
		{
			name:    "Keyword not in source",
			source:  "keyword text only",
			keyword: "example",
			nth:     1,
			url:     "http://example.com",
			want:    "keyword text only",
			wantErr: false,
		},
		{
			name:    "Nth is zero",
			source:  "keyword example text",
			keyword: "example",
			nth:     0,
			url:     "http://example.com",
			want:    "keyword example text",
			wantErr: false,
		},
		{
			name:    "Empty source",
			source:  "",
			keyword: "example",
			nth:     1,
			url:     "http://example.com",
			want:    "",
			wantErr: false,
		},
		{
			name:    "do not create link if already linked",
			source:  `This is an <a href="http://example.com">example</a> text.`,
			keyword: "example",
			nth:     1,
			url:     "http://example.com",
			want:    `This is an <a href="http://example.com">example</a> text.`,
			wantErr: false,
		},
		{
			name:    "keyword at the start of the source",
			source:  "example keyword text",
			keyword: "example",
			nth:     1,
			url:     "http://example.com",
			want:    "<a href=\"http://example.com\">example</a> keyword text",
			wantErr: false,
		},
		{
			name:    "keyword at the end of the source",
			source:  "keyword text example",
			keyword: "example",
			nth:     1,
			url:     "http://example.com",
			want:    "keyword text <a href=\"http://example.com\">example</a>",
			wantErr: false,
		},
		{
			name:    "keyword with punctuation",
			source:  "This is an example, indeed.",
			keyword: "example",
			nth:     1,
			url:     "http://example.com",
			want:    "This is an <a href=\"http://example.com\">example</a>, indeed.",
			wantErr: false,
		},
		{
			name:    "keyword preceded by punctuation",
			source:  "This is an!example text.",
			keyword: "example",
			nth:     1,
			url:     "http://example.com",
			want:    "This is an!example text.",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.Backlink(tt.source, tt.keyword, tt.nth, tt.url)
			if tt.wantErr {
				t.Fatal("Backlink() succeeded unexpectedly")
			}
			if tt.want != got {
				t.Errorf("Backlink() = %v, want %v", got, tt.want)
			}
		})
	}
}
