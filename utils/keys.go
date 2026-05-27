package utils

import tea "github.com/charmbracelet/bubbletea"

type Action int

const (
    ActionNone Action = iota
    ActionUp
    ActionDown
    ActionPageUp
    ActionPageDown
    ActionHome
    ActionEnd
    ActionQuit
    ActionLeft
    ActionRight
)

type KeyMap map[string]Action

func DefaultKeyMap() KeyMap {
    return KeyMap{
        "up":      ActionUp,
        "down":    ActionDown,
        "pgup":    ActionPageUp,
        "pgdown":  ActionPageDown,
        "home":    ActionHome,
        "end":     ActionEnd,
        "q":       ActionQuit,
        "ctrl+c":  ActionQuit,
        "left":    ActionLeft,
        "right":   ActionRight,
    }
}

func (km KeyMap) Lookup(msg tea.KeyMsg) Action {
    if action, ok := km[msg.String()]; ok {
        return action
    }
    return ActionNone
}
