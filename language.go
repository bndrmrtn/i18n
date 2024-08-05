package i18n

type Language struct {
	Language string
	Messages []Message
}

type Message struct {
	Key   string
	Value string
}