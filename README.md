# Lexigo

Lexigo is a golang library for generating locales.
It provides translation type safety and support for placeholders

***Disclaimer***
Lexigo is still in the early stage of development, expect breaking changes.

## Installation

#### Codegen: 
```bash
go install github.com/MysteriousPotato/lexigo/cmd/lexigogen@v0.3.1
```

#### Library:
```bash
go get github.com/MysteriousPotato/lexigo/pkg@v0.3.1
```

## Usage

### 1. Define your locales
- Currently only supports json files
- File names must be valid BCP 47 language tags Ex:. en.json, en-US.json
- All locale files must be under the same directory. Lexigo looks up the directory recursively, so you can structure the directory however you like.
- Only one file must have the "default" suffix in its name Ex.: en.default.json. Lexigo will use the default language as a reference for other languages.


#### You can nest locales:
```
{
  "vegetables": {
    "potato": "Potato",
    "zucchini": "Zucchini"
  },
}
```

#### You can define placeholders similarly to templates:
```
{
    "potatoStatement": "I {{.Statement}} potato",
}
```


##### You can also specify a type and a format.

Lexigo uses [fmt](https://pkg.go.dev/fmt) under the hood so refer to it for verbs.
When omitted, the default type and verb are `string` and `%s`
```
{
    "potatoStatement": "I {{.Statement:[]byte(%q)}} potato",
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

#### ***Or***
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
    fmt.Println(myLocale.FromTag(language.German))
	
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

## Roadmap

- Adding currency support
- Adding genderization support
- Adding pluralization support
- Adding `time.Time` support


## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://github.com/MysteriousPotato/lexigo/blob/main/LICENSE)