package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var greet = &cobra.Command{
	Use:   "greet",
	Short: "count the number of words in a string",
	Run: func(cmd *cobra.Command, args []string) {
		if h, _ := cmd.Flags().GetBool("hello"); h {
			fmt.Println("Hello")
		}
		if g, _ := cmd.Flags().GetBool("greetings"); g {
			fmt.Println("Greetings")
		}
		if w, _ := cmd.Flags().GetBool("welcome"); w {
			fmt.Println("Welcome")
		}
	},
}

var root = &cobra.Command{
	Use: "app",
}

func main() {
	greet.Flags().BoolP("hello", "H", false, "Greets as `Hello`")
	greet.Flags().BoolP("greetings", "g", false, "Greets as `Greetings`")
	greet.Flags().BoolP("welcome", "w", false, "Greets as `Welcome`")

	root.AddCommand(greet)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
