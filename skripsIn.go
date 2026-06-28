package main

import "fmt"

const MAKS int = 100

type Skripsi struct {
	nim    string
	nama   string
	judul  string
	status string
	nilai  float64
}

var data [MAKS]Skripsi
var jumlah int = 0
var keluar bool = false

func garis() {
	fmt.Println("--------------------------------------------------")
}

func tampilSatu(i int) {
	fmt.Printf("  NIM    : %s\n", data[i].nim)
	fmt.Printf("  Nama   : %s\n", data[i].nama)
	fmt.Printf("  Judul  : %s\n", data[i].judul)
	fmt.Printf("  Status : %s\n", data[i].status)
	fmt.Printf("  Nilai  : %.2f\n", data[i].nilai)
	garis()
}

func statusValid(s string) bool {
	if s == "Lulus" {
		return true
	}
	if s == "Revisi" {
		return true
	}
	if s == "TidakLulus" {
		return true
	}
	return false
}

func nimValid(nim string) bool {
	var i int
	var panjang int

	panjang = 0
	for range nim {
		panjang = panjang + 1
	}

	if panjang != 12 {
		return false
	}

	for i = 0; i < panjang; i++ {
		if nim[i] < '0' || nim[i] > '9' {
			return false
		}
	}

	return true
}

func cariNIM(nim string) int {
	var i int
	var hasil int
	hasil = -1
	for i = 0; i < jumlah; i++ {
		if data[i].nim == nim {
			hasil = i
		}
	}
	return hasil
}

func tambah() {
	var nim, nama, judul, status string
	var nilai float64

	fmt.Println("\n=== TAMBAH DATA ===")
	if jumlah >= MAKS {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("NIM    : ")
	fmt.Scan(&nim)

	if !nimValid(nim) {
		fmt.Println("NIM harus terdiri dari 12 angka!")
		return
	}

	if cariNIM(nim) != -1 {
		fmt.Println("NIM sudah terdaftar!")
		return
	}

	fmt.Print("Nama   : ")
	fmt.Scan(&nama)
	fmt.Print("Judul  : ")
	fmt.Scan(&judul)
	fmt.Print("Status (Lulus/Revisi/TidakLulus) : ")
	fmt.Scan(&status)
	if !statusValid(status) {
		fmt.Println("Status tidak valid!")
		return
	}
	fmt.Print("Nilai  : ")
	fmt.Scan(&nilai)
	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai harus 0-100!")
		return
	}

	data[jumlah] = Skripsi{nim, nama, judul, status, nilai}
	jumlah = jumlah + 1
	fmt.Println("Data berhasil ditambahkan.")
}

func tampilSemua() {
	var i int
	fmt.Println("\n=== DAFTAR SKRIPSI ===")
	garis()
	if jumlah == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i = 0; i < jumlah; i++ {
		fmt.Printf("  [%d]\n", i+1)
		tampilSatu(i)
	}
}

func edit() {
	var nim, input string
	var nilai float64
	var idx int

	fmt.Println("\n=== EDIT DATA ===")
	fmt.Print("NIM yang diedit : ")
	fmt.Scan(&nim)
	idx = cariNIM(nim)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("Data saat ini:")
	tampilSatu(idx)
	fmt.Println("(Ketik - untuk skip)")

	fmt.Print("Nama baru   : ")
	fmt.Scan(&input)
	if input != "-" {
		data[idx].nama = input
	}

	fmt.Print("Judul baru  : ")
	fmt.Scan(&input)
	if input != "-" {
		data[idx].judul = input
	}

	fmt.Print("Status baru (Lulus/Revisi/TidakLulus) : ")
	fmt.Scan(&input)
	if input != "-" {
		if statusValid(input) {
			data[idx].status = input
		} else {
			fmt.Println("Status tidak valid, dilewati.")
		}
	}

	fmt.Print("Nilai baru  : ")
	fmt.Scan(&nilai)
	if nilai >= 0 && nilai <= 100 {
		data[idx].nilai = nilai
	} else {
		fmt.Println("Nilai tidak valid, dilewati.")
	}

	fmt.Println("Data berhasil diperbarui.")
}

func hapus() {
	var nim, konfirm string
	var idx, i int

	fmt.Println("\n=== HAPUS DATA ===")
	fmt.Print("NIM yang dihapus : ")
	fmt.Scan(&nim)
	idx = cariNIM(nim)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	tampilSatu(idx)
	fmt.Print("Yakin hapus? (y/n) : ")
	fmt.Scan(&konfirm)
	if konfirm != "y" {
		fmt.Println("Dibatalkan.")
		return
	}

	for i = idx; i < jumlah-1; i++ {
		data[i] = data[i+1]
	}
	jumlah = jumlah - 1
	fmt.Println("Data berhasil dihapus.")
}

func BSNIM(nim string) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = jumlah - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].nim == nim {
			return tengah
		} else if data[tengah].nim < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func cariByNIM() {
	var nim string
	var idx int

	sortNIM()

	fmt.Print("NIM : ")
	fmt.Scan(&nim)

	idx = BSNIM(nim)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		tampilSatu(idx)
	}
}

