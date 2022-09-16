package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	sort    bool
	sortBy  string
	rootCmd = &cobra.Command{
		Use:              "outdoorsy",
		Short:            "Outdoor.sy CLI",
		Long:             "Load and print data or serve the Outdoor.sy API",
		PersistentPreRun: rootRun,
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")

	importCmd.Flags().BoolVarP(&sort, "sort", "s", false, "Sort customer list")
	importCmd.Flags().StringVarP(&sortBy, "sort-by", "b", "name", "Field by which to sort customer list (name|vehicle-type)")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(importCmd)
}

func rootRun(cmd *cobra.Command, args []string) {
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	log.Debug("Debug logging is on.")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
