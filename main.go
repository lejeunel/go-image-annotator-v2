package main

import (
	"fmt"
	cli "github.com/lejeunel/go-image-annotator-v2/adapters/cli"
	cmd "github.com/lejeunel/go-image-annotator-v2/cmd"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-image-annotator",
	Short: "Image annotation platform",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmd.ServeCmd)
	rootCmd.AddCommand(cli.IngestDirectoryCmd)
}

func main() {
	Execute()
}
