package model

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const JanjiTemuDB = "appointments"

type JanjiTemu struct {
	Id          int    `json:"id"`
	Tanggal     string `json:"tanggal"`
	Pasien      Pasien `json:"pasien"`
	TidakDatang bool   `json:"status"`
	CreatedAt   string `json:"created_at"`
}

func InitJanjiTemu(tanggal string, pasien Pasien) *JanjiTemu {
	appointment := JanjiTemu{
		Tanggal:     tanggal,
		Pasien:      pasien,
		TidakDatang: true,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05")}
	appointment.Id = appointment.GetId()
	return &appointment
}

func (janjiTemu *JanjiTemu) GetId() int {
	lastId := 0
	appointments, ok := janjiTemu.Retrieve().([]JanjiTemu)
	if ok {
		for _, appointment := range appointments {
			if appointment.Id > lastId {
				lastId = appointment.Id
			}
		}
	}
	return lastId + 1
}

func (janjiTemu *JanjiTemu) Create() {
	orders, ok := janjiTemu.Retrieve().([]JanjiTemu)
	if ok {

		file, err := os.Create(fmt.Sprintf("%s/%s.json", DbPath, JanjiTemuDB))
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		orders = append(orders, *janjiTemu)

		if err := encoder.Encode(orders); err != nil {
			fmt.Println("Error encoding JSON:", err)
		}
	}
}

func (janjiTemu *JanjiTemu) Retrieve() interface{} {
	file, err := os.Open(fmt.Sprintf("%s/%s.json", DbPath, JanjiTemuDB))
	if err != nil {
		fmt.Println("Retrieve : Error opening file:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var appointments []JanjiTemu
	if err := decoder.Decode(&appointments); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return appointments
}

func (janjiTemu *JanjiTemu) Save() {
	appointments := (janjiTemu.Retrieve()).([]JanjiTemu)
	file, err := os.Create(fmt.Sprintf("%s/%s.json", DbPath, JanjiTemuDB))
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	for _, appointment := range appointments {
		if appointment.Id == janjiTemu.Id {
			appointment = *janjiTemu
			break
		}
	}

	if err := encoder.Encode(appointments); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
