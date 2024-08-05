package i18n

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/bndrmrtn/i18n/internal/utils"
)

type I18n struct {
	fallback  string
	languages []string
	data      map[string]map[string]string
}

type Config struct {
	// FallbackLocale is the default language to use.
	FallbackLocale string

	// Data contains the translation data. (You can leave this empty)
	Data []Language

	// LoadDir is a path to the directory containing language files. (You can leave empty)
	LoadDir string
	// Unmarshal is a list of functions that enables to use custom  file types for translation.
	// (It only required if loadDir is not empty.)
	// By default, only JSON is supported.
	Unmarshallers []Unmarshal
}

func New(config *Config) *I18n {
	i := new(I18n)
	i.fallback = config.FallbackLocale
	i.languages = []string{}
	i.data = make(map[string]map[string]string)

	if len(config.Data) > 0 {
		i.loadLanguageData(config.Data)
	}

	if config.LoadDir != "" {
		if len(config.Unmarshallers) == 0 {
			config.Unmarshallers = DefaultUnmarshaler
		}
		i.loadLanguageFiles(config.LoadDir, config.Unmarshallers)
	}

	return i
}

func (i *I18n) Create(language string) Translator {
	if !slices.Contains(i.languages, language) {
		language = i.fallback
	}

	return &T{
		i18n:   i,
		locale: language,
	}
}

func (i *I18n) CreateT(language string) func(key string, values ...interface{}) string {
	t := i.Create(language)
	return t.Translate
}

func (i *I18n) loadLanguageData(langs []Language) {
	for _, data := range langs {
		if !slices.Contains(i.languages, data.Language) {
			i.languages = append(i.languages, data.Language)
		}

		_, ok := i.data[data.Language]
		if !ok {
			i.data[data.Language] = make(map[string]string)
		}
		for _, msg := range data.Messages {
			i.data[data.Language][msg.Key] = msg.Value
		}
	}
}

func (i *I18n) loadLanguageFiles(dir string, unmarshallers []Unmarshal) {
	f, err := utils.WalkDir(dir)
	if err != nil {
		log.Fatal("Failed to get files from dir: ", dir)
	}

	for _, file := range f {
		parts := strings.Split(file, "/")
		fileName := parts[len(parts)-1]
		extension := strings.TrimPrefix(filepath.Ext(fileName), ".")

		// getting the language code from the file names [en, hu, ...]
		languageKey := strings.TrimSuffix(fileName, "."+extension)

		st, err := os.Stat(file)
		if err != nil {
			fmt.Println("Failed to stat file: ", file)
		}

		if st.IsDir() {
			continue
		}

		b, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Failed to load language: %v\n", languageKey)
			continue
		}

		var data = make(map[string]string)

		unmarshaller, err := i.getUnmarshaller(unmarshallers, extension)
		if err != nil {
			fmt.Printf("No unmarshaller found for file: %v", extension)
		}

		err = unmarshaller(b, &data)
		if err != nil {
			fmt.Printf("Failed to unmarshal file for language: %v\n", languageKey)
			continue
		}

		_, ok := i.data[languageKey]
		if !ok {
			i.data[languageKey] = make(map[string]string)
		}

		for key, val := range data {
			i.data[languageKey][key] = val
		}

		if !slices.Contains(i.languages, languageKey) {
			i.languages = append(i.languages, languageKey)
		}
	}
}

func (i *I18n) getUnmarshaller(unmarshallers []Unmarshal, ext string) (func(data []byte, v any) error, error) {
	for _, u := range unmarshallers {
		if slices.Contains(u.Extensions, ext) {
			return u.Func, nil
		}
	}
	return nil, errors.New("missing unmarshaller")
}
