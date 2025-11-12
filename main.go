package main

import (
	"fmt"

	"github.com/matusstafura/coolsnippets/internal/cli"
)

func main() {
	cfg := cli.MustParse()

	err := cli.ExecuteCommand(cfg)
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}
