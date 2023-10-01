package generator

import (
	"fmt"
	"slices"
	"strings"
)

type languageMap map[string]any

func (m languageMap) get(namespace string) (any, error) {
	ref := m
	seg := slices.DeleteFunc(strings.Split(namespace, "."), func(s string) bool {
		return s == ""
	})

	for i, key := range seg {
		v, ok := ref[key]
		if !ok {
			return nil, fmt.Errorf("key not found")
		}

		if str, isStr := v.(string); isStr && i == len(seg)-1 {
			return str, nil
		}

		if ref, ok = v.(map[string]any); !ok {
			return nil, fmt.Errorf("expected map, got %v", v)
		}
	}
	return ref, nil
}
