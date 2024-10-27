package view

import (
	"context"
	"fmt"
	"project/model"
	"strings"
)

type Appointments struct {
}

func (list *Appointments) Render(ctx context.Context) {
	appointments := (&model.JanjiTemu{}).Retrieve().([]model.JanjiTemu)

	const WidthId = 14
	const WidthDate = 11
	const WidthNama = 20
	const WidthAlamat = 40

	// header
	fmt.Printf("%s%s%s%s%s%s%s%s%s\n", "┌─", strings.Repeat("─", WidthId), "┬─", strings.Repeat("─", WidthDate), "┬─", strings.Repeat("─", WidthNama), "─┬─", strings.Repeat("─", WidthAlamat), "─┐")
	fmt.Printf("%s%-*s%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthId, "ID Janji Temu", "│ ", WidthDate, "Tanggal", "│ ", WidthNama, "Nama Pasien", " │ ", WidthAlamat, "Alamat Pasien", " │")
	fmt.Printf("%s%s%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthId), "┼─", strings.Repeat("─", WidthDate), "┼─", strings.Repeat("─", WidthNama), "─┼─", strings.Repeat("─", WidthAlamat), "─┤")

	// body
	for _, appointment := range appointments {
		if appointment.TidakDatang {
			fmt.Printf("│  %-12s │ %-10s │ %-*s │ %-*s │\n", fmt.Sprintf("%2d", appointment.Id), appointment.Tanggal, WidthNama, appointment.Pasien.Nama, WidthAlamat, appointment.Pasien.Alamat)
		}
	}

	// footer
	fmt.Printf("%s%s%s%s%s%s%s%s%s\n", "╘═", strings.Repeat("═", WidthId), "╧═", strings.Repeat("═", WidthDate), "╧═", strings.Repeat("═", WidthNama), "═╧═", strings.Repeat("═", WidthAlamat), "═╛")
}
