package view

import (
	"context"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/util"
)

type Home struct {
	isLogout bool
}

func (home *Home) Render(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			util.ViewTitle("Klinik Dokter Umum - Menu Utama")
			fmt.Println("[1] Buat Janji")
			fmt.Println("[2] Ubah Jadwal (Reschedule)")
			fmt.Println("[3] Batalkan Janji")
			fmt.Println("[4] Kunjungan Pasien")
			fmt.Println("[5] Laporan")

			fmt.Println()
			fmt.Println("[0] Logout")

			menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[5] atau [0] untuk Logout"))))
			switch menuItem.(int) {
			case 1:
				(&MakeAppointment{}).Render(ctx)
			case 2:
				(&Reschedule{}).Render(ctx)
			case 3:
				(&CancelAppointment{}).Render(ctx)
			case 4:

			case 0:
				home.isLogout = true
				return
			}
		}
	}
}
