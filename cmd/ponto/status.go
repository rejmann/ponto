package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "mostra o andamento do dia",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		agora := time.Now()

		d, err := diaAberto(agora)
		if err != nil {
			fmt.Println("nenhuma jornada em andamento.")
			return nil
		}

		imprimirDia(d, agora)

		return nil
	},
}
