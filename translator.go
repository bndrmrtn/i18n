package i18n

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type Translator interface {
	// Locale returns the current locale.
	Locale() string
	ChangeLocale(locale string) error

	// Translate translates the text to the current locale.
	Translate(key string, value ...interface{}) string
	// T is a shortcut for Translate. It uses the current locale.
	T(key string, value ...interface{}) string

	// Languages returns the list of languages
	Languages() []string

	// New creates a new copy of the current translator
	New() Translator
}

type T struct {
	i18n   *I18n
	locale string
}

func (t *T) New() Translator {
	return &T{
		i18n:   t.i18n,
		locale: t.locale,
	}
}

func (t *T) Locale() string {
	return t.locale
}

func (t *T) ChangeLocale(locale string) error {
	if !slices.Contains(t.i18n.languages, locale) {
		return errors.New("language " + locale + " not found")
	}
	t.locale = locale
	return nil
}

func (t *T) Translate(key string, value ...interface{}) string {
	data, ok := t.i18n.data[t.locale][key]
	if !ok {
		return t.parseMessage(key, value...)
	}

	return t.parseMessage(data, value...)
}

func (t *T) T(key string, value ...interface{}) string {
	return t.Translate(key, value...)
}

func (t *T) Languages() []string {
	return t.i18n.languages
}

func (t *T) parseMessage(message string, value ...interface{}) string {
	if len(value) == 0 {
		return message
	}

	if len(value) == 1 {
		vT := reflect.TypeOf(value[0])

		if vT.Kind() == reflect.Map {
			return t.parseMap(message, value[0].(map[string]interface{}))
		}
	}

	for i, v := range value {
		message = strings.ReplaceAll(message, fmt.Sprintf("{%d}", i), fmt.Sprint(v))
	}

	return message
}

func (t *T) parseMap(message string, values map[string]interface{}) string {
	for k, v := range values {
		message = strings.ReplaceAll(message, fmt.Sprintf("{%v}", k), fmt.Sprint(v))
	}
	return message
}
