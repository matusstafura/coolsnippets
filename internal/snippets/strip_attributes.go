package snippets

import (
	"slices"
	"strings"
)

var defaultTagsToKeep = []string{"script", "style", "img", "input", "link", "meta"}

// StripAttributes removes HTML attributes from tags in the input string,
// except for specified tags which preserve their attributes. Self-closing
// tags remain intact and the cleaned HTML string is then returned.
func StripAttributes(input string, keep []string) string {
	if len(input) == 0 {
		return input
	}

	var output strings.Builder
	var specialTags []string
	output.Grow(len(input))

	if len(keep) == 0 {
		specialTags = defaultTagsToKeep
	} else {
		specialTags = keep
	}

	i := 0
	for i < len(input) {
		if input[i] != '<' {
			output.WriteByte(input[i])
			i++
			continue
		}

		tagStart := i + 1
		if tagStart < len(input) && input[tagStart] == '/' {
			tagStart++
		}

		tagEnd := tagStart
		for tagEnd < len(input) && isTagChar(input[tagEnd]) {
			tagEnd++
		}

		skipAttrStrip := false
		if tagEnd > tagStart {
			tagName := strings.ToLower(input[tagStart:tagEnd])
			skipAttrStrip = slices.Contains(specialTags, tagName)
		}

		output.WriteByte('<')
		i++

		if skipAttrStrip {
			for i < len(input) && input[i] != '>' {
				output.WriteByte(input[i])
				i++
			}
			if i < len(input) {
				output.WriteByte('>') // Write closing '>'
				i++
			}
			continue
		}

		inAttr := false
		for i < len(input) && input[i] != '>' {
			char := input[i]

			if char == ' ' || char == '\n' || char == '\t' || char == '=' {
				inAttr = true
			} else if !inAttr {
				output.WriteByte(char)
			}

			i++
		}

		if i < len(input) {
			output.WriteByte('>')
			i++
		}
	}

	return output.String()
}

func isTagChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
