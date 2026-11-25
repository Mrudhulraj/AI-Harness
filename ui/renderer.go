package ui

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	exec_color    = "#ffa657"
	config_color  = "#ffffff"
	chat_color    = "#66b3ff"
	help_color    = "#aaaaaa"
	error_color   = "#cc3333"
	warning_color = "#ffcc00"
	success_color = "#46b946"
)

type Renderer struct {
	contentRenderer *glamour.TermRenderer
	successRenderer lipgloss.Style
	warningRenderer lipgloss.Style
	errorRenderer   lipgloss.Style
	helpRenderer    lipgloss.Style
}
