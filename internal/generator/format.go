package generator

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var (
	errInvalidVerb     = errors.New("invalid fmt verb")
	errInvalidFormat   = errors.New("invalid format")
	errUnsupportedVerb = errors.New("unsupported verb")
)

var (
	verbMatcher       = regexp.MustCompile("^%[#+\\-0]?[0-9]*(\\.[0-9]*)?[bcdeEfFgGoOpqstTUvxX]$")
	formatVerbMatcher = regexp.MustCompile("^(?:\\[])?\\w*\\(?([\\w%#.+-]*)\\)?$")
)

var (
	strVerbs = formatType{
		supported:   []string{"s", "q", "x", "X"},
		defaultVerb: "s",
	}
	intVerbs = formatType{
		supported:   []string{"b", "c", "d", "o", "O", "q", "x", "X", "U"},
		defaultVerb: "d",
	}
	floatVerbs = formatType{
		supported:   []string{"b", "e", "E", "f", "F", "g", "G", "x", "X"},
		defaultVerb: "f",
	}
)

var (
	formatTypesMap = map[string]formatType{
		"string":     strVerbs,
		"[]byte":     strVerbs,
		"int":        intVerbs,
		"int8":       intVerbs,
		"int16":      intVerbs,
		"int32":      intVerbs,
		"int64":      intVerbs,
		"uint":       intVerbs,
		"uint8":      intVerbs,
		"uint16":     intVerbs,
		"uint32":     intVerbs,
		"uint64":     intVerbs,
		"float32":    floatVerbs,
		"float64":    floatVerbs,
		"complex64":  floatVerbs,
		"complex132": floatVerbs,
	}
	defaultFormatType = formatType{
		supported:   []string{"v"},
		defaultVerb: "v",
	}
	defaultFormat = format{typeStr: "string", verb: "%s"}
)

type format struct {
	typeStr string
	verb    string
}

type formatType struct {
	supported   []string
	defaultVerb string
}

func formatFromStr(s string) (format, error) {
	if s == "" {
		return defaultFormat, nil
	}

	f := format{}
	matches := formatVerbMatcher.FindStringSubmatch(s)
	matchLen := len(matches)
	if matchLen == 0 {
		return format{}, fmt.Errorf("%w %q", errInvalidFormat, s)
	}
	if len(matches) > 1 {
		f.verb = matches[1]
	}

	typeStr := strings.Split(s, "(")[0]
	ft, isKnownType := formatTypesMap[typeStr]
	if !isKnownType {
		ft = defaultFormatType
	}

	if typeStr == "" {
		typeStr = "any"
	}
	f.typeStr = typeStr

	if f.verb == "" {
		f.verb = "%" + ft.defaultVerb
	}

	if !verbMatcher.MatchString(f.verb) {
		return format{}, fmt.Errorf("%w %q", errInvalidVerb, f.verb)
	}

	if !slices.Contains(ft.supported, f.verb[len(f.verb)-1:]) {
		return format{}, fmt.Errorf("%w %q for type %q", errUnsupportedVerb, f.verb, typeStr)
	}

	return f, nil
}
