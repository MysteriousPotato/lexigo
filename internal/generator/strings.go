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
			seg := strings.Split(group, ":%")
			fieldName := seg[0]

			expr := "s"
			if len(seg) > 1 {
				expr = seg[1]
			}

			var fieldType string
			switch expr {
			case "v":
				fieldType = "any"
				break
			case "s", "q":
				fieldType = "string"
				break
			case "d", "b", "o", "x", "X":
				fieldType = "int64"
			case "f", "g", "e":
				fieldType = "float64"
			default:
				return "", nil, fmt.Errorf("unsupported format expression '%s'", expr)
			}

			str = strings.Replace(str, match[0], "%"+expr, 1)

			placeholders = append(placeholders, field{
				Name: toPascalCase(fieldName),
				Type: fieldType,
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
