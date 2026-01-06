package main

import (
	"github.com/kabirnayeem99/TaskTrackerGoLang/cmd/taskcli"
	"os"
)

func main() {
	if err := taskcli.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
