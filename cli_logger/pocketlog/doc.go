/*
Package pocketlog exposes an API to log whatever you want.

Instantiate a logger with pocketlog.New, and giving it a threshold level:
  - LevelDebug
  - LevelInfo
  - LevelWarn
  - LevelError
  - LevelFatal

Messages with less critical logs won't be logged.

The logger can be called to log messages on five levels:
  - Debug: information that is diagnostically helpful
  - Info: information that is generally useful to log
  - Warn: information about anything that can cause app oddities
  - Error: information about errors which are fatal to the operation
  - Fatal: information about errors that force a shutdown
*/
package pocketlog
