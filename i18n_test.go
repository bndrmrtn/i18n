package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_I18n_Create(t *testing.T) {
	i18n := New(&Config{
		FallbackLocale: "en",
		LoadDir:        "./lang",
	})

	assert.Equal(t, "en", i18n.Create("en").Locale(), "Default locale should be en")
	assert.Equal(t, "hu", i18n.Create("hu").Locale(), "Fallback locale should be hu")
	assert.Equal(t, "en", i18n.Create("unknown").Locale(), "Unknown locale should be fallbacked to en")
}

func Test_Translation(t *testing.T) {
	i18n := New(&Config{
		FallbackLocale: "en",
		LoadDir:        "./lang",
	})

	translator := i18n.Create("hu")

	data := translator.T("Hi")

	assert.Equal(t, "Szia", data, "Should be translated to hungarian")
}

func Test_UnnamedArgs(t *testing.T) {
	i18n := New(&Config{
		FallbackLocale: "en",
		LoadDir:        "./lang",
	})

	translator := i18n.CreateT("en")

	data := translator("Hey, {0}", "John")

	assert.Equal(t, "Hey, John", data, "Should replace the first argument with John.")
}

func Test_NamedArgs(t *testing.T) {
	i18n := New(&Config{
		FallbackLocale: "en",
		LoadDir:        "./lang",
	})

	translator := i18n.Create("en")

	data := translator.T("I am {name}, {age} years old.", map[string]interface{}{"name": "John", "age": 24})

	assert.Equal(t, "I am John, 24 years old.", data, "Should replace the named arguments with John and his age.")
}

func Test_ConfigMessages(t *testing.T) {
	i18n := New(&Config{
		FallbackLocale: "en",
		Data:           getLanguageData(),
	})

	translator := i18n.Create("en")

	data := translator.T("Hi")

	assert.Equal(t, "Hi", data, "Should be translated to English")

	translator = i18n.Create("hu")

	data = translator.T("Hi")

	assert.Equal(t, "Szia", data, "Should be translated to Hungarian")
}

func Test_BothMessages(t *testing.T) {
	i18n := New(&Config{
		FallbackLocale: "en",
		Data:           getLanguageData(),
		LoadDir:        "./lang",
	})

	translator := i18n.Create("en")

	data := translator.T("Bye")

	assert.Equal(t, "Bye", data, "Should be translated to English")

	data = "What's up?"

	assert.Equal(t, "What's up?", data, "Should be translated to English")

	translator = i18n.Create("hu")

	data = translator.T("Bye")

	assert.Equal(t, "Viszlát", data, "Should be translated to Hungarian")

	data = translator.T("What's up?")

	assert.Equal(t, "Mizu?", data, "Should be translated to Hungarian")
}

func Test_LanguageChange(t *testing.T) {
	i18n := New(&Config{
		FallbackLocale: "en",
		LoadDir:        "./lang",
	})

	translator := i18n.Create("en")
	assert.Equal(t, "Hi", translator.T("Hi"), "Should be translated to English")
	err := translator.ChangeLocale("hu")
	assert.Equal(t, nil, err, "Expecting no errors, because Hungarian is added.")
	assert.Equal(t, "Szia", translator.T("Hi"), "Should be translated to Hungarian")
}

func getLanguageData() []Language {
	return []Language{
		{
			Language: "en",
			Messages: []Message{
				{
					Key:   "Hi",
					Value: "Hi",
				},
				{
					Key:   "Bye",
					Value: "Bye",
				},
			},
		},
		{
			Language: "hu",
			Messages: []Message{
				{
					Key:   "Hi",
					Value: "Szia",
				},
				{
					Key:   "Bye",
					Value: "Viszlát",
				},
			},
		},
	}
}
