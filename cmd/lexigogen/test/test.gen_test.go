package test

import (
	"context"
	"github.com/MysteriousPotato/lexigo/pkg"
	"golang.org/x/text/language"
	"testing"
)

func assert(t *testing.T, expected, got string) {
	if expected != got {
		t.Errorf("expected '%s', got '%s'", expected, got)
	}
}

func TestLocales(t *testing.T) {
	t.Run("no placeholder", func(t *testing.T) {
		assert(t, "myLocale", Locales.MyLocale.NewFromTag(language.English))
	})

	t.Run("with placeholder", func(t *testing.T) {
		assert(t, "myLocale test", Locales.MyLocaleStr.NewFromTag(
			language.English,
			MyLocaleStrPlaceholders{
				Placeholder: "test",
			},
		))

		assert(t, "myLocale 10", Locales.MyLocaleInt.NewFromTag(
			language.English,
			MyLocaleIntPlaceholders{
				Placeholder: 10,
			},
		))

		assert(t, "myLocale 10.000000", Locales.MyLocaleFloat.NewFromTag(
			language.English,
			MyLocaleFloatPlaceholders{
				Placeholder: 10.0,
			},
		))

		assert(t, "myLocale <nil>", Locales.MyLocaleAny.NewFromTag(
			language.English,
			MyLocaleAnyPlaceholders{
				Placeholder: nil,
			},
		))
	})

	t.Run("nested", func(t *testing.T) {
		assert(t, "nested test1 10", Locales.Outer.Inner.NewFromTag(
			language.English,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: 10,
			},
		))
	})

	t.Run("different placeholder order", func(t *testing.T) {
		assert(t, "nested 10 test1 fr", Locales.Outer.Inner.NewFromTag(
			language.French,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: 10,
			},
		))
	})

	t.Run("language fallback", func(t *testing.T) {
		assert(t, "nested test1 10", Locales.Outer.Inner.NewFromTag(
			language.German,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: 10,
			},
		))
	})

	t.Run("best match", func(t *testing.T) {
		assert(t, "nested test1 10", Locales.Outer.Inner.NewFromTag(
			language.BritishEnglish,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: 10,
			},
		))
	})

	t.Run("from ctx", func(t *testing.T) {
		ctx := lexigo.WithLanguage(context.Background(), language.French)
		assert(t, "myLocale fr", Locales.MyLocale.New(ctx))
	})

	t.Run("from string", func(t *testing.T) {
		assert(t, "myLocale fr", Locales.MyLocale.NewFromString("fr"))
	})
}

func BenchmarkLocales(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.NewFromTag(language.English)
		}
	})

	b.Run("simple with matcher", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.NewFromTag(language.AmericanEnglish)
		}
	})

	b.Run("nested with multiple placeholders", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer.Inner.NewFromTag(language.English, InnerPlaceholders{
				Placeholder1: "potato",
				Placeholder2: "test",
			})
		}
	})

	b.Run("nested with multiple placeholders and matcher", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer.Inner.NewFromTag(language.AmericanEnglish, InnerPlaceholders{
				Placeholder1: "potato",
				Placeholder2: "test",
			})
		}
	})
}
