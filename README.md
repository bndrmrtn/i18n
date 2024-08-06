# I18n for Go

Make translations easier and better with I18n.

## Installation

```shell
go get github.com/bndrmrtn/i18n
```

## Usage

### Configuration

Setup the global translator.

```go
package main

import "github.com/bndrmrtn/i18n"

func main() {
    i := i18n.New(i18n.Config{
        FallbackLocale: "en", // Fallback locale if the desired language is not in the supported languages.
        Data:  []Language{}, // List of language data
        LoadDir: "./lang", // A directory containing the language files.
        Unmarshallers: []Unmarshaller{}, // Unmarshallers for the language files. By default JSON is supported.
    })
}
```

Create a new translator.

```go
t := i.Create("en")
fmt.Println(t.Translate("Hi"))
// or use T as a shortcut
fmt.Println(t.T("Hi"))
```

Get the translator's current language.

```go
fmt.Println(t.Locale()) // en
```

Changing translator's language.

```go
// Here we return an error instead of using the fallback locale.
err := t.ChangeLocale("unknown")
if err != nil {
    fmt.Println(err) // language unknown not found
}
```

Only use the translation function.

```go
t := i.CreateT("hu")
fmt.Println(t("Hi")) // Szia
```

### Messages & Translations

Simple messages.

```go
fmt.Println(t("Hi")) // Szia
```

Unnamed arguments.

```go
fmt.Println(t("Hi {0}", "Martin")) // Szia Martin
fmt.Println(t("Hi {0}, I'm {1}", "John", "Martin")) // Szia John, Martin vagyok
```

Named arguments.

```go
fmt.Println(t("Hi {name}", map[string]string{"name": "John"})) // Szia John
fmt.Println(t("Hi, I am {name} and I am {age} years old.", map[string]interface{}{"name": "Martin", "age": 18})) // Szia, Martin vagyok, 18 éves
```

## Support ❤️

- [PayPal](https://www.paypal.me/instasiteshu)
- [Ko-Fi](https://ko-fi.com/bndrmrtn)
