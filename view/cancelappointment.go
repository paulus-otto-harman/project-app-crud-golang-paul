package view

import (
	"context"
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/model"
	"project/util"
)

type CancelAppointment struct {
}

func (screen *CancelAppointment) Render(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			util.ViewTitle("Klinik Dokter Umum - Batalkan Janji")

			(&Appointments{}).Render(ctx)

			var idJanjiTemu int
			for err := errors.New(""); err != nil; {
				idJanjiTemu, err = gola.ToInt(gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("\n%s : ", "Masukkan ID Janji Temu atau [0] untuk kembali")))))
			}

			if idJanjiTemu == 0 {
				return
			}

			appointment, err := model.FindAppointmentById(idJanjiTemu)
			if err != nil {
				gola.Wait("ID Janji Temu tidak ditemukan")
			} else {

				var confirm string
				for err := errors.New(""); err != nil; {
					confirm, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("\n%s :", "Konfirmasi Pembatalan Janji Temu? [y/n]")))))
				}
				if confirm == "y" {
					appointment.Delete()
					gola.Wait("Pembatalan Jadwal Temu Pasien Berhasil. Tekan Enter untuk melanjutkan")
				}
			}

		}
	}
}
