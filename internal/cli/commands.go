package cli

import (
	"fmt"
	"strconv"

	"github.com/matusstafura/coolsnippets/internal/snippets"
)

type CommandFunc func(source string, args []string) (string, error)

var commandMap = map[string]CommandFunc{
	"strip-tags":         stripTagsCommand,
	"strip-tags-newline": stripTagsNewlineCommand,
	"unescape-html":      unescapeHtmlCommand,
	"extract-urls":       extractUrlsCommand,
	"extract-tag-values": extractTagValuesCommand,
	"backlink":           backlinkCommand,
	"strip-attributes":   stripAttributesCommand,
}

func stripTagsCommand(source string, args []string) (string, error) {
	return snippets.StripTags(source, ""), nil
}

func stripTagsNewlineCommand(source string, args []string) (string, error) {
	return snippets.StripTags(source, "\n"), nil
}

func unescapeHtmlCommand(source string, args []string) (string, error) {
	return snippets.UnescapeHtml(source), nil
}

func extractUrlsCommand(source string, args []string) (string, error) {
	return snippets.ExtractURLs(source), nil
}

func extractTagValuesCommand(source string, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("extract-tag-values requires: tag")
	}
	tag := args[0]

	return snippets.ExtractTagValues(source, tag), nil
}

func backlinkCommand(source string, args []string) (string, error) {
	if len(args) < 3 {
		return "", fmt.Errorf("backlink requires: keyword nth url")
	}

	keyword := args[0]
	nth, err := strconv.Atoi(args[1])
	if err != nil {
		return "", fmt.Errorf("nth must be a number: %w", err)
	}
	url := args[2]

	result := snippets.Backlink(source, keyword, nth, url)

	return result, nil
}

func stripAttributesCommand(source string, args []string) (string, error) {
	return snippets.StripAttributes(source, args), nil
}

func ExecuteCommand(cfg *Config) error {
	util := cfg.Utility
	source := cfg.Source
	args := cfg.Args

	result, err := commandMap[util](source, args)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}
