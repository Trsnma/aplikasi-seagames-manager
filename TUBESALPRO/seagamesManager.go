package main

import (
	"fmt"
	"strings"
)

// Kamus Global
const NMAX = 999
type negara struct {
	Nama string
	Emas int
	Perak int
	Perunggu int
}

type daftarNegara [NMAX]negara
var jumlahNegara int = 0

// Fungsi untuk menampilkan menu utama
func menu() {
	fmt.Println("------------------------------------")
	fmt.Println("          SEAGAMES MANAGER          ")
	fmt.Println("------------------------------------")
	fmt.Println("1. Tambah Negara")
	fmt.Println("2. Ubah Data Negara dan Mendali")
	fmt.Println("3. Hapus Negara Peserta")
	fmt.Println("4. Tampilkan Peringkat Klasemen")
	fmt.Println("5. Keluar")
	fmt.Println("------------------------------------")
}

// Main Function
func main(){
	var pilihan, pUbah, kriteria, index int
	var nama, desc string
	var N daftarNegara

	pilihan = 0

// Looping menu utama
	for pilihan != 5 {

		menu()
		fmt.Println("Pilih 1/2/3/4/5 ?")
		fmt.Scan(&pilihan)

		if  pilihan  == 1 {
			fmt.Println("Masukkan nama negara yang ingin ditambahkan:")
			fmt.Scan(&nama)
			tambahNegara(&N, nama)
		} else if pilihan == 2 {
			fmt.Println("Masukkan nama negara yang ingin diubah:")
			fmt.Scan(&nama)
			index = cariNegaraIndex(N, nama)
			if index != -1 {
				fmt.Println("Apa yang ingin diubah?")
				fmt.Println("1. Data Negara")
				fmt.Println("2. Jumlah Medali")
				fmt.Println("Pilih 1/2?")
				fmt.Scan(&pUbah)

				if pUbah == 1 {
					ubahDataNegara(&N, nama, index)
				} else if pUbah == 2 {
					ubahDataMedali(&N, nama, index)
				} else {
					fmt.Println("\nGagal: Pilihan tidak valid.")
				}
			} else {
				fmt.Printf("\nGagal: Negara '%s' tidak ditemukan.\n", nama)
			}
			
		} else if pilihan == 3 {
			fmt.Println("Masukkan nama negara yang ingin dihapus:")
			fmt.Scan(&nama)
			hapusNegara(&N, nama)
		} else if pilihan == 4 {
			fmt.Println("Pilih kriteria peringkat:")
			fmt.Println("1. Peringkat Tertinggi ke Rendah")
			fmt.Println("2. Peringkat Rendah ke Tertinggi")
			fmt.Scan(&kriteria)

			if kriteria == 1 {
				peringkatDescending(&N, jumlahNegara)
				desc = "Descending"
				tampilkanKlasemen(N, jumlahNegara, desc)
			} else if kriteria == 2 {
				peringkatAscending(&N, jumlahNegara)
				desc = "Ascending"
				tampilkanKlasemen(N, jumlahNegara, desc)
			}
		} else if pilihan == 5 {
			fmt.Println("Terima kasih telah menggunakan SEAGAMES MANAGER!")
		} else {
			fmt.Println("\nGagal: Pilihan tidak valid. Silakan pilih 1/2/3/4/5.")
		}
	}
}

// Fungsi untuk mencari index negara berdasarkan nama
func cariNegaraIndex(N daftarNegara, nama string)int{
	var index int = -1
	var i int = 0

	for i < jumlahNegara && index == -1 {
		if strings.EqualFold(N[i].Nama, nama) {
			index = i
		}
		i++
	}
	return index
}

// fungsi untuk menambahkan negara baru ke dalam daftar
func tambahNegara(N *daftarNegara, nama string){
	var emas, perak, perunggu int

	if jumlahNegara >= NMAX {
		fmt.Println("\nGagal: Kapasitas maksimum negara peserta telah penuh!")
		return
	}

	if cariNegaraIndex(*N, nama) != -1 {
		fmt.Printf("\nGagal: Negara '%s' sudah terdaftar.\n", nama)
		return
	}

	fmt.Println("Masukkan jumlah medali emas, perak, dan perunggu:")
	fmt.Scan(&emas, &perak, &perunggu)
	N[jumlahNegara] = negara{Nama: nama, Emas: emas, Perak: perak, Perunggu: perunggu}
	jumlahNegara++
	fmt.Printf("\nBerhasil: Negara '%s' berhasil ditambahkan.\n", nama)
}

