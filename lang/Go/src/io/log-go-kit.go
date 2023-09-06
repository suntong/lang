package main

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

const esc = "\x1b"
const start = esc + "[1m" + "starting task" + esc + "(B" + esc + "[m"

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	//logger = log.With(logger, "ts", log.DefaultTimestamp)

	type Task struct {
		ID int
	}

	fmt.Println(start)

	RunTask := func(task Task, logger log.Logger) {
		logger.Log("taskID", task.ID, "event", start)

		logger.Log("taskID", task.ID, "event", "task complete")
	}
	RunTask(Task{ID: 1}, logger)

	p := []interface{}{"taskID", "1.5", "event", "task complete"}
	logger.Log(p...)
	logger.Log("taskID", 2, "event", "task complete")
}
