package main

import (
	"os"

	"github.com/faddat/clint-demo/cmd/clint-demod/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
