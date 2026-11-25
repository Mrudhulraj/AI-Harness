package ai

import (
	"github.com/99designs/gqlgen/codegen/config"
)

const (
	noexec string = "[noexec]"
)

type Engine struct {
	model        EngineMode // Enum
	config       *config.Config
	client       *openai.Client
	execMessages []openai.ChatCompletionMessage
	chatMessages []openai.ChatCompletionMessage
	channel      chan EngineChatStreamOutput
	pipe         string
	running      bool
}
