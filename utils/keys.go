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
