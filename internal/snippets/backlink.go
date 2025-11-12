package snippets

import (
	"fmt"
	"strings"
)

// Backlink finds the nth occurrence of a keyword in source and wraps it
// with an HTML anchor tag. The keyword must be preceded by space or be
// at the start of the string. Returns the modified source text.
func Backlink(source, keyword string, nth int, url string) string {
	occurrence := 0
	i := 0
	keywordLen := len(keyword)
	var b strings.Builder

	for i < len(source) {

		if i+keywordLen <= len(source) && source[i:i+keywordLen] == keyword {
			if i == 0 || source[i-1] == ' ' {
				occurrence++
				link := fmt.Sprintf(`<a href="%s">%s</a>`, url, keyword)
				if occurrence == nth {
					b.WriteString(link)
					i += keywordLen
					continue
				}
			}
		}
		b.WriteByte(source[i])
		i++
	}

	return b.String()
}
