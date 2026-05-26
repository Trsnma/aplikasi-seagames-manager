package main

import (
	"fmt"
	"strings"
)

const NMAX = 999
type negara struct {
	Nama string
	Emas int
	Perak int
	Perunggu int
}

var daftarNegara [NMAX]negara
var jumlahNegara int = 0

func menu() {
	fmt.Println("------------------------------------")
	fmt.Println("          SEAGAMES MANAGER          ")
	fmt.Println("------------------------------------")
	fmt.Println("1. Tambah Negara")
	fmt.Println("2. Ubah Data Negara dan Mendali")
	fmt.Println("3. Hapus Negara Peserta")
	fmt.Println("4. Tampilkan Peringkat Klasemen")
	fmt.Println("------------------------------------")
}


func main(){
	var pilihan, pUbah, index int
	var nama string
	var N daftarNegara

	menu()
	fmt.Println("Pilih 1/2/3/4 ?")
	fmt.Scan(&pilihan)

	for pilihan != 5 {

		if  pilihan  == 1 {
			fmt.Scan(&nama)
			tambahNegara(nama)
		} else if pilihan == 2 {
			fmt.Println("Masukkan nama negara yang ingin diubah:")
			fmt.Scan(&nama)
			index = cariNegaraIndex(nama)
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
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				peringkatDescending(&N, jumlahNegara)
				tampilkanKlasemen()
			} else if pilihan == 2 {
				peringkatAscending(&N, jumlahNegara)
			}
		}
	}

	
	
}

func cariNegaraIndex(nama string)int{
	var index int = -1
	var i int = 0

	for i < jumlahNegara && index == -1 {
		if strings.EqualFold(daftarNegara[i].Nama, nama) {
			index = i
		}
		i++
	}
	return index
}

func tambahNegara(nama string){
	var emas, perak, perunggu int

	if jumlahNegara <= NMAX {
		fmt.Println("\nGagal: Kapasitas maksimum negara peserta telah penuh!")
		return
	}

	if cariNegaraIndex(nama) != -1 {
		fmt.Printf("\nGagal: Negara '%s' sudah terdaftar.\n", nama)
		return
	}

	fmt.Println("Masukkan jumlah medali emas, perak, dan perunggu:")
	fmt.Scan(&emas, &perak, &perunggu)
	daftarNegara[jumlahNegara] = negara{Nama: nama, Emas: emas, Perak: perak, Perunggu: perunggu}
	jumlahNegara++
	fmt.Printf("\nBerhasil: Negara '%s' berhasil ditambahkan.\n", nama)
}

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

func hapusNegara(N *daftarNegara, nama string){
	var index int = cariNegaraIndex(nama)
	var i int = index

	if index == -1 {
		fmt.Printf("\nGagal: Negara '%s' tidak ditemukan.\n", nama)
		return
	}

	for i < jumlahNegara-1 {
		N[i] = N[i+1]
		i++
	}

	jumlahNegara[jumlahNegara-1] = negara{}
	jumlahNegara--
	fmt.Printf("\nBerhasil: Negara '%s' berhasil dihapus.\n", nama)

}

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
		}
		temp = N[pass-1]
		N[pass-1] = N[index]
		N[index] = temp
		pass++
	}
}

func peringkatAscending(N *daftarNegara, jumlahNegara int) {
	var pass, i int
	var temp negara

	pass = 1
	for pass <= jumlahNegara-1 {
		i = pass
		temp = N[pass-1]

		for i > 0 && temp < N[i-1] {
			N[i] = N[i-1]
			i--
		}
		N[i] = temp
		pass++
	}
}

func tampilkanKlasemen() {
	var daftar daftarNegara
	if jumlahNegara == 0 {
		fmt.Println("\nBelum Ada Data Negara.")
		return
	}

	fmt.Println("\n--- KLASEMEN MEDALI SEAGAMES ---")
	fmt.Printf("%-3s | %-15s | %-5s | %-5s | %-5s\n", "No", "Negara", "Emas", "Perak", "Perunggu")
	fmt.Println("--------------------------------------------------")
	for i := 0; i < jumlahNegara; i++ {
		fmt.Printf("%-3d | %-15s | %-5d | %-5d | %-5d\n", i+1, daftar.Nama, daftar.Emas, daftar.Perak, daftar.Perunggu)
		fmt.Println("--------------------------------------------------")
	}
}
	

