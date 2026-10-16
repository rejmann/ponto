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
		return voltar(time.Now())
	},
}

func voltar(agora time.Time) error {
	d, err := diaAberto(agora)
	if err != nil {
		return err
	}

	p := d.pausaAtiva()
	if p == nil {
		return errors.New("nenhuma pausa em andamento")
	}
	p.Fim = agora

	fmt.Printf("de volta! pausa (%s) durou %s.\n",
		p.Motivo, fmtDur(agora.Sub(p.Inicio)))

	return salvarDia(d)
}
