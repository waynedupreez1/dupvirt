// Creates a logging interface used by all other packages
// pretty much uses slog for logging
// Author: Wayne du Preez

package logger

import (
    "os"
	"log/slog"
)

// LogLevel allows other packages to use this
type LogLevel int

// Debug Allows for different loglevels to be used
const (
    Debug LogLevel = iota
	Info
	Warn
	Error
)

// ILogger is the logger interface
type ILogger interface {
    Debug(msg string, keysAndValues ...any)
    Info(msg string, keysAndValues ...any)
    Warn(msg string, keysAndValues ...any)
    Error(msg string, keysAndValues ...any)
}

type slogLogger struct {
    logger *slog.Logger
}

// New creates a brand new logger
func New(logLevel LogLevel) ILogger {    

    level := new(slog.LevelVar)
   
    switch logLevel{
    case Debug:
        level.Set(slog.LevelDebug)
    case Info:
        level.Set(slog.LevelInfo)
    case Warn:
        level.Set(slog.LevelWarn)
    case Error:
        level.Set(slog.LevelError)
    }
    
    handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
    })
    
    logger := slogLogger {
        logger: slog.New(handler),
    }
    
    return &logger
}

// Debug allows for logging debug messages
func (t *slogLogger) Debug(msg string, keysAndValues ...any) {

    t.logger.Debug(msg, keysAndValues...)
}

// Info allows for logging info messages
func (t *slogLogger) Info(msg string, keysAndValues ...any) {
    t.logger.Info(msg, keysAndValues...)
}

// Warn allows for logging warn messages
func (t *slogLogger) Warn(msg string, keysAndValues ...any) {
    t.logger.Warn(msg, keysAndValues...)
}

// Error allows for logging error messages
func (t *slogLogger) Error(msg string, keysAndValues ...any) {
    t.logger.Error(msg, keysAndValues...)
}
