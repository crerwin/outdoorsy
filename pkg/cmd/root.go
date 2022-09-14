package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "outdoorsy",
	Short: "Outdoor.sy CLI",
	Long:  "Load and print data or serve the Outdoor.sy API",
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(importCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
