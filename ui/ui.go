package ui

import (
	"github.com/mrudhulraj/ai-harness/ai"
	"github.com/mrudhulraj/ai-harness/config"
	"github.com/mrudhulraj/ai-harness/history"
)

type UiDimensions struct {
	width  int
	height int
}
type UiState struct {
	error       error
	runMode     RunMode
	promptMode  PromptMode
	configuring bool
	querying    bool
	confirming  bool
	executing   bool
	args        string
	pipe        string
	buffer      string
	command     string
}

type UiComponents struct {
	prompt   *Prompt
	renderer *Renderer
	spinner  *Spinner
}

type UiDimensions struct {
}

type Ui struct {
	state      UiState
	dimensions UiDimensions
	components UiComponents
	config     *config.Config
	engine     *ai.Engine
	history    *history.History
}
