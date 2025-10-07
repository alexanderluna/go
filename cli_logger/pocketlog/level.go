package pocketlog

// Level represents an available logging level
type Level byte

const (
	// LevelDebug has information that is diagnostically helpful.
	LevelDebug Level = iota
	// LevelInfo has information that is generally useful to log.
	LevelInfo
	// LevelWarn has information about anything that can cause app oddities.
	LevelWarn
	// LevelError has information about errors which are fatal to the operation.
	LevelError
	// LevelFatal has information about errors that force a shutdown.
	LevelFatal
)
