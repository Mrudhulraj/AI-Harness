package config

const (
	openai_key         = ""
	openai_model       = ""
	openai_proxy       = ""
	openai_temperature = ""
	openai_max_tokens  = iota
)

type AiConfig struct {
	key         string
	model       string
	proxy       string
	temperature float64
	maxTokens   int64
}