// fungsi untuk mengubah data negara
func ubahDataNegara(N *daftarNegara, nama string, index int) {
	var emasBaru, perakBaru, perungguBaru int

	fmt.Println("Masukkan jumlah medali emas, perak, dan perunggu yang baru:")
	fmt.Print("Jumlah medali emas:")
	fmt.Scan(&emasBaru)
	N[index].Emas = emasBaru
	fmt.Print("Jumlah medali perak:")
	fmt.Scan(&perakBaru)
	N[index].Perak = perakBaru
	fmt.Println("Jumlah medali perunggu:")
	N[index].Perunggu = perungguBaru
	fmt.Printf("\nBerhasil: Data medali untuk negara '%s' berhasil diubah.\n", nama)

}

// Fungsi untuk mengubah data mendali yang spesifik
func ubahDataMedali(N *daftarNegara, nama string, index int) {
	var jumlahBaru int
	var mendaliDiCari string

	fmt.Println("Masukkan jenis medali yang ingin diubah (emas/perak/perunggu):")
	fmt.Scan(&mendaliDiCari)
	fmt.Println("Masukkan jumlah medali yang baru:")
	fmt.Scan(&jumlahBaru)
	if strings.EqualFold(mendaliDiCari, "emas") {
		N[index].Emas = jumlahBaru
	} else if strings.EqualFold(mendaliDiCari, "perak") {
		N[index].Perak = jumlahBaru
	} else if strings.EqualFold(mendaliDiCari, "perunggu") {
		N[index].Perunggu = jumlahBaru
	}
}

// Fungsi untuk menghapus negara dari daftar
func hapusNegara(N *daftarNegara, nama string){
	var index int = cariNegaraIndex(*N, nama)
	var i int = index

	if index == -1 {
		fmt.Printf("\nGagal: Negara '%s' tidak ditemukan.\n", nama)
		return
	}

	for i < jumlahNegara-1 {
		N[i] = N[i+1]
		i++
	}

	N[jumlahNegara-1] = negara{}
	jumlahNegara--
	fmt.Printf("\nBerhasil: Negara '%s' berhasil dihapus.\n", nama)

}

// Fungsi untuk mengurutkan dari tertinggi ke terendah berdasarkan jumlah medali emas, perak, dan perunggu
func peringkatDescending(N *daftarNegara, jumlahNegara int) {
	var pass, index, i int
	var temp negara

	pass = 1
	for pass <= jumlahNegara-1 {
		index = pass - 1
		i = pass

		for i < jumlahNegara {
			if N[index].Emas < N[i].Emas {
				index = i
			} else if N[index].Emas == N[i].Emas {
				if N[index].Perak < N[i].Perak {
					index = i
				} else if N[index].Perak == N[i].Perak {
					if N[index].Perunggu < N[i].Perunggu {
						index = i
					}
				}
			}
			i++
		}
		temp = N[pass-1]
		N[pass-1] = N[index]
		N[index] = temp
		pass++
	}
}

// Fungsi untuk mengurutkan dari terendah ke tertinggi berdasarkan jumlah medali emas, perak, dan perunggu
func peringkatAscending(N *daftarNegara, jumlahNegara int) {
	var pass, i int
	var temp negara

	pass = 1
	for pass <= jumlahNegara-1 {
		i = pass
		temp = N[pass]

		for i > 0 && (temp.Emas < N[i-1].Emas || (temp.Emas == N[i-1].Emas && temp.Perak < N[i-1].Perak) || (temp.Emas == N[i-1].Emas && temp.Perak == N[i-1].Perak && temp.Perunggu < N[i-1].Perunggu)) {
			N[i] = N[i-1]
			i--
		}
		N[i] = temp
		pass++
	}
}

// Fungsi untuk menampilkan klasemen berdasarkan
func tampilkanKlasemen(daftar daftarNegara, jumlahnegara int, desc string) {
	if jumlahNegara == 0 {
		fmt.Println("\nBelum Ada Data Negara.")
		return
	}

	fmt.Printf("\n[Klasemen Peringkat: %s]", desc)
	fmt.Print("\n====================================================")
	fmt.Print("\n=             KLASEMEN MEDALI SEAGAMES             =")
	fmt.Println("\n====================================================")
	fmt.Printf("| %-3s | %-15s | %-5s | %-5s | %-5s |\n", "No", "Negara", "Emas", "Perak", "Perunggu")
	fmt.Println("----------------------------------------------------")
	for i := 0; i < jumlahNegara; i++ {
		fmt.Printf("| %-3d | %-15s | %-6d | %-6d | %-6d |\n", i+1, daftar[i].Nama, daftar[i].Emas, daftar[i].Perak, daftar[i].Perunggu)
		fmt.Println("----------------------------------------------------")
	}
}
	

