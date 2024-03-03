package main

import (
	"fmt"
	"os"
	"strconv"
)

type Siswa struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	data_siswa := getNoSiswa()
	if len(os.Args) != 1 {
		for i := 1; i < len(os.Args); i++ {
			no, _ := strconv.Atoi(os.Args[i])

			if no <= 0 {
				fmt.Println("Data Biodata Tidak Valid!!", no)
				continue
			}
			if no >= len(data_siswa) {
				fmt.Println("Data Biodata Tidak Ditemukan !!")
				continue
			}

			Cetak(data_siswa, no)
		}
	} else {
		fmt.Println("Masukkan No Setelah biodata.go seperti contoh : biodata.go 1 !!!")
	}

}

func getNoSiswa() []Siswa {
	data_siswa := []Siswa{
		{},
		{nama: "Alifia", alamat: "Jl. Jambu", pekerjaan: "Mahasiswa", alasan: "Ingin Mempelajari Golang"},
		{nama: "Septi", alamat: "Jl. Mekar", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Hal Baru"},
		{nama: "Budi", alamat: "Jl. Bagus", pekerjaan: "Pengawas", alasan: "Belajar Sesuatu Yang Baru"},
		{nama: "Laila", alamat: "Jl. Jambu", pekerjaan: "Admin", alasan: "Belajar Bahasa Golang"},
		{nama: "Irene", alamat: "Jl. Sekitar", pekerjaan: "Admin", alasan: "Belajar Pemograman"},
		{nama: "Velvet", alamat: "Jl. Red", pekerjaan: "Belum Bekerja", alasan: "Beradaptasi Bahasa Pemograman"},
		{nama: "Yanto", alamat: "Jl. Mekar", pekerjaan: "Operator", alasan: "Mencoba Hal Baru"},
		{nama: "Tono", alamat: "Jl.Sari", pekerjaan: "Finance", alasan: "Tertarik Bahasa Golang"},
	}
	return data_siswa
}

func Cetak(data_siswa []Siswa, no int) {
	fmt.Println("Nama :", data_siswa[no].nama)
	fmt.Println("Alamat :", data_siswa[no].alamat)
	fmt.Println("Pekerjaan :", data_siswa[no].pekerjaan)
	fmt.Println("Alasan Memilih Kelas Golang :", data_siswa[no].alasan)
}
