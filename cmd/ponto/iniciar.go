package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var iniciarCmd = &cobra.Command{
	Use:   "iniciar",
	Short: "Inicia o dia de trabalho",
	Args:  cobra.NoArgs,
	RunE: func(c *cobra.Command, args []string) error {
		return iniciar(time.Now())
	},
}

func iniciar(agora time.Time) error {
	hoje := agora.Format(formatoData)

	d, err := carregarDia(hoje)
	if err != nil {
		return err
	}
	if d != nil {
		if d.Fim.IsZero() {
			return errors.New("a jornada de hoje ja esta em andamento")
		}

		return errors.New("a jornada de hoje ja foi encerrada")
	}

	d = NewDia(hoje, agora)

	fmt.Printf("bom dia! jornada iniciada as %s.\n", agora.Format("15:04"))

	return salvarDia(d)
}

func carregarDia(data string) (*Dia, error) {
	dir, err := diretorio()
	if err != nil {
		return nil, err
	}

	p := filepath.Join(dir, data+extensao)
	raw, err := os.ReadFile(p)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var d Dia
	if err := json.Unmarshal(raw, &d); err != nil {
		return nil, fmt.Errorf("arquivo %s corrompido: %w", p, err)
	}

	return &d, nil

}

func diretorio() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".ponto"), nil
}

func salvarDia(d *Dia) error {
	dir, err := diretorio()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, d.Data+extensao), raw, 0o600)
}
