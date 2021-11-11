package main

import (
	"github.com/fatih/color"
	gotools "github.com/zguillez/go-tools"
	"github.com/zguillez/go-tools/core"
)

func main() {
	version := core.FlagBool([]string{"version", "v"}, false, "Package version")
	help := core.FlagBool([]string{"help", "h"}, false, "Help")
	command := core.FlagString([]string{"command", "c"}, "", "Command to execute:\ntools -c version -l minor")
	input := core.FlagString([]string{"input", "i"}, "", "Input value text")
	output := core.FlagString([]string{"output", "o"}, "", "Output value text")
	level := core.FlagString([]string{"level", "l"}, "100", "Level value int")
	verbose := core.FlagBool([]string{"verbose"}, false, "Display console logs")
	core.FlagParse()

	if *help {
		gotools.Help()
	} else if *version {
		color.Yellow("v0.1.7-cli")
		gotools.Version()
	} else if *command != "" {
		var args []string
		switch *command {
		case "readfile":
			args = []string{*input}
		case "minimize":
			args = []string{*input, *output, *level}
		case "version":
			args = []string{*level}
		default:
			args = []string{}
		}
		gotools.Command(*command, *verbose, args...)
	} else {
		color.Red("[error: tool -c command empty]")
	}
}
