package utils

import "github.com/charmbracelet/glamour"

func GetDefaultRenderer() *glamour.TermRenderer {
	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		glamour.WithWordWrap(80),
		glamour.WithEmoji(),
		glamour.WithPreservedNewLines(),
	)

	return r
}
