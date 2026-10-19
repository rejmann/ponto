package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "ponto",
	Short:         "controle de ponto e pausas do dia de trabalho",
	SilenceUsage:  true,
	SilenceErrors: false,
}

func init() {
	cmds := []*cobra.Command{
		iniciarCmd,
		encerrarCmd,
		pausaCmd,
		voltarCmd,
		statusCmd,
		historicoCmd,
	}

	for _, c := range cmds {
		rootCmd.AddCommand(c)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
