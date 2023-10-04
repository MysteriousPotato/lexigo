package generator

import (
	"fmt"
	"strings"
)

func toPascalCase(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func toCamelCase(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func extractPlaceholders(str string) (string, fields, error) {
	matches := placeholderMatcher.FindAllStringSubmatch(str, -1)
	var placeholders []field

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
