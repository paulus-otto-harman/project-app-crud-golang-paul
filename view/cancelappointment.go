package view

import (
	"context"
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type CancelAppointment struct {
}

func (appointment *CancelAppointment) Render(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			gola.ClearScreen()

			var idPasien interface{}

			for err := errors.New(""); err != nil; {
				idPasien, err = gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%s : ", "ID Pasien atau [0] untuk kembali"))))
			}
			if idPasien == "0" {
				return
			}
		}
	}
}
