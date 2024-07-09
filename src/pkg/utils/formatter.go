package utils

import (
	"fmt"

	"github.com/fatih/color"
)

type Options struct {
	Label string
}

func PrintLog(title string, message string, options ...Options) string {
	var label string

	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	logTag := green("[server]:")

	if options == nil {
		label = blue(title)
	}

	for _, opt := range options {
		if opt.Label == "success" {
			label = green(title)
		} else if opt.Label == "warning" {
			label = yellow(title)
		} else if opt.Label == "error" {
			label = red(title)
		} else {
			label = blue(title)
		}
	}

	logMessage := cyan(message)
	result := fmt.Sprintf("%s %s %s", logTag, label, logMessage)

	return result
}
