package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/dave/jennifer/jen"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

type (
	Generator struct {
		languages              []language.Tag
		languagesMaps          map[language.Tag]languageMap
		defaultLang            language.Tag
		writer                 *writer
		nestedLocalesMap       map[string]*namespace
		localesMap             map[string]*namespace
		localesConflicts       map[string]struct{}
		nestedLocalesConflicts map[string]struct{}
	}
	Params struct {
		SrcPath string
		PkgName string
	}
)

type method struct {
	Alias       string
	Type        string
	Name        string
	Params      jen.Statement
	ReturnTypes jen.Statement
	Body        jen.Statement
}

func NewGenerator(params Params) (*Generator, error) {
	log.Printf("Starting Lexigo generator from src %q", params.SrcPath)

	var defaultLang language.Tag
	var languages []language.Tag
	languagesSrc := map[language.Tag][]byte{}
	var languagesDisplayNames []string
	if err := filepath.Walk(params.SrcPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if ext := filepath.Ext(info.Name()); ext == ".json" {
			seg := strings.Split(info.Name(), ".")
			tag, err := language.Parse(seg[0])
			if err != nil {
				return fmt.Errorf("failed to parse language from file name: %s: %w", info.Name(), err)
			}

			displayName := display.English.Languages().Name(tag)
			languagesDisplayNames = append(languagesDisplayNames, displayName)

			log.Printf("Found language file %q for %q", path, displayName)
			if seg[1] == "default" {
				if defaultLang != (language.Tag{}) {
					return fmt.Errorf("found 2 or more default language files")
				}
				log.Printf("Setting default language to %q", displayName)
				defaultLang = tag
			}

			languages = append(languages, tag)
			languagesSrc[tag], err = os.ReadFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if len(languagesSrc) == 0 {
		return nil, fmt.Errorf("no locale file found")
	}

	log.Printf("Found %d locale files - %v", len(languagesDisplayNames), languagesDisplayNames)

	maps := make(map[language.Tag]languageMap, len(languagesSrc))
	for _, lang := range languages {
		var data map[string]any
		if err := json.Unmarshal(languagesSrc[lang], &data); err != nil {
			panic(err)
		}
		maps[lang] = data
	}

	g := &Generator{
		languages:              languages,
		languagesMaps:          maps,
		defaultLang:            defaultLang,
		writer:                 &writer{},
		nestedLocalesMap:       map[string]*namespace{},
		localesMap:             map[string]*namespace{},
		localesConflicts:       map[string]struct{}{},
		nestedLocalesConflicts: map[string]struct{}{},
	}

	if err := jen.
		Comment("Code generated by lexigo - DO NOT EDIT").Line().
		Id("package").Id(params.PkgName).
		Line().
		Id("import").
		Defs(
			jen.Lit("fmt"),
			jen.Lit("context"),
			jen.Lit("golang.org/x/text/language"),
			jen.Lit("github.com/MysteriousPotato/lexigo/pkg"),
		).
		Line().
		Var().
		DefsFunc(func(group *jen.Group) {
			for _, lang := range languages {
				group.Add(jen.
					Id(lang.String()).
					Op("=").
					Qual("golang.org/x/text/language", "MustParse").
					Call(jen.Lit(lang.String())),
				)
			}
		}).
		Line().Line().
		Var().Id("Matcher").Op("=").
		Qual("golang.org/x/text/language", "NewMatcher").
		Call(jen.Index().Qual("golang.org/x/text/language", "Tag").ValuesFunc(func(group *jen.Group) {
			for _, lang := range languages {
				group.Add(jen.Id(lang.String()))
			}
		})).
		Render(g.writer); err != nil {
		return nil, err
	}

	if err := jen.Type().Id("locale").InterfaceFunc(func(group *jen.Group) {
		for _, lang := range languages {
			group.Add(jen.Id(lang.String()).Call().String())
		}
	}).Render(g.writer); err != nil {
		return nil, err
	}

	if err := jen.Type().Id("placeholders").InterfaceFunc(func(group *jen.Group) {
		for _, lang := range languages {
			group.Add(jen.Id(lang.String()).Call().Index().Any())
		}
	}).Render(g.writer); err != nil {
		return nil, err
	}

	if err := jen.Type().Id("Locale").Struct(
		jen.Id("locale").Id("locale"),
		jen.Id("placeholders").Id("placeholders"),
	).Render(g.writer); err != nil {
		return nil, err
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  "Locale",
		Name:  "parse",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("lang").Qual("golang.org/x/text/language", "Tag"))
				group.Add(jen.Id("level").Int())
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.If(jen.Id("level").Op("==").Lit(1)).Block(
				jen.Id("lang, _, _").Op("=").Id("Matcher.Match(lang)"),
			),
			jen.Switch(jen.Id("lang")).BlockFunc(func(group *jen.Group) {
				for _, tag := range g.languages {
					group.Add(jen.Case(jen.Id(tag.String())).Block(
						jen.If(jen.Id("l").Dot("placeholders").Op("==").Nil()).Block(
							jen.Return(jen.Id("l").Dot("locale").Dot(tag.String()).Call()),
						),
						jen.Return(
							jen.Qual("fmt", "Sprintf").Call(
								jen.Id("l").Dot("locale").Dot(tag.String()).Call(),
								jen.Id("l").Dot("placeholders").Dot(tag.String()).Call().Op("..."),
							),
						),
					))
				}
			}),
			jen.Id("parent").Op(":=").Id("lang"),
			jen.If(jen.Id("level").Op(">").Lit(0)).Block(
				jen.Id("parent").Op("=").Id("lang").Dot("Parent").Call(),
			),
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("parent"))
				group.Add(jen.Id("level+1"))
			})),
		},
	}); err != nil {
		return nil, err
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  "Locale",
		Name:  "FromCtx",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("ctx").Qual("context", "Context"))
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.List(jen.Id("lang"), jen.Id("_")).Op(":=").Qual("github.com/MysteriousPotato/lexigo", "FromCtx").Call(jen.Id("ctx")),
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("lang"))
				group.Add(jen.Lit(0))
			})),
		},
	}); err != nil {
		return nil, err
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  "Locale",
		Name:  "FromString",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("lang").String())
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.List(jen.Id("tag"), jen.Id("_")).Op(":=").Qual("golang.org/x/text/language", "Parse").Call(jen.Id("lang")),
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("tag"))
				group.Add(jen.Lit(0))
			})),
		},
	}); err != nil {
		return nil, err
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  "Locale",
		Name:  "FromTag",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("lang").Qual("golang.org/x/text/language", "Tag"))
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("lang"))
				group.Add(jen.Lit(0))
			})),
		},
	}); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Generator) Exec(varName string) ([]byte, error) {
	varName = toPascalCase(varName)

	if err := g.parseLocales(&namespace{
		fieldName: varName,
		varName:   varName,
		fieldType: toCamelCase(varName),
	}); err != nil {
		return nil, err
	}

	if err := g.generate(); err != nil {
		return nil, err
	}

	buf, err := io.ReadAll(&g.writer.internal)
	if err != nil {
		return nil, err
	}

	return imports.Process("", buf, nil)
}

