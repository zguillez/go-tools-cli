package main

import (
	"github.com/fatih/color"
	gotools "github.com/zguillez/go-tools"
	"github.com/zguillez/go-tools-cli/core"
	gotoolsc "github.com/zguillez/go-tools/core"
)

func main() {
	version := gotoolsc.FlagBool([]string{"version", "v"}, false, "Package version")
	help := gotoolsc.FlagBool([]string{"help", "h"}, false, "Help")
	command := gotoolsc.FlagString([]string{"command", "c"}, "", "Command to execute:\ntools -c version -l minor")
	input := gotoolsc.FlagString([]string{"input", "i"}, "", "Input value text")
	output := gotoolsc.FlagString([]string{"output", "o"}, "", "Output value text")
	level := gotoolsc.FlagString([]string{"level", "l"}, "100", "Level value int")
	verbose := gotoolsc.FlagBool([]string{"verbose"}, false, "Display console logs")
	gotoolsc.FlagParse()

	if *help {
		core.Help()
		gotools.Help()
	} else if *version {
		gotools.Version()
		color.Cyan("[gtools-cli v0.1.8]")
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
