package main

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "",
	Long:  ``,
}

func init() {
	cobra.OnInitialize(func() {
		if !verbose {
			log.SetOutput(io.Discard)
		}
	})
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
