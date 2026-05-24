
package utils

import "github.com/charmbracelet/lipgloss"

type ReaderStyles struct {
    Title lipgloss.Style
    Item  lipgloss.Style
}

func DefaultStyles() ReaderStyles {
    return ReaderStyles {
        Title: lipgloss.NewStyle().
            Background(lipgloss.Color("#957fb8")).
            Bold(true),
        Item: lipgloss.NewStyle().
            Background(lipgloss.Color("#223249")),
    }
}