func (g *Generator) parseLocales(ns *namespace) error {
	if ns.path != "" {
		log.Printf("Found %q locale", ns.path)
	}

	langData := g.languagesMaps[g.defaultLang]
	localeData, err := langData.get(ns)
	if err != nil {
		return fmt.Errorf("invalid namespace %q for %q locale file: %w", ns.path, g.defaultLang.String(), err)
	}

	valueOf := reflect.ValueOf(localeData)
	switch valueOf.Kind() {
	case reflect.Map:
		keys := valueOf.MapKeys()
		ns.fields = make([]string, len(keys))
		for i, key := range keys {
			ns.fields[i] = key.String()
		}

		return g.addNestedLocale(ns, 0)
	case reflect.String:
		g.addLocale(ns, 0)
	default:
		return fmt.Errorf("unsupported type %q in locales file", valueOf.Type())
	}

	return nil
}

func (g *Generator) generate() error {
	nestedLocales := make([]*namespace, 0, len(g.nestedLocalesMap))
	for _, ns := range g.nestedLocalesMap {
		nestedLocales = append(nestedLocales, ns)
	}

	slices.SortFunc(nestedLocales, func(a, b *namespace) int {
		if a.fieldType < b.fieldType {
			return -1
		}
		if a.fieldType > b.fieldType {
			return 1
		}
		return 0
	})

	for _, ns := range nestedLocales {
		slices.SortFunc(ns.children, func(a, b *namespace) int {
			if a.fieldType < b.fieldType {
				return -1
			}
			if a.fieldType > b.fieldType {
				return 1
			}
			return 0
		})

		typeCodes := make(jen.Statement, len(ns.children))
		valueCodes := make(jen.Statement, len(ns.children))
		for _, child := range ns.children {
			valueCodes.Add(jen.Id(child.fieldName).Op(":").Id(child.fieldType + "{}"))
			typeCodes.Add(jen.Id(child.fieldName).Id(child.fieldType))
		}

		if ns.varName != "" {
			if err := jen.Var().Id(ns.varName).Op("=").Id(ns.fieldType).Values(valueCodes...).Render(g.writer); err != nil {
				return err
			}
		}
		if err := jen.Type().Id(ns.fieldType).Struct(typeCodes...).Render(g.writer); err != nil {
			return err
		}
	}

	locales := make([]*namespace, 0, len(g.localesMap))
	for _, ns := range g.localesMap {
		locales = append(locales, ns)
	}

	slices.SortFunc(locales, func(a, b *namespace) int {
		if a.fieldType < b.fieldType {
			return -1
		}
		if a.fieldType > b.fieldType {
			return 1
		}
		return 0
	})

	for _, ns := range locales {
		if err := g.writeLocale(ns); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) addLocale(ns *namespace, level int) {
	if conflict, ok := g.localesMap[ns.fieldType]; ok {
		delete(g.localesMap, conflict.fieldType)
		conflict.prefixTypes()
		g.addLocale(conflict, level+1)

		g.localesConflicts[ns.fieldType] = struct{}{}

		ns.prefixTypes()
		g.addLocale(ns, level+1)
		return
	}

	if level == 0 {
		if _, ok := g.localesConflicts[ns.fieldType]; ok {
			ns.prefixTypes()
			g.addLocale(ns, level+1)
			return
		}
	}

	g.localesMap[ns.fieldType] = ns
}

func (g *Generator) addNestedLocale(ns *namespace, level int) error {
	if level == 0 {
		for _, key := range ns.fields {
			if err := g.parseLocales(ns.extend(key)); err != nil {
				return err
			}
		}
	}

	if conflict, ok := g.nestedLocalesMap[ns.fieldType]; ok {
		delete(g.nestedLocalesMap, conflict.fieldType)
		conflict.prefixTypes()
		if err := g.addNestedLocale(conflict, level+1); err != nil {
			return err
		}

		g.nestedLocalesConflicts[ns.fieldType] = struct{}{}

		ns.prefixTypes()
		return g.addNestedLocale(ns, level+1)
	}

	if level == 0 {
		if _, ok := g.nestedLocalesConflicts[ns.fieldType]; ok {
			ns.prefixTypes()
			return g.addNestedLocale(ns, level+1)
		}
	}

	g.nestedLocalesMap[ns.fieldType] = ns
	return nil
}

func (g *Generator) writeLocale(ns *namespace) error {
	log.Printf("Generating %q locale", ns.path)

	defaultLang := g.languagesMaps[g.defaultLang]
	defaultLocale, err := defaultLang.get(ns)
	_, placeholders, err := extractPlaceholders(defaultLocale.(string))
	if err != nil {
		return fmt.Errorf("failed to extract placeholders from %q: %w", ns.path, err)
	}

	if err := jen.Type().Id(ns.fieldType).Struct().Render(g.writer); err != nil {
		return err
	}

	langPlaceholdersMap := make(map[language.Tag][]field, len(g.languagesMaps))
	langSwitchCodes := make(jen.Statement, len(g.languagesMaps))
	for _, tag := range g.languages {
		lang := g.languagesMaps[tag]
		locale, err := lang.get(ns)
		if err != nil {
			return fmt.Errorf("invalid namespace %q for %q locale file: %w", ns.path, tag.String(), err)
		}

		localeStr, langPlaceholders, err := extractPlaceholders(locale.(string))
		if err != nil {
			return fmt.Errorf("failed to extract placeholders from %q: %w", ns.path, err)
		}

		if len(placeholders) != len(langPlaceholders) {
			return fmt.Errorf(
				"%q missmatched placeholders for %q file: expected %d placholders, got %d",
				ns.path,
				tag.String(),
				len(placeholders),
				len(langPlaceholders),
			)
		}
		for _, p := range placeholders {
			if !slices.Contains(langPlaceholders, p) {
				return fmt.Errorf(
					"%q missmatched placeholders for %q file: missing or invalid type for %q",
					ns.path,
					tag.String(),
					p.Name,
				)
			}
		}

		if err := jen.
			Func().
			Params(jen.Id("l").Id(ns.fieldType)).
			Id(tag.String()).
			Call().
			String().
			Block(jen.Return(jen.Lit(localeStr))).
			Render(g.writer); err != nil {
			return err
		}

		langPlaceholdersMap[tag] = langPlaceholders
		langSwitchCodes.Add(jen.Case(jen.Id(tag.String())).Block(
			jen.ReturnFunc(func(group *jen.Group) {
				if placeholders == nil {
					group.Add(jen.Id("l").Dot(tag.String()).Call())
					return
				}

				group.Add(jen.Qual("fmt", "Sprintf").Call(
					jen.Id("l").Dot(tag.String()).Call(),
					jen.Id("placeholders").Dot(tag.String()).Call().Op("..."),
				))
			}),
		))
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  ns.fieldType,
		Name:  "parse",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("lang").Qual("golang.org/x/text/language", "Tag"))
				group.Add(jen.Id("level").Int())
				if placeholders != nil {
					group.Add(jen.Id("placeholders").Id(ns.placeholdersType))
				}
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.If(jen.Id("level").Op("==").Lit(1)).Block(
				jen.Id("lang, _, _").Op("=").Id("Matcher.Match(lang)"),
			),
			jen.Switch(jen.Id("lang")).Block(langSwitchCodes...),
			jen.Id("parent").Op(":=").Id("lang"),
			jen.If(jen.Id("level").Op(">").Lit(0)).Block(
				jen.Id("parent").Op("=").Id("lang").Dot("Parent").Call(),
			),
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("parent"))
				group.Add(jen.Id("level+1"))
				if placeholders != nil {
					group.Add(jen.Id("placeholders"))
				}
			})),
		},
	}); err != nil {
		return err
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  ns.fieldType,
		Name:  "FromCtx",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("ctx").Qual("context", "Context"))
				if placeholders != nil {
					group.Add(jen.Id("placeholders").Id(ns.placeholdersType))
				}
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.List(jen.Id("lang"), jen.Id("_")).Op(":=").Qual("github.com/MysteriousPotato/lexigo", "FromCtx").Call(jen.Id("ctx")),
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("lang"))
				group.Add(jen.Lit(0))
				if placeholders != nil {
					group.Add(jen.Id("placeholders"))
				}
			})),
		},
	}); err != nil {
		return err
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  ns.fieldType,
		Name:  "FromString",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("lang").String())
				if placeholders != nil {
					group.Add(jen.Id("placeholders").Id(ns.placeholdersType))
				}
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.List(jen.Id("tag"), jen.Id("_")).Op(":=").Qual("golang.org/x/text/language", "Parse").Call(jen.Id("lang")),
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("tag"))
				group.Add(jen.Lit(0))
				if placeholders != nil {
					group.Add(jen.Id("placeholders"))
				}
			})),
		},
	}); err != nil {
		return err
	}

	if err := g.writeMethod(method{
		Alias: "l",
		Type:  ns.fieldType,
		Name:  "FromTag",
		Params: jen.Statement{
			jen.CustomFunc(jen.Options{Separator: ","}, func(group *jen.Group) {
				group.Add(jen.Id("lang").Qual("golang.org/x/text/language", "Tag"))
				if placeholders != nil {
					group.Add(jen.Id("placeholders").Id(ns.placeholdersType))
				}
			}),
		},
		ReturnTypes: jen.Statement{jen.String()},
		Body: jen.Statement{
			jen.Return(jen.Id("l").Dot("parse").CallFunc(func(group *jen.Group) {
				group.Add(jen.Id("lang"))
				group.Add(jen.Lit(0))
				if placeholders != nil {
					group.Add(jen.Id("placeholders"))
				}
			})),
		},
	}); err != nil {
		return err
	}

	var localeMethodParams *jen.Statement
	if placeholders != nil {
		localeMethodParams = jen.Id("placeholders").Id(ns.placeholdersType)
	}

	if err := g.writeMethod(method{
		Alias:       "l",
		Type:        ns.fieldType,
		Name:        "Locale",
		Params:      jen.Statement{localeMethodParams},
		ReturnTypes: jen.Statement{jen.Id("Locale")},
		Body: jen.Statement{
			jen.Return(jen.Id("Locale").ValuesFunc(func(group *jen.Group) {
				group.Add(jen.Id("locale").Op(": ").Id("l"))
				if placeholders != nil {
					group.Add(jen.Id("placeholders").Op(": ").Id("placeholders"))
				}
			})),
		},
	}); err != nil {
		return err
	}

	if placeholders != nil {
		if err := jen.Type().Id(ns.placeholdersType).Struct(placeholders.toJen()...).Render(g.writer); err != nil {
			return err
		}

		for _, tag := range g.languages {
			langPlaceholders := langPlaceholdersMap[tag]
			if err := g.writeMethod(method{
				Alias:       "p",
				Type:        ns.placeholdersType,
				Name:        tag.String(),
				ReturnTypes: jen.Statement{jen.Index().Any()},
				Body: jen.Statement{
					jen.Return(jen.Index().Any().ValuesFunc(func(group *jen.Group) {
						for _, p := range langPlaceholders {
							group.Add(jen.Id("p").Dot(p.Name))
						}
					})),
				},
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Generator) writeMethod(method method) error {
	return jen.
		Func().
		Params(jen.Id(method.Alias).Id(method.Type)).
		Id(method.Name).
		Params(method.Params...).
		Params(method.ReturnTypes...).
		Block(method.Body...).
		Render(g.writer)
}
