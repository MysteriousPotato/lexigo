// Code generated by lexigo - DO NOT EDIT
package test

import (
	"context"
	"fmt"
	"github.com/MysteriousPotato/lexigo/pkg"
	"golang.org/x/text/language"
)

var (
	en = language.MustParse("en")
	fr = language.MustParse("fr")
)

var Matcher = language.NewMatcher([]language.Tag{en, fr})

type locale interface {
	en() string
	fr() string
}

type placeholders interface {
	en() []any
	fr() []any
}

type Locale struct {
	locale       locale
	placeholders placeholders
}

func (l Locale) parse(lang language.Tag, level int) string {
	if level == 1 {
		lang, _, _ = Matcher.Match(lang)
	}
	switch lang {
	case en:
		if l.placeholders == nil {
			return l.locale.en()
		}
		return fmt.Sprintf(l.locale.en(), l.placeholders.en()...)
	case fr:
		if l.placeholders == nil {
			return l.locale.fr()
		}
		return fmt.Sprintf(l.locale.fr(), l.placeholders.fr()...)
	}
	parent := lang
	if level > 0 {
		parent = lang.Parent()
	}
	return l.parse(parent, level+1)
}

func (l Locale) FromCtx(ctx context.Context) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0)
}

func (l Locale) FromString(lang string) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0)
}

func (l Locale) FromTag(lang language.Tag) string {
	return l.parse(lang, 0)
}

type myLocale struct{}

func (l myLocale) en() string {
	return "myLocale"
}

func (l myLocale) fr() string {
	return "myLocale fr"
}

func (l myLocale) parse(lang language.Tag, level int) string {
	if level == 1 {
		lang, _, _ = Matcher.Match(lang)
	}
	switch lang {
	case en:
		return l.en()
	case fr:
		return l.fr()
	}
	parent := lang
	if level > 0 {
		parent = lang.Parent()
	}
	return l.parse(parent, level+1)
}

func (l myLocale) FromCtx(ctx context.Context) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0)
}

func (l myLocale) FromString(lang string) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0)
}

func (l myLocale) FromTag(lang language.Tag) string {
	return l.parse(lang, 0)
}

func (l myLocale) Locale() Locale {
	return Locale{locale: l}
}

type myLocaleAny struct{}

func (l myLocaleAny) en() string {
	return "myLocale %+v"
}

func (l myLocaleAny) fr() string {
	return "myLocale %+v fr"
}

func (l myLocaleAny) parse(lang language.Tag, level int, placeholders MyLocaleAnyPlaceholders) string {
	if level == 1 {
		lang, _, _ = Matcher.Match(lang)
	}
	switch lang {
	case en:
		return fmt.Sprintf(l.en(), placeholders.en()...)
	case fr:
		return fmt.Sprintf(l.fr(), placeholders.fr()...)
	}
	parent := lang
	if level > 0 {
		parent = lang.Parent()
	}
	return l.parse(parent, level+1, placeholders)
}

