package snippets

import "html"

// UnescapeHtml converts HTML entities in the input string back to their
// original characters. Common entities are decoded and the unescaped
// string is returned ready for display or further processing.
func UnescapeHtml(input string) string {
	return html.UnescapeString(input)
}
