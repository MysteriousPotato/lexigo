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

var Locales = localesNested{MyLocaleAny: myLocaleAnyLocale{}, MyLocaleFloat: myLocaleFloatLocale{}, MyLocaleInt: myLocaleIntLocale{}, MyLocale: myLocaleLocale{}, MyLocaleStr: myLocaleStrLocale{}, Outer1: outer1LocaleNested{}, Outer2: outer2LocaleNested{}, Outer3: outer3LocaleNested{}, Outer4: outer4LocaleNested{}, Outer5: outer5LocaleNested{}}

type localesNested struct {
	MyLocaleAny   myLocaleAnyLocale
	MyLocaleFloat myLocaleFloatLocale
	MyLocaleInt   myLocaleIntLocale
	MyLocale      myLocaleLocale
	MyLocaleStr   myLocaleStrLocale
	Outer1        outer1LocaleNested
	Outer2        outer2LocaleNested
	Outer3        outer3LocaleNested
	Outer4        outer4LocaleNested
	Outer5        outer5LocaleNested
}

type outer1LocaleNested struct {
	Inner outer1InnerLocale
}

type outer2LocaleNested struct {
	Inner outer2InnerLocale
}

type outer3LocaleNested struct {
	Outer1 outer3Outer1LocaleNested
}

type outer3Outer1LocaleNested struct {
	Inner outer3Outer1InnerLocale
}

type outer4LocaleNested struct {
	Outer1 outer4Outer1LocaleNested
}

type outer4Outer1LocaleNested struct {
	Inner outer4Outer1InnerLocale
}

type outer5LocaleNested struct {
	Outer6 outer6LocaleNested
}

type outer6LocaleNested struct {
	Outer1 outer6Outer1LocaleNested
}

type outer6Outer1LocaleNested struct {
	Inner outer6Outer1InnerLocale
}

type myLocaleAnyLocale struct{}

func (l myLocaleAnyLocale) en() string {
	return "myLocale %+v"
}

func (l myLocaleAnyLocale) fr() string {
	return "myLocale %+v fr"
}

