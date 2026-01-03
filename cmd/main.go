package main

import (
	"fmt"
	"os"

	"github.com/catmika/mika/internal/prompt"
	"github.com/catmika/mika/internal/ui"
	// "github.com/catmika/mika/internal/scaffold"
)

func main() {
	ui.PrintBanner()

	config, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%t\n", config.GuidanceSwitch)

	// if err := scaffold.Generate(config); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to generate project: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// ui.PrintSuccess(config.Name)
}
