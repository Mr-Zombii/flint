package cli

import "fmt"

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  flint <command> [options]")
	fmt.Println()

	fmt.Println("Commands:")
	for _, cmd := range commands {
		printCmd(cmd.Name, cmd.Description)
	}

	fmt.Println()
	fmt.Println("Use 'flint help <command>' for more details.")
}

func printCmd(name, desc string) {
	fmt.Printf("  %-14s  %s\n", name, desc)
}
