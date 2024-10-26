package view

import (
	"context"
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/util"
)

type Login struct {
	Username string
}

func (login *Login) Render(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			util.ViewTitle("Klinik Dokter Umum - Login")

			for err := errors.New(""); err != nil; {
				login.Username, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Username atau [0] untuk keluar")))))
			}

			if login.Username == "0" {
				return
			}

			password, _ := gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Password"))))

			if login.Username == "x" && password == "x" {
				menuUtama := Home{}
				(&menuUtama).Render(ctx)
				if !menuUtama.isLogout {
					gola.Wait("Sesi telah berakhir. Tekan Enter untuk login kembali")
				}
			}
		}
	}
}
