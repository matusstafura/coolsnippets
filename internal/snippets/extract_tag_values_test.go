package snippets_test

import (
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func TestExtractImageURLs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test with href relative URLs",
			input: `<img src="/images/pic1.jpg" alt="Pic 1"> and <img src="https://external.com/pic2.png" alt="Pic 2">`,
			want:  "/images/pic1.jpg\nhttps://external.com/pic2.png",
		},
		{
			name:  "Test with no href URLs",
			input: `<a>No link here</a>`,
			want:  "",
		},
		{
			name:  "Test with mixed href URLs",
			input: `<img src="/images/home.png" alt="Home"> <img src="http://example.com/logo.jpg" alt="Logo"> <img src="/images/about.png" alt="About">`,
			want:  "/images/home.png\nhttp://example.com/logo.jpg\n/images/about.png",
		},
		{
			name:  "Test with complex attributes",
			input: `<img src="/images/contact.png" class="nav-image" alt="Contact"> <img src="https://secure.com/banner.jpg" id="banner-image" alt="Banner">`,
			want:  "/images/contact.png\nhttps://secure.com/banner.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.ExtractTagValues(tt.input, "src")
			if tt.want != got {
				t.Errorf("ExtractHrefURLs() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExtractHrefURLs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test with href relative URLs",
			input: `<a href="/local/path">Local Link</a> and <a href="https://external.com/page">External Link</a>`,
			want:  "/local/path\nhttps://external.com/page",
		},
		{
			name:  "Test with no href URLs",
			input: `<a>No link here</a>`,
			want:  "",
		},
		{
			name:  "Test with mixed href URLs",
			input: `<a href="/home">Home</a> <a href="http://example.com">Example</a> <a href="/about">About Us</a>`,
			want:  "/home\nhttp://example.com\n/about",
		},
		{
			name:  "Test with complex href attributes",
			input: `<a href="/contact" class="nav-link">Contact</a> <a href="https://secure.com/login" id="login-link">Login</a>`,
			want:  "/contact\nhttps://secure.com/login",
		},
		{
			name:  "Test with href URLs containing query parameters",
			input: `<a href="/search?q=golang&page=2">Search</a> <a href="http://example.com/view?id=123">View</a>`,
			want:  "/search?q=golang&page=2\nhttp://example.com/view?id=123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.ExtractTagValues(tt.input, "href")
			if tt.want != got {
				t.Errorf("ExtractHrefURLs() = %q, want %q", got, tt.want)
			}
		})
	}
}
