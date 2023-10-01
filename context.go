package lexigo

import (
	"context"
	"golang.org/x/text/language"
)

const langCtxKey ctxKey = "lang"

type ctxKey string

// WithLanguage injects the (language tag)[language.Tag] into the context
//
// If the provided value is not a valid language according to [language.Parse], an empty value with be injected,
// in which case creating a locale with the ctx will result in the default tag being used.
func WithLanguage(ctx context.Context, lang language.Tag) context.Context {
	return context.WithValue(ctx, langCtxKey, lang)
}

// WithLanguageString injects the (language tag)[language.Tag] corresponding to the given string into the context
//
// Refer to [WithLanguage] for more info.
func WithLanguageString(ctx context.Context, lang string) (context.Context, error) {
	langTag, err := language.Parse(lang)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, langCtxKey, langTag), nil
}

// FromCtx retrieves the language tag from the context
func FromCtx(ctx context.Context) (language.Tag, bool) {
	lang, ok := ctx.Value(langCtxKey).(language.Tag)
	return lang, ok
}
