package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/authit/authit/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "authit"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
