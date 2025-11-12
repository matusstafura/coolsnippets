package cli

import "testing"

func Test_stripTagsCommand(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name:   "basic test",
			source: "<p>Hello <b>World</b></p>",
			args:   []string{},
			want:   "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := stripTagsCommand(tt.source, tt.args)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("stripTagsCommand() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("stripTagsCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backlinkCommand(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name:   "basic test",
			source: "example keyword text",
			args:   []string{"keyword", "1", "http://example.com"},
			want:   "example <a href=\"http://example.com\">keyword</a> text",
		},
		{
			name:    "missing args",
			source:  "<p>This is a sample text with keyword.</p>",
			args:    []string{"keyword", "1"},
			wantErr: true,
		},
		{
			name:    "non-numeric nth",
			source:  "<p>This is a sample text with keyword.</p>",
			args:    []string{"keyword", "first", "http://example.com"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := backlinkCommand(tt.source, tt.args)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("backlinkCommand() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("backlinkCommand() succeeded unexpectedly")
			}
			if tt.want != got {
				t.Errorf("backlinkCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripTagsNewlineCommand(t *testing.T) {
	tests := []struct {
		name   string
		source string
		args   []string
		want   string
	}{
		{
			name:   "basic test",
			source: "Hello <b>World",
			args:   []string{},
			want:   "Hello \nWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := stripTagsNewlineCommand(tt.source, tt.args)
			if tt.want != got {
				t.Errorf("stripTagsNewlineCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unescapeHtmlCommand(t *testing.T) {
	tests := []struct {
		name   string
		source string
		args   []string
		want   string
	}{
		{
			name:   "basic test",
			source: "Hello &amp; World &lt;3",
			args:   []string{},
			want:   "Hello & World <3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := unescapeHtmlCommand(tt.source, tt.args)
			if tt.want != got {
				t.Errorf("unescapeHtmlCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractUrlsCommand(t *testing.T) {
	tests := []struct {
		name   string
		source string
		args   []string
		want   string
	}{
		{
			name:   "basic test",
			source: `<a href="http://example.com">Example</a> <a href="https://test.com">Test</a>`,
			args:   []string{},
			want:   "http://example.com\nhttps://test.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := extractUrlsCommand(tt.source, tt.args)
			if tt.want != got {
				t.Errorf("extractUrlsCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripAttributesCommand(t *testing.T) {
	tests := []struct {
		name   string
		source string
		args   []string
		want   string
	}{
		{
			name:   "basic test",
			source: `<a href="http://example.com" class="link">Example</a>`,
			args:   []string{},
			want:   `<a>Example</a>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := stripAttributesCommand(tt.source, tt.args)
			if tt.want != got {
				t.Errorf("stripAttributesCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractTagValuesCommand(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name:   "basic test",
			source: `<a href="http://example.com">Example</a> <a href="https://test.com">Test</a>`,
			args:   []string{"href"},
			want:   "http://example.com\nhttps://test.com",
		},
		{
			name:    "missing tag arg",
			source:  `<a href="http://example.com">Example</a>`,
			args:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := extractTagValuesCommand(tt.source, tt.args)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("extractTagValuesCommand() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractTagValuesCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		wantErr bool
	}{
		{
			name: "strip-tags command",
			cfg: &Config{
				Utility: "strip-tags",
				Source:  "<p>Hello <b>World</b></p>",
				Args:    []string{},
			},
		},
		{
			name: "backlink command with missing args",
			cfg: &Config{
				Utility: "backlink",
				Source:  "example keyword text",
				Args:    []string{"keyword", "1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := ExecuteCommand(tt.cfg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ExecuteCommand() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ExecuteCommand() succeeded unexpectedly")
			}
		})
	}
}