func cariByStatus() {
	var status string
	var i int
	var ditemukan bool
	ditemukan = false
	fmt.Print("Status (Lulus/Revisi/TidakLulus) : ")
	fmt.Scan(&status)
	garis()
	for i = 0; i < jumlah; i++ {
		if data[i].status == status {
			tampilSatu(i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada data dengan status tersebut.")
	}
}

func cariByNama() {
	var nama string
	var i int
	var ditemukan bool

	ditemukan = false

	fmt.Print("Nama : ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if data[i].nama == nama {
			tampilSatu(i)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

func menuCari() {
	var pilihan string

	fmt.Println("\n=== CARI SKRIPSI ===")
	fmt.Println("1. Cari berdasarkan NIM")
	fmt.Println("2. Cari berdasarkan Status")
	fmt.Println("3. Cari berdasarkan Nama")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	if pilihan == "1" {
		cariByNIM()
	} else if pilihan == "2" {
		cariByStatus()
	} else if pilihan == "3" {
		cariByNama()
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func sortNIM() {
	var i, j int
	var kunci Skripsi
	for i = 1; i < jumlah; i++ {
		kunci = data[i]
		j = i - 1
		for j >= 0 && data[j].nim > kunci.nim {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = kunci
	}
	fmt.Println("Data diurutkan berdasarkan NIM.")
}

func sortNama() {
	var i, j int
	var kunci Skripsi

	for i = 1; i < jumlah; i++ {
		kunci = data[i]
		j = i - 1

		for j >= 0 && data[j].nama > kunci.nama {
			data[j+1] = data[j]
			j = j - 1
		}

		data[j+1] = kunci
	}

	fmt.Println("Data diurutkan berdasarkan Nama.")
}

func sortNilai() {
	var i, j, idxMaks int
	var temp Skripsi
	for i = 0; i < jumlah-1; i++ {
		idxMaks = i
		for j = i + 1; j < jumlah; j++ {
			if data[j].nilai > data[idxMaks].nilai {
				idxMaks = j
			}
		}
		if idxMaks != i {
			temp = data[i]
			data[i] = data[idxMaks]
			data[idxMaks] = temp
		}
	}
	fmt.Println("Data diurutkan berdasarkan Nilai (tertinggi ke terendah).")
}

func menuSort() {
	var pilihan string

	fmt.Println("\n=== URUTKAN DATA ===")
	fmt.Println("1. Urutkan berdasarkan Nama")
	fmt.Println("2. Urutkan berdasarkan NIM")
	fmt.Println("3. Urutkan berdasarkan Nilai")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	if pilihan == "1" {
		sortNama()
	} else if pilihan == "2" {
		sortNIM()
	} else if pilihan == "3" {
		sortNilai()
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	tampilSemua()
}

func laporan() {
	var i int
	var lulus, revisi, tidakLulus int
	var total, rataRata float64
	var idxMax, idxMin int

	fmt.Println("\n=== LAPORAN STATISTIK ===")
	garis()
	if jumlah == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	idxMax = 0
	idxMin = 0
	for i = 0; i < jumlah; i++ {
		if data[i].status == "Lulus" {
			lulus = lulus + 1
		} else if data[i].status == "Revisi" {
			revisi = revisi + 1
		} else if data[i].status == "TidakLulus" {
			tidakLulus = tidakLulus + 1
		}
		total = total + data[i].nilai
		if data[i].nilai > data[idxMax].nilai {
			idxMax = i
		}
		if data[i].nilai < data[idxMin].nilai {
			idxMin = i
		}
	}
	rataRata = total / float64(jumlah)

	fmt.Printf("Total data    : %d\n", jumlah)
	fmt.Printf("Lulus         : %d\n", lulus)
	fmt.Printf("Revisi        : %d\n", revisi)
	fmt.Printf("Tidak Lulus   : %d\n", tidakLulus)
	garis()
	fmt.Printf("Rata-rata     : %.2f\n", rataRata)
	fmt.Printf("Nilai tertinggi: %.2f (%s - %s)\n", data[idxMax].nilai, data[idxMax].nim, data[idxMax].nama)
	fmt.Printf("Nilai terendah : %.2f (%s - %s)\n", data[idxMin].nilai, data[idxMin].nim, data[idxMin].nama)
	garis()
}

func main() {
	var pilihan string

	for !keluar {
		fmt.Println("\n========== SKRIPSIN ==========")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Tampilkan Semua")
		fmt.Println("3. Edit Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Cari Skripsi")
		fmt.Println("6. Urutkan Data")
		fmt.Println("7. Laporan Statistik")
		fmt.Println("8. Keluar")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)

		if pilihan == "1" {
			tambah()
		} else if pilihan == "2" {
			tampilSemua()
		} else if pilihan == "3" {
			edit()
		} else if pilihan == "4" {
			hapus()
		} else if pilihan == "5" {
			menuCari()
		} else if pilihan == "6" {
			menuSort()
		} else if pilihan == "7" {
			laporan()
		} else if pilihan == "8" {
			fmt.Println("Sampai jumpa!")
			keluar = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}