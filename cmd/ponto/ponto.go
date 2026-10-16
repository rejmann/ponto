package main

import "time"

const (
	formatoData = "2026-01-02"
	extensao    = ".json"
)

type Periodo struct {
	Inicio time.Time `json:"inicio"`
	Fim    time.Time `json:"fim,omitempty"`
}

type Pausa struct {
	Motivo string `json:"motivo"`
	Periodo
}

type Dia struct {
	Data   string  `json:"data"`
	Pausas []Pausa `json:"pausas"`
	Periodo
}

func (d *Dia) pausaAtiva() *Pausa {
	if n := len(d.Pausas); n > 0 && d.Pausas[n-1].Fim.IsZero() {
		return &d.Pausas[n-1]
	}
	return nil
}

func (d *Dia) trabalhado(agora time.Time) time.Duration {
	fim := d.Fim
	if fim.IsZero() {
		fim = agora
	}

	total := fim.Sub(d.Inicio)
	for _, p := range d.Pausas {
		f := p.Fim
		if f.IsZero() {
			f = agora
		}

		total -= f.Sub(p.Inicio)
	}

	return total
}

func NewDia(hoje string, agora time.Time) *Dia {
	return &Dia{Data: hoje, Periodo: Periodo{Inicio: agora}}
}

func NewPausa(motivo string, agora time.Time) *Pausa {
	return &Pausa{Motivo: motivo, Periodo: Periodo{Inicio: agora}}
}
