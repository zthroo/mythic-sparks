package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Get a spark",
	Long:    "Get a spark from the provided Subject and Subsubject",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generation of spark for %s and %s: %s\n\n", args[0], args[1], Generate(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
