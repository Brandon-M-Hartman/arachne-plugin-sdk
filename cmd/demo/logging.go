package main

import (
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger initializes the logger with the provided log level
func InitLogger(LogLevel string, LogFilePath string, LogFileMaxMB int, LogFileMaxBackups int, LogFileMaxAge int, LogFileCompress bool) {

	// Declare a variable of type slog.Leveler
	var logLevel slog.Leveler

	// Determine the log level based on the provided string
	switch LogLevel {
	case "Debug":
		logLevel = slog.LevelDebug
	case "Warn":
		logLevel = slog.LevelWarn
	case "Error":
		logLevel = slog.LevelError
	case "Info":
		logLevel = slog.LevelInfo
	default:
		// Default to Info level if no matching level is found
		logLevel = slog.LevelInfo
	}

	// Create a new HandlerOptions with the determined log level and source addition
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}

	// Configure rolling logs
	logFile := &lumberjack.Logger{
		Filename:   LogFilePath,
		MaxSize:    LogFileMaxMB, // megabytes
		MaxBackups: LogFileMaxBackups,
		MaxAge:     LogFileMaxAge,   //days
		Compress:   LogFileCompress, // disabled by default
	}

	// Create a new JSONHandler with the options
	handler := slog.NewJSONHandler(logFile, opts)

	// If the application environment is production, create a new HandlerOptions without source addition
	if AppEnv == "production" {
		opts = &slog.HandlerOptions{
			AddSource: false,
			Level:     logLevel,
		}

		handler = slog.NewJSONHandler(logFile, opts)
	}

	// Create a new logger with the handler
	logger := slog.New(handler)
	// Set the created logger as the default logger
	slog.SetDefault(logger)

}
