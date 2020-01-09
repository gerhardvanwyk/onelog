package log

import (
	"os"
	"time"

	log "github.com/gerhardvanwyk/onelogplus"
)

var logger = log.New(os.Stdout, log.FINEST).Hook(func(e log.Entry) {
	e.Int64("time", time.Now().Unix())
})

// Finest logs an entry with FINEST level.
func Finest(msg string) {
	logger.Log(log.FINEST, msg)
}

// FinestWith return an ChainEntry with INFO level.
func FinestWith(msg string) log.ChainEntry {
	return logger.LogWith(log.FINEST, msg)
}

// FinestWithFields logs an entry with INFO level and custom fields.
func FinestWithFields(msg string, fields func(log.Entry)) {
	logger.LogWithFields(log.FINEST, msg, fields)
}

// Finer logs an entry with FINER level.
func Finer(msg string) {
	logger.Log(log.FINER, msg)
}

// FinerWith return an ChainEntry with FINER level.
func FinerWith(msg string) log.ChainEntry {
	return logger.LogWith(log.FINER, msg)
}

// FinerWithFields logs an entry with FINER level and custom fields.
func FinerWithFields(msg string, fields func(log.Entry)) {
	logger.LogWithFields(log.FINER, msg, fields)
}

// Fine logs an entry with FINE level.
func Fine(msg string) {
	logger.Log(log.FINE, msg)
}

// FineWith return an ChainEntry with FINE level.
func FineWith(msg string) log.ChainEntry {
	return logger.LogWith(log.FINE, msg)
}

// FineWithFields logs an entry with FINE level and custom fields.
func FineWithFields(msg string, fields func(log.Entry)) {
	logger.LogWithFields(log.FINE, msg, fields)
}

// Config logs an entry with CONFIG level.
func Config(msg string) {
	logger.Log(log.CONFIG, msg)
}

// FineWith return an ChainEntry with CONFIG level.
func ConfigWith(msg string) log.ChainEntry {
	return logger.LogWith(log.CONFIG, msg)
}

// ConfigWithFields logs an entry with CONFIG level and custom fields.
func ConfigWithFields(msg string, fields func(log.Entry)) {
	logger.LogWithFields(log.CONFIG, msg, fields)
}

// Info prints a message with log level Info.
func Info(msg string) {
	logger.Info(msg)
}

// InfoWith return an ChainEntry with INFO level.
func InfoWith(msg string) log.ChainEntry {
	return logger.LogWith(log.INFO, msg)
}

// InfoWithFields prints a message with log level INFO and fields.
func InfoWithFields(msg string, fields func(e log.Entry)) {
	logger.InfoWithFields(msg, fields)
}

// Debug prints a message with log level DEBUG.
func Debug(msg string) {
	logger.Debug(msg)
}

// DebugWith return ChainEntry with DEBUG level.
func DebugWith(msg string) log.ChainEntry {
	return logger.LogWith(log.DEBUG, msg)
}

// DebugWithFields prints a message with log level DEBUG and fields.
func DebugWithFields(msg string, fields func(e log.Entry)) {
	logger.DebugWithFields(msg, fields)
}

// Warn prints a message with log level INFO.
func Warn(msg string) {
	logger.Warn(msg)
}

// WarnE prints a message with log level WARN and error.
func WarnE(msg string, err error) {
	logger.WarnWithFields(msg, func(e log.Entry) {
		e.Err("ERROR", err)
	})
}

// WarnWith returns a ChainEntry with WARN level
func WarnWith(msg string) log.ChainEntry {
	return logger.LogWith(log.WARN, msg)
}

// WarnWithFields prints a message with log level WARN and fields.
func WarnWithFields(msg string, fields func(e log.Entry)) {
	logger.WarnWithFields(msg, fields)
}

// Error prints a message with log level ERROR.
func Error(msg string) {
	logger.Error(msg)
}

// ErrorE prints a message with log level ERROR and error.
func ErrorE(msg string, err error) {
	logger.ErrorWithFields(msg, func(e log.Entry) {
		e.Err("ERROR", err)
	})
}

// ErrorWith returns a ChainEntry with ERROR level.
func ErrorWith(msg string) log.ChainEntry {
	return logger.LogWith(log.ERROR, msg)
}

// ErrorWithFields prints a message with log level ERROR and fields.
func ErrorWithFields(msg string, fields func(e log.Entry)) {
	logger.ErrorWithFields(msg, fields)
}

// Severe logs an entry with SEVERE level.
func Severe(msg string) {
	logger.Log(log.SEVERE, msg)
}

// SevereE prints a message with log level SEVERE and error.
func SevereE(msg string, err error) {
	logger.SevereWithFields(msg, func(e log.Entry) {
		e.Err("ERROR", err)
	})
}

// SevereWith returns a ChainEntry with SEVERE level.
func SevereWith(msg string) log.ChainEntry {
	return logger.LogWith(log.SEVERE, msg)
}

// SevereWithFields logs an entry with SEVERE level and custom fields.
func SevereWithFields(msg string, fields func(log.Entry)) {
	logger.LogWithFields(log.SEVERE, msg, fields)
}

// Fatal prints a message with log level FATAL.
func Fatal(msg string) {
	logger.Fatal(msg)
}

// FatalE prints a message with log level FATAL and error.
func FatalE(msg string, err error) {
	logger.FatalWithFields(msg, func(e log.Entry) {
		e.Err("ERROR", err)
	})
}

// FatalWith returns a ChainEntry with FATAL level.
func FatalWith(msg string) log.ChainEntry {
	return logger.LogWith(log.FATAL, msg)
}

// FatalWithFields prints a message with log level FATAL and fields.
func FatalWithFields(msg string, fields func(e log.Entry)) {
	logger.FatalWithFields(msg, fields)
}