func (l myLocaleAny) FromCtx(ctx context.Context, placeholders MyLocaleAnyPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleAny) FromString(lang string, placeholders MyLocaleAnyPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleAny) FromTag(lang language.Tag, placeholders MyLocaleAnyPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleAny) Locale(placeholders MyLocaleAnyPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type MyLocaleAnyPlaceholders struct {
	Placeholder any
}

func (p MyLocaleAnyPlaceholders) en() []any {
	return []any{p.Placeholder}
}

func (p MyLocaleAnyPlaceholders) fr() []any {
	return []any{p.Placeholder}
}

type myLocaleFloat struct{}

func (l myLocaleFloat) en() string {
	return "myLocale %.2f"
}

func (l myLocaleFloat) fr() string {
	return "myLocale %.2f fr"
}

func (l myLocaleFloat) parse(lang language.Tag, level int, placeholders MyLocaleFloatPlaceholders) string {
	if level == 1 {
		lang, _, _ = Matcher.Match(lang)
	}
	switch lang {
	case en:
		return fmt.Sprintf(l.en(), placeholders.en()...)
	case fr:
		return fmt.Sprintf(l.fr(), placeholders.fr()...)
	}
	parent := lang
	if level > 0 {
		parent = lang.Parent()
	}
	return l.parse(parent, level+1, placeholders)
}

func (l myLocaleFloat) FromCtx(ctx context.Context, placeholders MyLocaleFloatPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleFloat) FromString(lang string, placeholders MyLocaleFloatPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleFloat) FromTag(lang language.Tag, placeholders MyLocaleFloatPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleFloat) Locale(placeholders MyLocaleFloatPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type MyLocaleFloatPlaceholders struct {
	Placeholder float32
}

func (p MyLocaleFloatPlaceholders) en() []any {
	return []any{p.Placeholder}
}

func (p MyLocaleFloatPlaceholders) fr() []any {
	return []any{p.Placeholder}
}

type myLocaleInt struct{}

func (l myLocaleInt) en() string {
	return "myLocale %d"
}

func (l myLocaleInt) fr() string {
	return "myLocale %d"
}

func (l myLocaleInt) parse(lang language.Tag, level int, placeholders MyLocaleIntPlaceholders) string {
	if level == 1 {
		lang, _, _ = Matcher.Match(lang)
	}
	switch lang {
	case en:
		return fmt.Sprintf(l.en(), placeholders.en()...)
	case fr:
		return fmt.Sprintf(l.fr(), placeholders.fr()...)
	}
	parent := lang
	if level > 0 {
		parent = lang.Parent()
	}
	return l.parse(parent, level+1, placeholders)
}

func (l myLocaleInt) FromCtx(ctx context.Context, placeholders MyLocaleIntPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleInt) FromString(lang string, placeholders MyLocaleIntPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleInt) FromTag(lang language.Tag, placeholders MyLocaleIntPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleInt) Locale(placeholders MyLocaleIntPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type MyLocaleIntPlaceholders struct {
	Placeholder int
}

func (p MyLocaleIntPlaceholders) en() []any {
	return []any{p.Placeholder}
}

func (p MyLocaleIntPlaceholders) fr() []any {
	return []any{p.Placeholder}
}

type myLocaleStr struct{}

func (l myLocaleStr) en() string {
	return "myLocale %q"
}

func (l myLocaleStr) fr() string {
	return "myLocale %q fr"
}

func (l myLocaleStr) parse(lang language.Tag, level int, placeholders MyLocaleStrPlaceholders) string {
	if level == 1 {
		lang, _, _ = Matcher.Match(lang)
	}
	switch lang {
	case en:
		return fmt.Sprintf(l.en(), placeholders.en()...)
	case fr:
		return fmt.Sprintf(l.fr(), placeholders.fr()...)
	}
	parent := lang
	if level > 0 {
		parent = lang.Parent()
	}
	return l.parse(parent, level+1, placeholders)
}

func (l myLocaleStr) FromCtx(ctx context.Context, placeholders MyLocaleStrPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleStr) FromString(lang string, placeholders MyLocaleStrPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleStr) FromTag(lang language.Tag, placeholders MyLocaleStrPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleStr) Locale(placeholders MyLocaleStrPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type MyLocaleStrPlaceholders struct {
	Placeholder string
}

func (p MyLocaleStrPlaceholders) en() []any {
	return []any{p.Placeholder}
}

func (p MyLocaleStrPlaceholders) fr() []any {
	return []any{p.Placeholder}
}

type inner struct{}

func (l inner) en() string {
	return "nested %s %s"
}

func (l inner) fr() string {
	return "nested %s %s fr"
}

func (l inner) parse(lang language.Tag, level int, placeholders InnerPlaceholders) string {
	if level == 1 {
		lang, _, _ = Matcher.Match(lang)
	}
	switch lang {
	case en:
		return fmt.Sprintf(l.en(), placeholders.en()...)
	case fr:
		return fmt.Sprintf(l.fr(), placeholders.fr()...)
	}
	parent := lang
	if level > 0 {
		parent = lang.Parent()
	}
	return l.parse(parent, level+1, placeholders)
}

func (l inner) FromCtx(ctx context.Context, placeholders InnerPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l inner) FromString(lang string, placeholders InnerPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l inner) FromTag(lang language.Tag, placeholders InnerPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l inner) Locale(placeholders InnerPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type InnerPlaceholders struct {
	Placeholder1 string
	Placeholder2 string
}

func (p InnerPlaceholders) en() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

func (p InnerPlaceholders) fr() []any {
	return []any{p.Placeholder2, p.Placeholder1}
}

type outer struct {
	Inner inner
}

var Locales = locales{MyLocale: myLocale{}, MyLocaleAny: myLocaleAny{}, MyLocaleFloat: myLocaleFloat{}, MyLocaleInt: myLocaleInt{}, MyLocaleStr: myLocaleStr{}, Outer: outer{}}

type locales struct {
	MyLocale      myLocale
	MyLocaleAny   myLocaleAny
	MyLocaleFloat myLocaleFloat
	MyLocaleInt   myLocaleInt
	MyLocaleStr   myLocaleStr
	Outer         outer
}
