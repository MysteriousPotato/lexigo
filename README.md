# Lexigo

Lexigo is a golang library for generating locales.
It provides translation type safety and support for placeholders

***Disclaimer***
Lexigo is still in the early stage of development, expect breaking changes.

## Installation

#### Codegen: 
```bash
go install github.com/MysteriousPotato/lexigo/cmd/lexigogen@v0.2.0
```

#### Library:
```bash
go get github.com/MysteriousPotato/lexigo/pkg@v0.2.0
```

## Usage

### 1. Define your locales
- Currently only supports json files
- File names must be valid language codes Ex:. en.json, en_US.json
- All locale files must be under the same directory. Lexigo looks up the directory recursively, so you can structure the directory however you like.
- Only one file must have the "default" suffix in its name Ex.: en.default.json. Lexigo will use the default language as a reference for other languages.

#### ./src/en.default.json:
```
{
  // You can nest locales
  "vegetables": {
    "potato": "Potato",
    "zucchini": "Zucchini"
  },
  // You can define placeholders similarly to templates
  // You can also specify a format, in which case lexigo will infer the type if possible.
  // Lexigo uses "fmt" under the hood so specifiers will result in the same format.
  //
  // Currently supported formats and their type:
  //  - %v: any
  //  - %s, %q: string
  //  - %d, %b, %o, %x, %X: int64
  //  - %f, %g, %e: float64
  //
  // The default specifier is "%s"
  "potatoStatement": "I {{.Statement:%s}} potato"
}
```

### 2. Run code generator
```
  -h    Help
  -o string
        Output path
  -p string
        Source path (Directory containing locale files)
  -pkg string
        Package name
  -var string
        Name of the variable containing the locales (default "Locales")
```

The `-pkg`, `-p` and `-o` arguments are required.

Ex.:
```bash
 lexigogen -p ./src -o ./mypkg/locales.gen.go -pkg mypkg 
 ```

### Or
Using go generate:
```go
package mypkg

//go:generate lexigogen -p ./src -o ./locales.gen.go -pkg mypkg
```

### 3. Use the generated code
```go
import 	(
	"golang.org/x/text/language"
	"https://github.com/MysteriousPotato/lexigo/pkg"
)

func main() {
    // Prints "Potato"
    fmt.Println(mypkg.Locales.Vegetables.Potato.FromTag(language.English))

    // Prints "I hate potato"
    fmt.Println(mypkg.Locales.PotatoStatement.FromTag(
        language.English,
        mypkg.PotatoStatementPlaceholders{
            Statement: "hate",
        },
    ))

    // You can create Locale instances for parsing the same locale using different languages
    myLocale := mypkg.Locales.PotatoStatement.Locale(mypkg.PotatoStatementPlaceholders{
        Statement: "hate",
    })
    fmt.Println(myLocale.FromTag(language.English))
    fmt.Println(myLocale.FromTag(language.German)
	
    // You can manually inject the language into the context for using "New"
    ctx := lexigo.WithLanguage(ctx, language.English) 
    fmt.Println(mypkg.Locales.Vegetables.Potato.FromCtx(ctx))
	
    // Or use the built-in middleware "LanguageMiddleware"
    http.DefaultServeMux.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(mypkg.Locales.Vegetables.Potato.New(r.Context()))	
    })
    // Use the language.Matcher generated along with the locales
    mux := lexigo.LanguageMiddleware(mypkg.Matcher)(http.DefaultServeMux)
    http.ListenAndServe("", mux)
}
```


## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://github.com/MysteriousPotato/lexigo/LICENSE)