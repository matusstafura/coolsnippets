package snippets_test

import (
	"testing"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

func TestStripAttributes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		tags  []string
	}{
		{
			name:  "basic test",
			input: `<a href="url" title="link">Click Here</a>`,
			want:  `<a>Click Here</a>`,
			tags:  []string{},
		},
		{
			name:  "empty input",
			input: "",
			want:  "",
			tags:  []string{},
		},
		{
			name:  "multiple attributes",
			input: `<img src="image.jpg" alt="An image" width="500" height="600"/>`,
			want:  `<img src="image.jpg" alt="An image" width="500" height="600"/>`,
			tags:  []string{},
		},
		{
			name:  "no attributes",
			input: `<div>Hello World</div>`,
			want:  `<div>Hello World</div>`,
			tags:  []string{},
		},
		{
			name:  "self-closing tag with attributes",
			input: `<input type="text" name="username" value="user1"/>`,
			want:  `<input type="text" name="username" value="user1"/>`,
			tags:  []string{},
		},
		{
			name:  "nested tags with attributes",
			input: `<div class="container"><p id="paragraph">Text</p></div>`,
			want:  `<div><p>Text</p></div>`,
			tags:  []string{},
		},
		{
			name:  "empty tag with attributes",
			input: `<br style="clear:both;"/>`,
			want:  `<br>`,
			tags:  []string{},
		},
		{
			name:  "tag with multiple spaces in attributes",
			input: `<a    href="url"    title="link"   >Link</a>`,
			want:  `<a>Link</a>`,
			tags:  []string{},
		},
		{
			name:  "tag with newline in attributes",
			input: "<div\n class=\"container\"\n id=\"main\">\nContent\n</div>",
			want:  "<div>\nContent\n</div>",
			tags:  []string{},
		},
		{
			name:  "tag with tab in attributes",
			input: "<span\tstyle=\"color:red;\"\tdata-info=\"info\">\tText\t</span>",
			want:  "<span>\tText\t</span>",
			tags:  []string{},
		},
		{
			name:  "complex tag with attributes",
			input: `<section data-section="1" class="main-section" aria-label="Section 1">Content</section>`,
			want:  `<section>Content</section>`,
			tags:  []string{},
		},
		{
			name:  "tag with single quotes in attributes",
			input: `<a href='url' title='link'>Link</a>`,
			want:  `<a href='url' title='link'>Link</a>`,
			tags:  []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := snippets.StripAttributes(tt.input, tt.tags)
			if tt.want != got {
				t.Errorf("StripAttributes() %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
