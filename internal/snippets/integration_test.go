package snippets

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

var test_binary_path = "../../bin/test-binary"

func TestCLI(t *testing.T) {
	build := exec.Command("go", "build", "-o", test_binary_path, "../..")
	if err := build.Run(); err != nil {
		t.Fatalf("Failed to build: %v", err)
	}

	tests := []struct {
		name     string
		args     []string
		stdin    string
		expected string
	}{
		{
			name:     "strip-tags-newline",
			args:     []string{"-u", "strip-tags-newline"},
			stdin:    "hello<br>world",
			expected: "hello\nworld\n",
		},
		{
			name:     "backlink",
			args:     []string{"-u", "backlink", "example", "1", "http://test.com"},
			stdin:    "example text",
			expected: "<a href=\"http://test.com\">example</a> text\n",
		},
		{
			name:     "strip-tags-newline",
			args:     []string{"-u", "strip-tags-newline"},
			stdin:    "<p>Line1</p><br><div>Line2</div>",
			expected: "Line1\n\n\nLine2\n",
		},
		{
			name:     "extract-urls",
			args:     []string{"-u", "extract-urls"},
			stdin:    `Check out http://example.com, https://example.org! Also visit http://test.com/page?arg=1.`,
			expected: "http://example.com\nhttps://example.org\nhttp://test.com/page?arg=1\n",
		},
		{
			name:     "extract-tag-values",
			args:     []string{"-u", "extract-tag-values", "href"},
			stdin:    `<a href="http://example.com">Example</a> <a href="https://test.org/page">Test</a>`,
			expected: "http://example.com\nhttps://test.org/page\n",
		},
		{
			name:     "extract-tag-values styles",
			args:     []string{"-u", "extract-tag-values", "style"},
			stdin:    `<div style="color: red;">Red Text</div> <span style="font-size: 12px;">Small Text</span>`,
			expected: "color: red;\nfont-size: 12px;\n",
		},
		{
			name:     "unescape-html",
			args:     []string{"-u", "unescape-html"},
			stdin:    "Hello &amp; welcome to the world of &lt;coding&gt;!",
			expected: "Hello & welcome to the world of <coding>!\n",
		},
		{
			name:     "strip-tags",
			args:     []string{"-u", "strip-tags"},
			stdin:    "<h1>Title</h1><p>This is a <strong>test</strong> paragraph.</p>",
			expected: "TitleThis is a test paragraph.\n",
		},
		{
			name:     "strip-attributes",
			args:     []string{"-u", "strip-attributes"},
			stdin:    `<a href="http://example.com" title="Example">Link</a><img src="image.jpg" alt="Image">`,
			expected: `<a>Link</a><img src="image.jpg" alt="Image">` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(test_binary_path, tt.args...)
			cmd.Stdin = strings.NewReader(tt.stdin)
			var out bytes.Buffer
			cmd.Stdout = &out

			if err := cmd.Run(); err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			if got := out.String(); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
