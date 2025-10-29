package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Get a spark",
	Long:    "Get a spark from the provided Subject and Subtype",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := Generate(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("Generation of spark for %s and %s: %s\n\n", args[0], args[1], res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
