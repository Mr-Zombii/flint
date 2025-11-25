package cli

import (
	"flag"
	"fmt"
)

func RunCLI(args []string) {
	if len(args) < 2 {
		printHelp()
		return
	}

	cmdName := args[1]

	for _, cmd := range commands {
		if cmd.Name == cmdName {
			fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
			cmd.Run(fs)
			return
		}
	}

	fmt.Printf("Unknown command '%s'\n\n", cmdName)
	printHelp()
}
