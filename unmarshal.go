package i18n

import "encoding/json"

type Unmarshal struct {
	Extensions []string
	Func       func(data []byte, v any) error
}

var DefaultUnmarshaler = []Unmarshal{
	{[]string{"json"}, json.Unmarshal},
}
