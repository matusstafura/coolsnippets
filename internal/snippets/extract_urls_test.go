package snippets_test

import (
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func TestExtractURLs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test with multiple URLs",
			input: "Here are some links: https://example.com and http://test.com/page.",
			want:  "https://example.com\nhttp://test.com/page",
		},
		{
			name:  "Test with no URLs",
			input: "This string has no links.",
			want:  "",
		},
		{
			name:  "Test with mixed content",
			input: "Visit us at https://site.com! Also check http://another-site.org for more info.",
			want:  "https://site.com\nhttp://another-site.org",
		},
		{
			name:  "Test with URLs at the start and end",
			input: "https://start.com is the start. The end is http://end.com",
			want:  "https://start.com\nhttp://end.com",
		},
		{
			name:  "Test with URLs containing special characters",
			input: "Check out https://example.com/path?query=param&another=value for details.",
			want:  "https://example.com/path?query=param&another=value",
		},
		{
			name:  "Test with URL followed by punctuation",
			input: "visit https://www.example.com/sub/product?q=height%20&page=1 and enjoy.",
			want:  "https://www.example.com/sub/product?q=height%20&page=1",
		},
		{
			name:  "Test with URL followed by HTML link",
			input: "visit <a href=\"https://www.example.com/sub/product?q=height%20&page=1\">this link</a> and enjoy.",
			want:  "https://www.example.com/sub/product?q=height%20&page=1",
		},
		{
			name:  "Test with URL followed by HTML tags",
			input: "visit <a href=\"https://www.example.com/product\" target=\"_blank\" rel=\"noopener\">click</a><p>hello</p>",
			want:  "https://www.example.com/product",
		},
		// xml
		{
			name:  "Test with URL followed by XML tags",
			input: `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml" <url><loc>https://example.com/posts/</loc><lastmod>2025-10-11T09:56:32+02:00</lastmod><changefreq>monthly</changefreq><priority>0.5</priority></url>`,
			want:  "http://www.sitemaps.org/schemas/sitemap/0.9\nhttp://www.w3.org/1999/xhtml\nhttps://example.com/posts/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.ExtractURLs(tt.input)
			if tt.want != got {
				t.Errorf("ExtractURLs() = %q, want %q", got, tt.want)
			}
		})
	}
}
