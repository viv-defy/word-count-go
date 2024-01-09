package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hello = &cobra.Command{
	Use:   "hello",
	Short: "count the number of words in a string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

var greet = &cobra.Command{
	Use:   "greet",
	Short: "count the number of words in a string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Greetings")
	},
}

var root = &cobra.Command{
	Use: "app",
}

func main() {
	root.AddCommand(hello)
	root.AddCommand(greet)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
