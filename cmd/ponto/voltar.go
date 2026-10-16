package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var voltarCmd = &cobra.Command{
	Use:   "voltar",
	Short: "encerra a pausa atual",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		agora := time.Now()

		d, err := diaAberto(agora)
		if err != nil {
			return err
		}

		p := d.pausaAtiva()
		if p == nil {
			return errors.New("nenhuma pausa em andamento")
		}
		p.Fim = agora

		detalhe := ""
		if p.Motivo != "" {
			detalhe = fmt.Sprintf(" (%s)", p.Motivo)
		}

		fmt.Printf("de volta! pausa%s durou %s.\n",
			detalhe, fmtDur(agora.Sub(p.Inicio)))

		return salvarDia(d)
	},
}
