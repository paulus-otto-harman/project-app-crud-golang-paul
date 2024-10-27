package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

const PasienDB = "patients"

type Pasien struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	CreatedAt string `json:"created_at"`
}

func InitPasien(nama string, alamat string) Pasien {
	pasien := Pasien{Nama: nama, Alamat: alamat, CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
	pasien.Id = pasien.GetId()
	return pasien
}

func (pasien *Pasien) GetId() int {
	lastId := 0
	patients, ok := pasien.Retrieve().([]Pasien)
	if ok {
		for _, patient := range patients {
			if patient.Id > lastId {
				lastId = patient.Id
			}
		}
	}
	return lastId + 1
}

func (pasien *Pasien) Create() {
	patients, ok := pasien.Retrieve().([]Pasien)
	if ok {

		file, err := os.Create(fmt.Sprintf("%s/%s.json", DbPath, PasienDB))
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		patients = append(patients, *pasien)

		if err := encoder.Encode(patients); err != nil {
			fmt.Println("Error encoding JSON:", err)
		}
	}
}

func (pasien *Pasien) Retrieve() interface{} {
	file, err := os.Open(fmt.Sprintf("%s/%s.json", DbPath, PasienDB))
	if err != nil {
		fmt.Println("Retrieve : Error opening file:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var patients []Pasien
	if err := decoder.Decode(&patients); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return patients
}

func (pasien *Pasien) Save() {

}

func (pasien *Pasien) Delete() {

}

func (pasien *Pasien) FirstOrCreate(nama string, alamat string) Pasien {
	patients := pasien.Retrieve().([]Pasien)
	i := slices.IndexFunc(patients, func(pasien Pasien) bool {
		return pasien.Nama == nama && pasien.Alamat == alamat
	})
	if i == -1 {
		pasien := InitPasien(nama, alamat)
		pasien.Create()
		return pasien
	}
	return patients[i]

}

func (pasien *Pasien) First(nama string, alamat string) (Pasien, error) {
	patients := pasien.Retrieve().([]Pasien)
	i := slices.IndexFunc(patients, func(pasien Pasien) bool {
		return pasien.Nama == nama && pasien.Alamat == alamat
	})
	if i == -1 {
		return Pasien{}, errors.New("Not Found")
	}
	return patients[i], nil

}
