package generator

import (
	"strings"
)

type namespace struct {
	level            int
	path             string
	paths            []string
	placeholdersType string
	fieldType        string
	fieldName        string
	varName          string
	parent           *namespace
	children         []*namespace
	fields           []string
}

func newNamespace(path string) *namespace {
	seg := strings.Split(path, ".")
	name := toPascalCase(seg[len(seg)-1])
	return &namespace{
		level:            len(seg) - 1,
		path:             path,
		paths:            seg,
		placeholdersType: name + "Placeholders",
		fieldType:        toCamelCase(name) + "Locale",
		fieldName:        name,
	}
}

func (ns *namespace) extend(key string) *namespace {
	paths := make([]string, 0, 2)
	if ns.path != "" {
		paths = append(paths, ns.path)
	}
	paths = append(paths, key)
	path := strings.Join(paths, ".")

	newNS := newNamespace(path)
	newNS.parent = ns
	newNS.parent.children = append(newNS.parent.children, newNS)

	if len(ns.children) == 1 {
		ns.fieldType += "Nested"
	}

	return newNS
}

func (ns *namespace) prefixTypes() {
	if ns.level == 0 {
		return
	}
	ns.level--
	ns.placeholdersType = toPascalCase(ns.paths[ns.level]) + ns.placeholdersType
	ns.fieldType = toCamelCase(ns.paths[ns.level]) + toPascalCase(ns.fieldType)
}
