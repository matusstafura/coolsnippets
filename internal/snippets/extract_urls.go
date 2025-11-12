package snippets

import (
	"strings"
)

// ExtractURLs scans the input string and extracts all HTTP and HTTPS URLs
// found within the text. Trimmed punctuation and whitespace boundaries
// are respected. Returns URLs joined by newlines in a single string.
func ExtractURLs(input string) string {

	// Characters to trim off the end of URLs
	const trimChars = ".,;!?)]}"

	const (
		h1 = "http://"
		h2 = "https://"
	)

	var urls []string

	for i := 0; i < len(input); {
		if strings.HasPrefix(input[i:], h1) || strings.HasPrefix(input[i:], h2) {

			start := i
			j := i

			for j < len(input) {
				c := input[j]
				if c <= 32 || c == '<' || c == '>' || c == '"' || c == '\'' || c == '(' || c == ')' {
					break
				}
				j++
			}

			url := input[start:j]

			url = strings.TrimRight(url, trimChars)

			urls = append(urls, url)

			i = j
		} else {
			i++
		}
	}

	return saveStrings(urls)
}

// saveStrings joins a slice of strings into a single string with newlines.
func saveStrings(input []string) string {
	if len(input) == 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(len(input) / 2)

	for i, u := range input {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(u)
	}

	return b.String()
}
