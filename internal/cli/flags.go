package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Utility string
	Source  string
	Args    []string
}

func Parse(args []string, stdinR io.Reader) (*Config, error) {
	fs := flag.NewFlagSet("coolsnippets", flag.ContinueOnError)

	fs.SetOutput(io.Discard)

	var util, src string
	fs.StringVar(&util, "u", "", "Utility to use")
	fs.StringVar(&src, "s", "", "Source string or read from stdin")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if util == "" {
		return nil, fmt.Errorf("utility is required")
	}

	// Read from stdin if -s not provided
	if src == "" {
		data, err := io.ReadAll(stdinR)
		if err != nil {
			return nil, fmt.Errorf("failed to read stdin: %w", err)
		}
		src = string(data)
	}

	return &Config{
		Utility: util,
		Source:  src,
		Args:    fs.Args(),
	}, nil
}

func MustParse() *Config {
	cfg, err := Parse(os.Args[1:], os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	return cfg
}
