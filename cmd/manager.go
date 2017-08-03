package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "transer",
	Short: "Data converter",
	Long: "Multi data converter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run")
	},
}

func Run() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print version",
	Long: "Print transer's version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.1")
	},
}