func (l myLocaleAnyLocale) parse(lang language.Tag, level int, placeholders MyLocaleAnyPlaceholders) string {
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

func (l myLocaleAnyLocale) FromCtx(ctx context.Context, placeholders MyLocaleAnyPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleAnyLocale) FromString(lang string, placeholders MyLocaleAnyPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleAnyLocale) FromTag(lang language.Tag, placeholders MyLocaleAnyPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleAnyLocale) Locale(placeholders MyLocaleAnyPlaceholders) Locale {
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

type myLocaleFloatLocale struct{}

func (l myLocaleFloatLocale) en() string {
	return "myLocale %.2f"
}

func (l myLocaleFloatLocale) fr() string {
	return "myLocale %.2f fr"
}

func (l myLocaleFloatLocale) parse(lang language.Tag, level int, placeholders MyLocaleFloatPlaceholders) string {
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

func (l myLocaleFloatLocale) FromCtx(ctx context.Context, placeholders MyLocaleFloatPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleFloatLocale) FromString(lang string, placeholders MyLocaleFloatPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleFloatLocale) FromTag(lang language.Tag, placeholders MyLocaleFloatPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleFloatLocale) Locale(placeholders MyLocaleFloatPlaceholders) Locale {
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

type myLocaleIntLocale struct{}

func (l myLocaleIntLocale) en() string {
	return "myLocale %d"
}

func (l myLocaleIntLocale) fr() string {
	return "myLocale %d"
}

func (l myLocaleIntLocale) parse(lang language.Tag, level int, placeholders MyLocaleIntPlaceholders) string {
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

func (l myLocaleIntLocale) FromCtx(ctx context.Context, placeholders MyLocaleIntPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleIntLocale) FromString(lang string, placeholders MyLocaleIntPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleIntLocale) FromTag(lang language.Tag, placeholders MyLocaleIntPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleIntLocale) Locale(placeholders MyLocaleIntPlaceholders) Locale {
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

type myLocaleLocale struct{}

func (l myLocaleLocale) en() string {
	return "myLocale %"
}

func (l myLocaleLocale) fr() string {
	return "myLocale % fr"
}

func (l myLocaleLocale) parse(lang language.Tag, level int) string {
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

func (l myLocaleLocale) FromCtx(ctx context.Context) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0)
}

func (l myLocaleLocale) FromString(lang string) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0)
}

func (l myLocaleLocale) FromTag(lang language.Tag) string {
	return l.parse(lang, 0)
}

func (l myLocaleLocale) Locale() Locale {
	return Locale{locale: l}
}

type myLocaleStrLocale struct{}

func (l myLocaleStrLocale) en() string {
	return "myLocale %% %q"
}

func (l myLocaleStrLocale) fr() string {
	return "myLocale %% %q fr"
}

func (l myLocaleStrLocale) parse(lang language.Tag, level int, placeholders MyLocaleStrPlaceholders) string {
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

func (l myLocaleStrLocale) FromCtx(ctx context.Context, placeholders MyLocaleStrPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleStrLocale) FromString(lang string, placeholders MyLocaleStrPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l myLocaleStrLocale) FromTag(lang language.Tag, placeholders MyLocaleStrPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l myLocaleStrLocale) Locale(placeholders MyLocaleStrPlaceholders) Locale {
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

type outer1InnerLocale struct{}

func (l outer1InnerLocale) en() string {
	return "nested 1 %s %s"
}

func (l outer1InnerLocale) fr() string {
	return "nested 1 %s %s fr"
}

func (l outer1InnerLocale) parse(lang language.Tag, level int, placeholders Outer1InnerPlaceholders) string {
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

func (l outer1InnerLocale) FromCtx(ctx context.Context, placeholders Outer1InnerPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l outer1InnerLocale) FromString(lang string, placeholders Outer1InnerPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l outer1InnerLocale) FromTag(lang language.Tag, placeholders Outer1InnerPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l outer1InnerLocale) Locale(placeholders Outer1InnerPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type Outer1InnerPlaceholders struct {
	Placeholder1 string
	Placeholder2 string
}

func (p Outer1InnerPlaceholders) en() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

func (p Outer1InnerPlaceholders) fr() []any {
	return []any{p.Placeholder2, p.Placeholder1}
}

type outer2InnerLocale struct{}

func (l outer2InnerLocale) en() string {
	return "nested 2 %s %s"
}

func (l outer2InnerLocale) fr() string {
	return "nested 2 %s %s fr"
}

func (l outer2InnerLocale) parse(lang language.Tag, level int, placeholders Outer2InnerPlaceholders) string {
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

func (l outer2InnerLocale) FromCtx(ctx context.Context, placeholders Outer2InnerPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l outer2InnerLocale) FromString(lang string, placeholders Outer2InnerPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l outer2InnerLocale) FromTag(lang language.Tag, placeholders Outer2InnerPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l outer2InnerLocale) Locale(placeholders Outer2InnerPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type Outer2InnerPlaceholders struct {
	Placeholder1 string
	Placeholder2 string
}

func (p Outer2InnerPlaceholders) en() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

func (p Outer2InnerPlaceholders) fr() []any {
	return []any{p.Placeholder2, p.Placeholder1}
}

type outer3Outer1InnerLocale struct{}

func (l outer3Outer1InnerLocale) en() string {
	return "nested 3 %s %s"
}

func (l outer3Outer1InnerLocale) fr() string {
	return "nested 3 %s %s fr"
}

func (l outer3Outer1InnerLocale) parse(lang language.Tag, level int, placeholders Outer3Outer1InnerPlaceholders) string {
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

func (l outer3Outer1InnerLocale) FromCtx(ctx context.Context, placeholders Outer3Outer1InnerPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l outer3Outer1InnerLocale) FromString(lang string, placeholders Outer3Outer1InnerPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l outer3Outer1InnerLocale) FromTag(lang language.Tag, placeholders Outer3Outer1InnerPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l outer3Outer1InnerLocale) Locale(placeholders Outer3Outer1InnerPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type Outer3Outer1InnerPlaceholders struct {
	Placeholder1 string
	Placeholder2 string
}

func (p Outer3Outer1InnerPlaceholders) en() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

func (p Outer3Outer1InnerPlaceholders) fr() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

type outer4Outer1InnerLocale struct{}

func (l outer4Outer1InnerLocale) en() string {
	return "nested 4 %s %s"
}

func (l outer4Outer1InnerLocale) fr() string {
	return "nested 4 %s %s fr"
}

func (l outer4Outer1InnerLocale) parse(lang language.Tag, level int, placeholders Outer4Outer1InnerPlaceholders) string {
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

func (l outer4Outer1InnerLocale) FromCtx(ctx context.Context, placeholders Outer4Outer1InnerPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l outer4Outer1InnerLocale) FromString(lang string, placeholders Outer4Outer1InnerPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l outer4Outer1InnerLocale) FromTag(lang language.Tag, placeholders Outer4Outer1InnerPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l outer4Outer1InnerLocale) Locale(placeholders Outer4Outer1InnerPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type Outer4Outer1InnerPlaceholders struct {
	Placeholder1 string
	Placeholder2 string
}

func (p Outer4Outer1InnerPlaceholders) en() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

func (p Outer4Outer1InnerPlaceholders) fr() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

type outer6Outer1InnerLocale struct{}

func (l outer6Outer1InnerLocale) en() string {
	return "nested 4 %s %s"
}

func (l outer6Outer1InnerLocale) fr() string {
	return "nested 4 %s %s fr"
}

func (l outer6Outer1InnerLocale) parse(lang language.Tag, level int, placeholders Outer6Outer1InnerPlaceholders) string {
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

func (l outer6Outer1InnerLocale) FromCtx(ctx context.Context, placeholders Outer6Outer1InnerPlaceholders) string {
	lang, _ := lexigo.FromCtx(ctx)
	return l.parse(lang, 0, placeholders)
}

func (l outer6Outer1InnerLocale) FromString(lang string, placeholders Outer6Outer1InnerPlaceholders) string {
	tag, _ := language.Parse(lang)
	return l.parse(tag, 0, placeholders)
}

func (l outer6Outer1InnerLocale) FromTag(lang language.Tag, placeholders Outer6Outer1InnerPlaceholders) string {
	return l.parse(lang, 0, placeholders)
}

func (l outer6Outer1InnerLocale) Locale(placeholders Outer6Outer1InnerPlaceholders) Locale {
	return Locale{locale: l, placeholders: placeholders}
}

type Outer6Outer1InnerPlaceholders struct {
	Placeholder1 string
	Placeholder2 string
}

func (p Outer6Outer1InnerPlaceholders) en() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}

func (p Outer6Outer1InnerPlaceholders) fr() []any {
	return []any{p.Placeholder1, p.Placeholder2}
}
