package cli_test

import (
	"io"
	"strings"
	"testing"

	"github.com/matusstafura/coolsnippets/internal/cli"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		stdinR  io.Reader
		want    *cli.Config
		wantErr bool
	}{
		{
			name:    "missing utility",
			args:    []string{"-s", "test input"},
			wantErr: true,
		},
		{
			name:   "source from flag",
			args:   []string{"-u", "strip-tags", "-s", "<b>bold</b>"},
			stdinR: nil,
			want: &cli.Config{
				Utility: "strip-tags",
				Source:  "<b>bold</b>",
				Args:    []string{},
			},
			wantErr: false,
		},
		{
			name:   "source from stdin",
			args:   []string{"-u", "strip-tags"},
			stdinR: strings.NewReader("<i>italic</i>"),
			want: &cli.Config{
				Utility: "strip-tags",
				Source:  "<i>italic</i>",
				Args:    []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := cli.Parse(tt.args, tt.stdinR)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Utility != tt.want.Utility || got.Source != tt.want.Source {
					t.Errorf("Parse() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
