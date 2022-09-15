package cmd

import (
	"github.com/crerwin/outdoorsy/pkg/outdoorsy"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import data from text file(s)",
	Long:  "Import data from text file(s)",
	Run:   importRun,
}

func importRun(cmd *cobra.Command, args []string) {
	customers := outdoorsy.LoadFiles((args))
	outputCustomers(customers)
}
