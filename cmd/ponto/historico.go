package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var historicoCmd = &cobra.Command{
	Use:   "historico <AAAA-MM-DD>",
	Short: "mostra o registro de um dia",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return historico(args[0])
	},
}

func historico(data string) error {
	if _, err := time.Parse(formatoData, data); err != nil {
		return fmt.Errorf("data invalida %q, use o formato AAAA-MM-DD", data)
	}
	h, err := carregarDia(data)
	if err != nil {
		return err
	}
	if h == nil {
		fmt.Printf("nenhum registro para %s.\n", data)
		return nil
	}

	imprimirDia(h, time.Now())

	return nil
}
