package view

import (
	"context"
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/model"
	"project/util"
)

type MakeAppointment struct {
}

func (appointment *MakeAppointment) Render(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			util.ViewTitle("Klinik Dokter Umum - Buat Janji")

			var tanggal string
			for err := errors.New(""); err != nil; {
				tanggal, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-67s :", "Tanggal (DD-MM-YYYY) atau [0] untuk kembali")))))
			}
			if tanggal == "0" {
				return
			}

			var nama string
			for err := errors.New(""); err != nil; {
				nama, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-67s :", "Nama Pasien")))))
			}

			var alamat string
			for err := errors.New(""); err != nil; {
				alamat, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-67s :", "Alamat Pasien")))))
			}

			save, _ := gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-61s [y/n] :", "Buat Janji?"))))
			if save == "y" {

				model.InitJanjiTemu(tanggal, (&model.Pasien{}).FirstOrCreate(nama, alamat)).Create()

				repeat, _ := gola.Input(gola.Args(gola.P("label", fmt.Sprintf("\n%s :", "Enter untuk membuat janji lagi atau [0] untuk kembali ke Menu Utama"))))
				if repeat == "0" {
					return
				}
			}

		}
	}
}
