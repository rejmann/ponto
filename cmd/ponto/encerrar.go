package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

const formatoHora = "15:04"

var errNaoIniciada = errors.New("nenhuma jornada em andamento. Use 'ponto iniciar'")

var encerrarCmd = &cobra.Command{
	Use:   "encerrar",
	Short: "mostrar o andamento do dia",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return encerrar(time.Now())
	},
}

func encerrar(agora time.Time) error {
	d, err := diaAberto(agora)
	if err != nil {
		return err
	}
	if p := d.pausaAtiva(); p != nil {
		p.Fim = agora
	}

	d.Fim = agora

	return salvarDia(d)
}

func diaAberto(agora time.Time) (*Dia, error) {
	d, err := carregarDia(agora.Format(formatoData))
	if err != nil {
		return nil, err
	}

	if d == nil || !d.Fim.IsZero() {
		return nil, errNaoIniciada
	}

	imprimirDia(d, agora)

	return d, nil
}

func imprimirDia(d *Dia, agora time.Time) {
	fmt.Printf("data: %s\n", d.Data)
	fmt.Printf("inicio: %s\n", d.Inicio.Format(formatoHora))
	fmt.Printf("tempo trabalhado: %s\n", fmtDur(d.trabalhado(agora)))

	if len(d.Pausas) == 0 {
		fmt.Println("pausas: nenhuma")
		return
	}

	fmt.Println("pausas:")
	var total time.Duration
	for _, p := range d.Pausas {
		fim := p.Fim
		emAndamento := ""
		if fim.IsZero() {
			fim = agora
			emAndamento = " (em andamento)"
		}

		dur := fim.Sub(p.Inicio)
		total += dur
		fmt.Printf("  - %s  %s-%s  %s%s\n",
			p.Motivo, p.Inicio.Format(formatoHora), fim.Format(formatoHora), fmtDur(dur), emAndamento)
	}
}

func fmtDur(d time.Duration) string {
	d = d.Round(time.Minute)
	return fmt.Sprintf("%dh%02dm", int(d.Hours()), int(d.Minutes())%60)
}
