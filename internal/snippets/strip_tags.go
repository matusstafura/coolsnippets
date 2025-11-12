package snippets

import "strings"

// StripTags removes all HTML tags from the input string and replaces each
// tag with the specified delimiter. Leading and trailing whitespace is
// trimmed from the result before returning the cleaned string.
func StripTags(input, delimiter string) string {
	var output strings.Builder
	inTag := false

	for _, char := range input {
		switch char {
		case '<':
			inTag = true
		case '>':
			inTag = false
			output.WriteString(delimiter)
		default:
			if !inTag {
				output.WriteRune(char)
			}
		}
	}

	return strings.TrimSpace(output.String())
}
