package test

import (
	"context"
	"github.com/MysteriousPotato/lexigo/cmd/lexigogen/test/other"
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
			language.AmericanEnglish,
			MyLocaleStrPlaceholders{
				Placeholder: "test",
			},
		))

		assert(t, "myLocale 10", Locales.MyLocaleInt.FromTag(
			language.AmericanEnglish,
			MyLocaleIntPlaceholders{
				Placeholder: 10,
			},
		))

		assert(t, "myLocale 10.00", Locales.MyLocaleFloat.FromTag(
			language.AmericanEnglish,
			MyLocaleFloatPlaceholders{
				Placeholder: 10.0,
			},
		))

		assert(t, "myLocale <nil>", Locales.MyLocaleAny.FromTag(
			language.AmericanEnglish,
			MyLocaleAnyPlaceholders{
				Placeholder: nil,
			},
		))

		assert(t, "myLocale other", Locales.MyLocaleOther.FromTag(
			language.AmericanEnglish,
			MyLocaleOtherPlaceholders{
				Placeholder: other.Other("other"),
			},
		))
	})

	t.Run("nested", func(t *testing.T) {
		assert(t, "nested 1 test1 10", Locales.Outer1.Inner.FromTag(
			language.AmericanEnglish,
			Outer1InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: "10",
			},
		))
	})

	t.Run("different placeholder order", func(t *testing.T) {
		assert(t, "nested 1 10 test1 fr", Locales.Outer1.Inner.FromTag(
			language.French,
			Outer1InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: "10",
			},
		))
	})

	t.Run("language fallback", func(t *testing.T) {
		assert(t, "nested 1 test1 10", Locales.Outer1.Inner.FromTag(
			language.German,
			Outer1InnerPlaceholders{
				Placeholder1: "test1",
				Placeholder2: "10",
			},
		))
	})

	t.Run("best match", func(t *testing.T) {
		assert(t, "nested 1 test1 10", Locales.Outer1.Inner.FromTag(
			language.BritishEnglish,
			Outer1InnerPlaceholders{
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
			Locales.MyLocale.FromTag(language.BritishEnglish)
		}
	})

	b.Run("simple exact match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.FromTag(language.AmericanEnglish)
		}
	})

	b.Run("simple locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.Locale().FromTag(language.BritishEnglish)
		}
	})

	b.Run("simple locale exact match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.MyLocale.Locale().FromTag(language.AmericanEnglish)
		}
	})

	b.Run("nested with multiple placeholders", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer1.Inner.FromTag(language.BritishEnglish, Outer1InnerPlaceholders{
				Placeholder1: "potato",
				Placeholder2: "test",
			})
		}
	})

	b.Run("nested with multiple placeholders exact match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer1.Inner.FromTag(language.AmericanEnglish, Outer1InnerPlaceholders{
				Placeholder1: "potato",
				Placeholder2: "test",
			})
		}
	})

	b.Run("nested locale with multiple placeholders", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer1.Inner.
				Locale(Outer1InnerPlaceholders{
					Placeholder1: "potato",
					Placeholder2: "test",
				}).
				FromTag(language.BritishEnglish)
		}
	})

	b.Run("nested locale with multiple placeholders exact match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Locales.Outer1.Inner.
				Locale(Outer1InnerPlaceholders{
					Placeholder1: "potato",
					Placeholder2: "test",
				}).
				FromTag(language.AmericanEnglish)
		}
	})
}
