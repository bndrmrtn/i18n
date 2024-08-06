package i18n

import (
	"encoding/json"
	"encoding/xml"
)

type Unmarshal struct {
	Extensions []string
	Func       func(data []byte, v any) error
}

var DefaultUnmarshaler = []Unmarshal{
	{[]string{"json"}, json.Unmarshal},
	{[]string{"xml"}, xml.Unmarshal},
}
