package model

import "time"

type Pasien struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	CreatedAt string `json:"created_at"`
}

func InitPasien(nama string, alamat string) Pasien {
	return Pasien{Nama: nama, Alamat: alamat, CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
}

func (pasien *Pasien) FirstOrCreate(nama string, alamat string) Pasien {

	return Pasien{}
}

func (pasien *Pasien) First(nama string, alamat string) Pasien {

	return Pasien{}
}
