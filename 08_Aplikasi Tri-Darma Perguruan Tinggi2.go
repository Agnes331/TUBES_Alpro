package main

import (
	"fmt"
	"strings"
)

// "Penelitian" atau "Abdimas"
type PenelitianAbdimas struct {
	Ketua      string
	Anggota    [4]string
	Prodi      string
	Judul      string
	SumberDana string
	Luaran     string
	Jenis      string
	Tahun      int
}

const MaxData = 100

var data [MaxData]PenelitianAbdimas
var jumlahData int

// Tambah Data
func tambahData(pa PenelitianAbdimas) {
	if jumlahData >= MaxData {
		fmt.Println("Data sudah penuh, tidak dapat menambahkan data baru.")
		return
	}
	data[jumlahData] = pa
	jumlahData++
	fmt.Println("Data berhasil ditambahkan!")
}

// Edit Data
func editData(index int, pa PenelitianAbdimas) {
	if index < 0 || index >= jumlahData {
		fmt.Println("Index tidak valid!")
		return
	}
	data[index] = pa
	fmt.Println("Data berhasil diubah!")
}

// Hapus Data
func hapusData(index int) {
	if index < 0 || index >= jumlahData {
		fmt.Println("Index tidak valid!")
		return
	}
	for i := index; i < jumlahData-1; i++ {
		data[i] = data[i+1]
	}
	jumlahData--
	fmt.Println("Data berhasil dihapus!")
}

// Tampilkan Data
func tampilkanData(filterProdi string, filterTahun int) {
	for i := 0; i < jumlahData; i++ {
		pa := data[i]
		if (filterProdi == "" || strings.EqualFold(pa.Prodi, filterProdi)) && (filterTahun == 0 || pa.Tahun == filterTahun) {
			fmt.Printf("%d. Ketua: %s, Anggota: %v, Prodi: %s, Judul: %s, Sumber Dana: %s, Luaran: %s, Jenis: %s, Tahun: %d\n",
				i+1, pa.Ketua, pa.Anggota, pa.Prodi, pa.Judul, pa.SumberDana, pa.Luaran, pa.Jenis, pa.Tahun)
		}
	}
}

// Urutkan Data Berdasarkan Tahun (Selection Sort)
func urutkanDataBerdasarkanTahun(ascending bool) {
	for i := 0; i < jumlahData-1; i++ {
		idx := i
		for j := i + 1; j < jumlahData; j++ {
			if ascending {
				if data[j].Tahun < data[idx].Tahun {
					idx = j
				}
			} else {
				if data[j].Tahun > data[idx].Tahun {
					idx = j
				}
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tahun!")
}

// Cari Data Berdasarkan Judul (Sequential Search)
func cariDataJudul(judul string) int {
	for i := 0; i < jumlahData; i++ {
		if strings.EqualFold(data[i].Judul, judul) {
			return i
		}
	}
	return -1
}

func main() {
	var pilihan int
	for {
		fmt.Println("\n=== Aplikasi Pengolahan Data PPM ===")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("5. Urutkan Data Berdasarkan Tahun")
		fmt.Println("6. Cari Data Berdasarkan Judul")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var pa PenelitianAbdimas
			fmt.Print("Masukkan Ketua: ")
			fmt.Scan(&pa.Ketua)

			fmt.Println("Masukkan Anggota (maksimal 4, pisahkan dengan koma):")
			var anggota string
			fmt.Scan(&anggota)
			anggotaSlice := strings.Split(anggota, ",")
			for i := 0; i < len(anggotaSlice) && i < 4; i++ {
				pa.Anggota[i] = strings.TrimSpace(anggotaSlice[i])
			}

			fmt.Print("Masukkan Prodi: ")
			fmt.Scan(&pa.Prodi)
			fmt.Print("Masukkan Judul: ")
			fmt.Scan(&pa.Judul)
			fmt.Print("Masukkan Sumber Dana: ")
			fmt.Scan(&pa.SumberDana)
			fmt.Print("Masukkan Luaran: ")
			fmt.Scan(&pa.Luaran)
			fmt.Print("Masukkan Jenis (Penelitian/Abdimas): ")
			fmt.Scan(&pa.Jenis)
			fmt.Print("Masukkan Tahun: ")
			fmt.Scan(&pa.Tahun)

			tambahData(pa)

		case 2:
			var index int
			fmt.Print("Masukkan index data yang ingin diubah: ")
			fmt.Scan(&index)
			index--

			if index < 0 || index >= jumlahData {
				fmt.Println("Index tidak valid!")
				continue
			}

			var pa PenelitianAbdimas
			fmt.Print("Masukkan Ketua baru: ")
			fmt.Scan(&pa.Ketua)
			fmt.Println("Masukkan Anggota baru (maksimal 4, pisahkan dengan koma):")
			var anggota string
			fmt.Scan(&anggota)
			anggotaSlice := strings.Split(anggota, ",")
			for i := 0; i < len(anggotaSlice) && i < 4; i++ {
				pa.Anggota[i] = strings.TrimSpace(anggotaSlice[i])
			}
			fmt.Print("Masukkan Prodi baru: ")
			fmt.Scan(&pa.Prodi)
			fmt.Print("Masukkan Judul baru: ")
			fmt.Scan(&pa.Judul)
			fmt.Print("Masukkan Sumber Dana baru: ")
			fmt.Scan(&pa.SumberDana)
			fmt.Print("Masukkan Luaran baru: ")
			fmt.Scan(&pa.Luaran)
			fmt.Print("Masukkan Jenis baru (Penelitian/Abdimas): ")
			fmt.Scan(&pa.Jenis)
			fmt.Print("Masukkan Tahun baru: ")
			fmt.Scan(&pa.Tahun)

			editData(index, pa)

		case 3:
			var index int
			fmt.Print("Masukkan index data yang ingin dihapus: ")
			fmt.Scan(&index)
			hapusData(index - 1)

		case 4:
			var filterProdi string
			var filterTahun int
			fmt.Print("Masukkan filter Prodi (kosongkan jika tidak ada): ")
			fmt.Scan(&filterProdi)
			fmt.Print("Masukkan filter Tahun (0 jika tidak ada): ")
			fmt.Scan(&filterTahun)
			tampilkanData(filterProdi, filterTahun)

		case 5:
			var ascending bool
			fmt.Print("Urutkan secara ascending? (true/false): ")
			fmt.Scan(&ascending)
			urutkanDataBerdasarkanTahun(ascending)

		case 6:
			var judul string
			fmt.Print("Masukkan judul yang dicari: ")
			fmt.Scan(&judul)
			index := cariDataJudul(judul)
			if index == -1 {
				fmt.Println("Data tidak ditemukan!")
			} else {
				fmt.Printf("Data ditemukan di index %d: %+v\n", index+1, data[index])
			}

		case 7:
			fmt.Println("Keluar dari aplikasi.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
