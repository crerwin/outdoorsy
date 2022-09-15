package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	sort    bool
	sortBy  string
	rootCmd = &cobra.Command{
		Use:   "outdoorsy",
		Short: "Outdoor.sy CLI",
		Long:  "Load and print data or serve the Outdoor.sy API",
	}
)

func init() {
	importCmd.Flags().BoolVarP(&sort, "sort", "s", false, "Sort customer list")
	importCmd.Flags().StringVarP(&sortBy, "sort-by", "b", "name", "Field by which to sort customer list (name|vehicle-type)")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(importCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
