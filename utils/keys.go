package utils

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
    }
}
