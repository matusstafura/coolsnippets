package snippets

import "strings"

// ExtractTagValues parses the input string and extracts all values for a
// specified HTML attribute tag. Values are captured between quotes and
// returned as a newline-separated string with all found attributes.
func ExtractTagValues(input string, tag string) string {
	var hrefPrefix = tag + `="`

	var urls []string

	for i := 0; i < len(input); {
		if strings.HasPrefix(input[i:], hrefPrefix) {
			start := i + len(hrefPrefix)
			j := start

			for j < len(input) && input[j] != '"' {
				j++
			}

			url := input[start:j]
			urls = append(urls, url)

			i = j
		} else {
			i++
		}
	}

	return saveStrings(urls)
}
