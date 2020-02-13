package log

import (
	"context"
	"github.com/rs/zerolog"
	"io"
	"log"
	"os"
)
var ZLogger Logger

type Logger = zerolog.Logger

func New()  {
	ZLogger = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) Logger {
	return ZLogger.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return ZLogger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) Logger {
	return ZLogger.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) Logger {
	return ZLogger.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) Logger {
	return ZLogger.Hook(h)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return ZLogger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return ZLogger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return ZLogger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return ZLogger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return ZLogger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return ZLogger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return ZLogger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return ZLogger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	ZLogger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	ZLogger.Printf(format, v...)
}

// Ctx returns the logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *Logger {
	return zerolog.Ctx(ctx)
}

func StdLogger() *log.Logger {
	return log.New(ZLogger, "", 0)
}



