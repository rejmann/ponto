package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var pausaCmd = &cobra.Command{
	Use:   "pausar [motivo]",
	Short: "inicia uma pausa com descricao (ex: ponto pausa almoco)",
	RunE: func(cmd *cobra.Command, args []string) error {
		agora := time.Now()
		motivo := strings.Join(args, " ")

		d, err := diaAberto(agora)
		if err != nil {
			return err
		}

		if p := d.pausaAtiva(); p != nil {
			return fmt.Errorf("ja existe uma pausa em andamento (%s)", p.Motivo)
		}

		motivo = strings.TrimSpace(motivo)
		if motivo == "" {
			motivo = "pausa"
		}

		d.Pausas = append(d.Pausas, *NewPausa(motivo, agora))
		fmt.Printf("pausa iniciada com descricao (%s) as %s.\n",
			motivo, agora.Format(formatoHora))

		return salvarDia(d)
	},
}
