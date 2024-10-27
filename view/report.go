package view

import (
	"context"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/model"
	"project/util"
	"strings"
)

type Report struct {
}

func (list *Report) Render(ctx context.Context) {
	appointments := (&model.JanjiTemu{}).Retrieve().([]model.JanjiTemu)
	statusKedatangan := map[bool]string{true: "Belum Datang", false: "Sudah Ke Klinik"}
	const WidthId = 14
	const WidthDate = 11
	const WidthNama = 20
	const WidthAlamat = 40
	const WidthStatus = 17

	util.ViewTitle("Klinik Dokter Umum - Laporan Kunjungan Pasien")

	// header
	fmt.Printf("%s%s%s%s%s%s%s%s%s%s%s\n", "┌─", strings.Repeat("─", WidthId), "┬─", strings.Repeat("─", WidthDate), "┬─", strings.Repeat("─", WidthNama), "─┬─", strings.Repeat("─", WidthAlamat), "─┬─", strings.Repeat("─", WidthStatus), "─┐")
	fmt.Printf("%s%-*s%s%-*s%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthId, "ID Janji Temu", "│ ", WidthDate, "Tanggal", "│ ", WidthNama, "Nama Pasien", " │ ", WidthAlamat, "Alamat Pasien", " │ ", WidthStatus, "Status Kedatangan", " │")
	fmt.Printf("%s%s%s%s%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthId), "┼─", strings.Repeat("─", WidthDate), "┼─", strings.Repeat("─", WidthNama), "─┼─", strings.Repeat("─", WidthAlamat), "─┼─", strings.Repeat("─", WidthStatus), "─┤")

	// body
	for _, appointment := range appointments {
		fmt.Printf("│  %-12s │ %-10s │ %-*s │ %-*s │ %-*s │\n", fmt.Sprintf("%2d", appointment.Id), appointment.Tanggal, WidthNama, appointment.Pasien.Nama, WidthAlamat, appointment.Pasien.Alamat, WidthStatus, statusKedatangan[appointment.TidakDatang])
	}

	// footer
	fmt.Printf("%s%s%s%s%s%s%s%s%s%s%s\n", "╘═", strings.Repeat("═", WidthId), "╧═", strings.Repeat("═", WidthDate), "╧═", strings.Repeat("═", WidthNama), "═╧═", strings.Repeat("═", WidthAlamat), "═╧═", strings.Repeat("═", WidthStatus), "═╛")

	gola.Wait("Tekan Enter untuk kembali ke Menu Utama")
}
