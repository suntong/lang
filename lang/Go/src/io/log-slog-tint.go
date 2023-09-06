package main

import (
	"context"
	_ "embed"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

const (
	LevelTrace  = slog.Level(-8)
	LevelDbg3   = slog.Level(-7)
	LevelDbg2   = slog.Level(-6)
	LevelDbg1   = slog.Level(-5)
	LevelNotice = slog.Level(2)
	LevelFatal  = slog.Level(12)
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", p.Name),
		slog.Int("age", p.Age))
}

func main() {

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
		NoColor:    !isatty.IsTerminal(os.Stderr.Fd()),
	})))

	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	tintOptions := tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
	}
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, &tintOptions)))
	slog.Debug("Debug message")

	person := Person{
		Name: "John",
		Age:  30,
	}
	slog.Info("This is the person", slog.Any("person", person))

	// Creating custom log levels
	// https://betterstack.com/community/guides/logging/logging-in-go/#creating-custom-log-levels

	{
		opts := &slog.HandlerOptions{
			Level: LevelTrace,
		}

		logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

		ctx := context.Background()
		logger.Log(ctx, LevelTrace, "Trace message")
		logger.Log(ctx, LevelNotice, "Notice message")
		logger.Log(ctx, LevelFatal, "Fatal level")
	}

	{
		opts := &tint.Options{
			Level:      LevelTrace,
			TimeFormat: time.TimeOnly,
		}

		logger := slog.New(tint.NewHandler(os.Stdout, opts))

		ctx := context.Background()
		logger.Log(ctx, LevelDbg1, "Hidden debug message level1")
		logger.Log(ctx, LevelDbg2, "Hidden debug message level2")
		logger.Log(ctx, LevelDbg3, "Hidden debug message level3")
		logger.Log(ctx, LevelTrace, "Trace message")
		logger.Log(ctx, LevelNotice, "Notice message")
		logger.Log(ctx, LevelFatal, "Fatal level")
	}

}
