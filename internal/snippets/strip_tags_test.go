package snippets_test

import (
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func TestStripTags(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "basic test",
			input: "<p>Hello World</p>",
			want:  "Hello World",
		},
		{
			name:  "with line break",
			input: "<style>Hello<br>World</style>",
			want:  "HelloWorld",
		},
		{
			name:  "multiple tags",
			input: "<div><h1>Title</h1><p>This is a paragraph.</p></div>",
			want:  "TitleThis is a paragraph.",
		},
		{
			name:  "nested tags",
			input: "<div><p><span>Nested</span> Tags</p></div>",
			want:  "Nested Tags",
		},
		{
			name:  "no tags",
			input: "Just a plain text.",
			want:  "Just a plain text.",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "tags with attributes",
			input: `<a href="url">Click Here</a>`,
			want:  "Click Here",
		},
		{
			name:  "self-closing tags",
			input: "Line1<br/>Line2",
			want:  "Line1Line2",
		},
		{
			name:  "tags with spaces",
			input: "Hello < br > World",
			want:  "Hello  World",
		},
		{
			name:  "complex HTML",
			input: `<div><h1>Title</h1><p>This is a <strong>test</strong>.</p></div>`,
			want:  "TitleThis is a test.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.StripTags(tt.input, "")
			if tt.want != got {
				t.Errorf(tt.name, "StripTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripTagsNewlines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "basic test",
			input: "Hello<br>World",
			want:  "Hello\nWorld",
		},
		{
			name:  "multiple tags",
			input: "<p>Hello</p><br><div>World</div>",
			want:  "Hello\n\n\nWorld",
		},
		{
			name:  "nested tags",
			input: "<div><p>Hello<br>World</p></div>",
			want:  "Hello\nWorld",
		},
		{
			name:  "no tags",
			input: "Just a plain text.",
			want:  "Just a plain text.",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "tags with attributes",
			input: `<a href="url">Click<br>Here</a>`,
			want:  "Click\nHere",
		},
		{
			name:  "consecutive tags",
			input: "Line1<br><br>Line2",
			want:  "Line1\n\nLine2",
		},
		{
			name:  "tags with spaces",
			input: "Hello < br > World",
			want:  "Hello \n World",
		},
		{
			name:  "complex HTML",
			input: `<div><h1>Title</h1><p>This is a <strong>test</strong>.<br>New line here.</p></div>`,
			want:  "Title\n\nThis is a \ntest\n.\nNew line here.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.StripTags(tt.input, "\n")
			if got != tt.want {
				t.Errorf(tt.name+" failed: got %q, want %q", got, tt.want)
			}
		})
	}
}
