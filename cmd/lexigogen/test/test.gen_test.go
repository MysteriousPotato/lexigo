package test

import (
	"context"
	"github.com/MysteriousPotato/lexigo/pkg"
	"golang.org/x/text/language"
	"testing"
)

func assert(t *testing.T, expected, got string) {
	if expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestLocales(t *testing.T) {
	t.Run("no placeholder", func(t *testing.T) {
		assert(t, "myLocale %", Locales.MyLocale.FromTag(language.English))
	})

	t.Run("with placeholder", func(t *testing.T) {
		assert(t, "myLocale % \"test\"", Locales.MyLocaleStr.FromTag(
			language.English,
			MyLocaleStrPlaceholders{
				Placeholder: "test",
			},
		))

		assert(t, "myLocale 10", Locales.MyLocaleInt.FromTag(
			language.English,
			MyLocaleIntPlaceholders{
				Placeholder: 10,
			},
		))

		assert(t, "myLocale 10.00", Locales.MyLocaleFloat.FromTag(
			language.English,
			MyLocaleFloatPlaceholders{
				Placeholder: 10.0,
			},
		))

		assert(t, "myLocale <nil>", Locales.MyLocaleAny.FromTag(
			language.English,
			MyLocaleAnyPlaceholders{
				Placeholder: nil,
			},
		))
	})

	t.Run("nested", func(t *testing.T) {
		assert(t, "nested test1 10", Locales.Outer.Inner.FromTag(
			language.English,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: "10",
			},
		))
	})

	t.Run("different placeholder order", func(t *testing.T) {
		assert(t, "nested 10 test1 fr", Locales.Outer.Inner.FromTag(
			language.French,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: "10",
			},
		))
	})

	t.Run("language fallback", func(t *testing.T) {
		assert(t, "nested test1 10", Locales.Outer.Inner.FromTag(
			language.German,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: "10",
			},
		))
	})

	t.Run("best match", func(t *testing.T) {
		assert(t, "nested test1 10", Locales.Outer.Inner.FromTag(
			language.BritishEnglish,
			InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: "10",
			},
		))
	})

	t.Run("from ctx", func(t *testing.T) {
		ctx := lexigo.WithLanguage(context.Background(), language.French)
		assert(t, "myLocale % fr", Locales.MyLocale.FromCtx(ctx))
	})

	t.Run("from string", func(t *testing.T) {
		assert(t, "myLocale % fr", Locales.MyLocale.FromString("fr"))
	})

	t.Run("using Locale", func(t *testing.T) {
		assert(t, "myLocale % fr", Locales.MyLocale.Locale().FromTag(language.French))
	})

	t.Run("using Locale with placeholders", func(t *testing.T) {
		assert(t, "myLocale % \"test\" fr", Locales.MyLocaleStr.
			Locale(MyLocaleStrPlaceholders{Placeholder: "test"}).
			FromTag(language.French),
		)
	})
}

func BenchmarkLocales(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.FromTag(language.English)
		}
	})

	b.Run("simple locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.Locale().FromTag(language.English)
		}
	})

	b.Run("simple with matcher", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.FromTag(language.AmericanEnglish)
		}
	})

	b.Run("simple locale with matcher", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.Locale().FromTag(language.AmericanEnglish)
		}
	})

	b.Run("nested with multiple placeholders", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer.Inner.FromTag(language.English, InnerPlaceholders{
				Placeholder1: "potato",
				Placeholder2: "test",
			})
		}
	})

	b.Run("nested locale with multiple placeholders", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer.Inner.
				Locale(InnerPlaceholders{
					Placeholder1: "potato",
					Placeholder2: "test",
				}).
				FromTag(language.English)
		}
	})

	b.Run("nested with multiple placeholders and matcher", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer.Inner.FromTag(language.AmericanEnglish, InnerPlaceholders{
				Placeholder1: "potato",
				Placeholder2: "test",
			})
		}
	})

	b.Run("nested with multiple placeholders and matcher", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer.Inner.
				Locale(InnerPlaceholders{
					Placeholder1: "potato",
					Placeholder2: "test",
				}).
				FromTag(language.AmericanEnglish)
		}
	})
}
