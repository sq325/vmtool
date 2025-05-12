package main

import (
	"fmt"

	"github.com/sq325/vmtool/cmd"
)

func main() {
	// Execute the root command
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
