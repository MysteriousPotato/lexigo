package lexigo

import (
	"golang.org/x/text/language"
	"net/http"
)

// LanguageMiddleware inserts the language.Tag into the request context according to "Accept-Language" header
//
// Usage:
//
//	func main() {
//		// Assuming the package name where the generated locales are is "i18n"
//		mux := LanguageMiddleware(i18n.Matcher)(http.DefaultServeMux)
//		//...
//	}
func LanguageMiddleware(matcher *Matcher) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// ignore the error here
			// worst case scenario, no accepted language matches and the language will fall back to the default one.
			acceptedLanguages, _, _ := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
			bestMatch := matcher.MatchAll(acceptedLanguages...)
			ctx := WithLanguage(r.Context(), bestMatch)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
