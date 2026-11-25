package ui

import "github.com/charmbracelet/bubbles/textinput"

const (
	exec_icon          = "Rocky Bhai >"
	exec_placeholder   = "Worker"
	config_icon        = "LawNOrder >"
	config_placeholder = "Enter your Opnen AI key..."
	chat_icon          = "Chat >"
	chat_placeholder   = "Ask me something..."
)

type Prompt struct {
	mode  PromptMode
	input textinput.Model
}
