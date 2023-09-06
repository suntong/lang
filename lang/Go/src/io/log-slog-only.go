// https://betterstack.com/community/guides/logging/logging-in-go/

package main

import (
	"context"
	"log/slog"
	"os"
)

const (
	LevelTrace  = slog.Level(-8)
	LevelDbg3   = slog.Level(-7)
	LevelDbg2   = slog.Level(-6)
	LevelDbg1   = slog.Level(-5)
	LevelNotice = slog.Level(2)
	LevelFatal  = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelTrace:  "TRC",
	LevelDbg3:   "D-3",
	LevelDbg2:   "D-2",
	LevelDbg1:   "D-1",
	LevelNotice: "NTC",
	LevelFatal:  "FTL",
}

func main() {
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	//

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	logLevel := &slog.LevelVar{} // INFO
	// you can change the level anytime like this
	logLevel.Set(slog.LevelDebug)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	logger.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)

	logger.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)

	// child loggers
	child := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", "buildInfo.GoVersion"),
		),
	)
	// all records created by the child logger will contain the specified attributes under the program_info property
	child.Info("image upload successful", slog.String("image_id", "39ud88"))
	child.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 mb"),
	)

	// custom log levels
	{
		opts := &slog.HandlerOptions{
			Level: LevelTrace,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.LevelKey {
					level := a.Value.Any().(slog.Level)
					levelLabel, exists := LevelNames[level]
					if !exists {
						levelLabel = level.String()
					}

					a.Value = slog.StringValue(levelLabel)
				}

				return a
			},
		}

		logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

		ctx := context.Background()
		logger.Log(ctx, LevelDbg1, "Hidden debug message level1")
		logger.Log(ctx, LevelDbg2, "Hidden debug message level2")
		logger.Log(ctx, LevelDbg3, "Hidden debug message level3")
		logger.Log(ctx, LevelTrace, "Trace message")
		logger.Log(ctx, LevelNotice, "Notice message")
		logger.Log(ctx, LevelFatal, "Fatal level")
	}
}
