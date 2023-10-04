package generator

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var placeholderMatcher = regexp.MustCompile(`{{\.(\S+)}}`)

func toPascalCase(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func toCamelCase(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func extractPlaceholders(str string) (string, fields, error) {
	matches := placeholderMatcher.FindAllStringSubmatch(str, -1)
	var placeholders fields

	if len(matches) > 0 {
		str = escapePercentChars(str)
		slices.Grow(placeholders, len(matches)-1)
	}

	for _, match := range matches {
		if len(match) == 0 {
			continue
		}

		for _, group := range match[1:] {
			seg := strings.Split(group, ":")
			fieldName := seg[0]

			if len(seg) == 1 {
				seg = append(seg, "")
			}

			f, err := formatFromStr(seg[1])
			if err != nil {
				return "", nil, fmt.Errorf("failed to parse format %q: %w", seg[1], err)
			}

			str = strings.Replace(str, match[0], f.verb, 1)

			placeholders = append(placeholders, field{
				Name: toPascalCase(fieldName),
				Type: f.typeStr,
			})
		}
	}

	return str, placeholders, nil
}

func escapePercentChars(s string) string {
	var indexes []int
	var isPlaceholder bool
	for i := 0; i < len(s); i++ {
		if !isPlaceholder && len(s) > i+1 && s[i:i+2] == "{{" {
			isPlaceholder = true
			i++
			continue
		}
		if isPlaceholder && len(s) > i+1 && s[i:i+2] == "}}" {
			isPlaceholder = false
			i++
			continue
		}

		if !isPlaceholder && s[i] == '%' {
			indexes = append(indexes, i)
		}
	}

	b := make([]byte, len(s)+len(indexes))
	var lastInsert int
	for i, idx := range indexes {
		var lastIdx int
		if i > 0 {
			lastIdx = indexes[i-1]
		}

		copy(b[lastIdx+i:idx+i+1], s[lastIdx:idx+1])
		b[idx+i+1] = '%'
		lastInsert = idx + 1
	}
	copy(b[lastInsert+len(indexes):], s[lastInsert:])

	return string(b)
}

func extendNamespace(namespace, key string) string {
	if namespace == "" {
		return key
	}
	return namespace + "." + key
}

func getNames(fieldName string) (string, string, string, string) {
	typeName := toCamelCase(fieldName)
	fieldName = toPascalCase(fieldName)

	return fieldName, typeName, typeName + "Value", fieldName + "Placeholders"
}
