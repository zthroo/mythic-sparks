package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// main entry point and root of command tree
var rootCmd = &cobra.Command{
	Use:   "spark",
	Short: "spark is a cli tool for generating random sparks for fantasy RPGs",
	Long:  "spark is a cli tool for generating random sparks for fantasy RPGs using the tables provided in Mythic Bastionland by Chris McDowall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// handles execution of the root command rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occured while executing spark '%s'\n", err)
		os.Exit(1)
	}
}
