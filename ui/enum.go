package ui

type PromptMode int

// 3 modes from which we select to ask questions
const (
	ExecPromptMode PromptMode = iota
	ConfigPromptMode
	ChatPromptMode
	defaultPromptMode
)

// These are modes while running the output of returned values from GPT
type RunMode int

const (
	CliMode RunMode = iota
	ReplMode
)
