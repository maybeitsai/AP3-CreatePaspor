package main

import (
	"fmt"
	"encoding/json"
	"time"
	"io/ioutil"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// DataPaspor adalah struktur data untuk menyimpan informasi paspor
type DataPaspor struct {
	Kategori       string    `json:"kategori"`
	NomorKK        string    `json:"nomor_kk"`
	NIK            string    `json:"nik"`
	Nama           string    `json:"nama"`
	TempatLahir    string    `json:"tempat_lahir"`
	TanggalLahir   string    `json:"tanggal_lahir"`
	Alamat         string    `json:"alamat"`
	JenisKelamin   string    `json:"jenis_kelamin"`
	WaktuPembuatan time.Time `json:"waktu_pembuatan"`
}

var pasporList []DataPaspor

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Pembuatan Paspor")

	// Formulir input
	kategoriEntry := widget.NewSelect([]string{"Anak-Anak", "Dewasa"}, nil)
	nomorKKEntry := widget.NewEntry()
	nikEntry := widget.NewEntry()
	namaEntry := widget.NewEntry()
	jenisKelaminRadio := widget.NewRadioGroup([]string{"Laki-Laki", "Perempuan"}, func(selected string) {
		// Do nothing for now
	})
	tempatLahirEntry := widget.NewEntry()
	tanggalLahirEntry := widget.NewEntry()
	alamatEntry := widget.NewEntry()

	// Tombol submit
	submitButton := widget.NewButton("Submit", func() {
		// Validasi input
		if kategoriEntry.Selected == "" || nomorKKEntry.Text == "" || nikEntry.Text == "" ||
			namaEntry.Text == "" || tempatLahirEntry.Text == "" || tanggalLahirEntry.Text == "" ||
			alamatEntry.Text == "" || jenisKelaminRadio.Selected == "" {

			dialog.ShowError(fmt.Errorf("Mohon isi semua kolom"), myWindow)
			return
		}

		pasporList = readDataFromJSON()
			for _, paspor := range pasporList {
				if paspor.NIK == nikEntry.Text {
					dialog.ShowInformation("Informasi", "NIK sudah terdaftar.", myWindow)
					return
				}
			} 

		// Simpan data paspor ke dalam list
		paspor := DataPaspor{
			Kategori:       kategoriEntry.Selected,
			NomorKK:        nomorKKEntry.Text,
			NIK:            nikEntry.Text,
			Nama:           namaEntry.Text,
			JenisKelamin:   jenisKelaminRadio.Selected,
			TempatLahir:    tempatLahirEntry.Text,
			TanggalLahir:   tanggalLahirEntry.Text,
			Alamat:         alamatEntry.Text,
			WaktuPembuatan: time.Now(),
		}
		pasporList = append(pasporList, paspor)

		// Write data to JSON file
		writeDataToJSON(pasporList)	

		// Reset formulir
		kategoriEntry.SetSelected("")
		nomorKKEntry.SetText("")
		nikEntry.SetText("")
		namaEntry.SetText("")
		jenisKelaminRadio.SetSelected("")
		tempatLahirEntry.SetText("")
		tanggalLahirEntry.SetText("")
		alamatEntry.SetText("")

		// Tampilkan konfirmasi
		showSubmitDialog(myWindow, paspor)
	})
	// Membuat tata letak
	form := container.NewVBox(
		widget.NewLabel("Pilih Kategori:"),
		kategoriEntry,
		widget.NewLabel("Nomor KK:"),
		nomorKKEntry,
		widget.NewLabel("NIK:"),
		nikEntry,
		widget.NewLabel("Nama:"),
		namaEntry,
		widget.NewLabel("Jenis Kelamin:"),
		jenisKelaminRadio,
		widget.NewLabel("Tempat Lahir:"),
		tempatLahirEntry,
		widget.NewLabel("Tanggal Lahir (DD/MM/YYYY):"),
		tanggalLahirEntry,
		widget.NewLabel("Alamat:"),
		alamatEntry,
		layout.NewSpacer(),
		submitButton,
	)

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}

func showSubmitDialog(window fyne.Window, paspor DataPaspor) {
	content := container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Kategori: %s", paspor.Kategori)),
		widget.NewLabel(fmt.Sprintf("Nomor KK: %s", paspor.NomorKK)),
		widget.NewLabel(fmt.Sprintf("NIK: %s", paspor.NIK)),
		widget.NewLabel(fmt.Sprintf("Nama: %s", paspor.Nama)),
		widget.NewLabel(fmt.Sprintf("Jenis Kelamin: %s", paspor.JenisKelamin)),
		widget.NewLabel(fmt.Sprintf("Tempat Lahir: %s", paspor.TempatLahir)),
		widget.NewLabel(fmt.Sprintf("Tanggal Lahir: %s", paspor.TanggalLahir)),
		widget.NewLabel(fmt.Sprintf("Alamat: %s", paspor.Alamat)),
		widget.NewLabel(fmt.Sprintf("Waktu Pembuatan: %s", paspor.WaktuPembuatan.Format(time.RFC3339))),
	)

	dialog.ShowCustom("Data Paspor", "Close", content, window)
}

func writeDataToJSON(data []DataPaspor) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("data_paspor.json", file, 0644)
}

func readDataFromJSON() []DataPaspor {
	file, _ := ioutil.ReadFile("data_paspor.json")
	var data []DataPaspor
	_ = json.Unmarshal(file, &data)
	return data
}