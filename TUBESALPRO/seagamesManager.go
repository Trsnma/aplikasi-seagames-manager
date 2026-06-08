package main

import (
	"fmt"
	"strings"
)

// Kamus Global
const NMAX = 999

type negara struct {
	Nama     string
	Emas     int
	Perak    int
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
	fmt.Println("5. Analisis Kemenangan")
	fmt.Println("6. Keluar")
	fmt.Println("------------------------------------")
}

// Main Function
func main() {
	var pilihan, pUbah, kriteria, kriteriaCari, index int
	var nama, desc string
	var N daftarNegara

	pilihan = 0

	// Looping menu utama
	for pilihan != 6 {

		menu()
		fmt.Println("Pilih 1/2/3/4/5/6 ?")
		fmt.Print("> ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println("Masukkan nama negara yang ingin ditambahkan:")
			fmt.Print("> ")
			fmt.Scan(&nama)
			tambahNegara(&N, nama)
		} else if pilihan == 2 {
			fmt.Println("Masukkan nama negara yang ingin diubah:")
			fmt.Print("> ")
			fmt.Scan(&nama)
			index = cariNegaraIndex(N, nama)
			if index != -1 {
				fmt.Println("Apa yang ingin diubah?")
				fmt.Println("1. Data Negara")
				fmt.Println("2. Jumlah Medali")
				fmt.Println("Pilih 1/2?")
				fmt.Print("> ")
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
			fmt.Print("> ")
			fmt.Scan(&nama)
			hapusNegara(&N, nama)
		} else if pilihan == 4 {
			fmt.Println("Pilih kriteria peringkat:")
			fmt.Println("1. Peringkat Tertinggi ke Rendah")
			fmt.Println("2. Peringkat Rendah ke Tertinggi")
			fmt.Print("> ")
			fmt.Scan(&kriteria)

			if kriteria == 1 {
				peringkatDescending(&N, jumlahNegara)
				desc = "Descending"
				tampilkanKlasemen(N, jumlahNegara, desc)

				if jumlahNegara > 0 {
					fmt.Println("Ingin mencari negara dengan mendali tertentu?")
					fmt.Println("1. ya")
					fmt.Println("2. tidak")
					fmt.Print("> ")
					fmt.Scan(&kriteriaCari)

					pilihanSesuaiMendali(N, kriteriaCari, desc)
				}

			} else if kriteria == 2 {
				peringkatAscending(&N, jumlahNegara)
				desc = "Ascending"
				tampilkanKlasemen(N, jumlahNegara, desc)

				if jumlahNegara > 0 {
					fmt.Println("Ingin mencari negara dengan mendali tertentu?")
					fmt.Println("1. ya")
					fmt.Println("2. tidak")
					fmt.Print("> ")
					fmt.Scan(&kriteriaCari)

					pilihanSesuaiMendali(N, kriteriaCari, desc)
				}
			}
		} else if pilihan == 5 {
			fmt.Println("Masukan Nama Negara yang Ingin Dianalisis:")
			fmt.Print("> ")
			fmt.Scan(&nama)
			analisisKemenangan(N, nama)
		} else if pilihan == 6 {
			fmt.Println("Terima kasih telah menggunakan SEAGAMES MANAGER!")
		} else {
			fmt.Println("\nGagal: Pilihan tidak valid. Silakan pilih 1/2/3/4/5.")
		}
	}
}

// Fungsi untuk mencari index negara berdasarkan nama
func cariNegaraIndex(N daftarNegara, nama string) int {
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

// fungsi untuk menangani pilihan mencari negara berdasarkan kriteria mendali tertentu
func pilihanSesuaiMendali(N daftarNegara, kriteriaC int, desc string) {
	var jenisMendali string
	var jumlahMendali int

	if kriteriaC == 1 {
		fmt.Println("Masukkan jenis & jumlah mendali yang ingin dicari (Contoh : Emas 10):")
		fmt.Print("> ")
		fmt.Scan(&jenisMendali, &jumlahMendali)

		if desc == "Descending" {
			cetakCariNegaraD(N, jenisMendali, jumlahMendali)
		} else if desc == "Ascending" {
			cetakCariNegaraA(N, jenisMendali, jumlahMendali)
		}

	} else if kriteriaC == 2 {
		fmt.Println("Kembali ke menu utama...")
	}
}

// fungsi untuk mencari negara berdasarkan kriteria mendali tertentu pada urutan descending
func cariNegaraDescending(N daftarNegara, jenis string, target int) int {
	var cariMendali, i int
	var kiri, kanan, tengah int
	var cekMendali int = target
	var found int = -1

	if strings.EqualFold(jenis, "emas") {
		kiri = 0
		kanan = jumlahNegara - 1

		for kiri <= kanan && found == -1 {
			tengah = (kiri + kanan) / 2
			cariMendali = N[tengah].Emas
			if cariMendali == target {
				found = tengah
			} else if cariMendali < target {
				kanan = tengah - 1
			} else if cariMendali > target {
				kiri = tengah + 1
			}
		}

		if found != -1 {
			cekMendali = target
			for found > 0 && cekMendali == target {
				cekMendali = N[found-1].Emas
				if cekMendali == target {
					found--
				}
			}
		}
		return found
	}

	i = 0
	for i < jumlahNegara && found == -1 {
		if strings.EqualFold(jenis, "perak") {
			cariMendali = N[i].Perak
		} else if strings.EqualFold(jenis, "perunggu") {
			cariMendali = N[i].Perunggu
		}

		if cariMendali == target {
			found = i
		}
		i++
	}
	return found
}

// fungsi untuk mencari negara berdasarkan kriteria mendali tertentu pada urutan ascending
func cariNegaraAscending(N daftarNegara, jenis string, target int) int {
	var cariMendali int
	var kiri, kanan, tengah int
	var found int = -1
	kiri = 0
	kanan = jumlahNegara - 1
	tengah = (kiri + kanan) / 2

	if strings.EqualFold(jenis, "emas") {
		cariMendali = N[tengah].Emas
	} else if strings.EqualFold(jenis, "perak") {
		cariMendali = N[tengah].Perak
	} else if strings.EqualFold(jenis, "perunggu") {
		cariMendali = N[tengah].Perunggu
	}

	for kiri <= kanan && found == -1 {
		if cariMendali == target {
			found = tengah
			kanan = tengah - 1
		} else if cariMendali < target {
			kiri = tengah + 1
		} else if cariMendali > target {
			kanan = tengah - 1
		}
	}
	return found
}

// fungsi untuk mencetak negara yang ditemukan berdasarkan kriteria mendali tertentu pada urutan descending
func cetakCariNegaraD(N daftarNegara, jenis string, target int) {
	var idxAwal, i, j int
	var jumlahTemuan int = 0
	var cariMendali int
	var ketemuKriteria bool = false
	var nomerUrut int = 1

	if strings.EqualFold(jenis, "emas") {
		idxAwal = cariNegaraDescending(N, jenis, target)
		if idxAwal != -1 {
			i = idxAwal
			cariMendali = N[i].Emas

			// Hitung jumlah negara dengan emas yang sama (pasti berurutan)
			for i < jumlahNegara && cariMendali == target {
				jumlahTemuan++
				i++
				if i < jumlahNegara {
					cariMendali = N[i].Emas
				}
			}

			// Cetak Tabel
			fmt.Printf("| %-3s | %-15s | %-5s | %-5s | %-5s |\n", "No", "Negara", "Emas", "Perak", "Perunggu")
			for j = idxAwal; j < idxAwal+jumlahTemuan; j++ {
				fmt.Printf("| %-3d | %-15s | %-6d | %-6d | %-6d |\n", j-idxAwal+1, N[j].Nama, N[j].Emas, N[j].Perak, N[j].Perunggu)
			}
		} else {
			fmt.Printf("\nTidak ditemukan negara dengan %d mendali %s.\n", target, jenis)
		}
	} else {
		for i = 0; i < jumlahNegara; i++ {
			if strings.EqualFold(jenis, "perak") {
				cariMendali = N[i].Perak
			} else if strings.EqualFold(jenis, "perunggu") {
				cariMendali = N[i].Perunggu
			}

			if cariMendali == target {
				if !ketemuKriteria {
					fmt.Printf("| %-3s | %-15s | %-5s | %-5s | %-5s |\n", "No", "Negara", "Emas", "Perak", "Perunggu")
					ketemuKriteria = true
				}
				fmt.Printf("| %-3d | %-15s | %-6d | %-6d | %-6d |\n", nomerUrut, N[i].Nama, N[i].Emas, N[i].Perak, N[i].Perunggu)
				nomerUrut++
			}
		}
		if !ketemuKriteria {
			fmt.Printf("\nTidak ditemukan negara dengan %d mendali %s.\n", target, jenis)
		}
	}
}

// fungsi untuk mencetak negara yang ditemukan berdasarkan kriteria mendali tertentu pada urutan ascending
func cetakCariNegaraA(N daftarNegara, jenis string, target int) {
	var idxAwal, i, j int
	var jumlahTemuan int = 0
	var cariMendali int

	if strings.EqualFold(jenis, "emas") {
		cariMendali = N[i].Emas
	} else if strings.EqualFold(jenis, "perak") {
		cariMendali = N[i].Perak
	} else if strings.EqualFold(jenis, "perunggu") {
		cariMendali = N[i].Perunggu
	}

	idxAwal = cariNegaraAscending(N, jenis, target)
	if idxAwal != -1 {
		i = idxAwal
		for i < jumlahNegara && cariMendali == target {
			jumlahTemuan++
			i++
		}

		for j = idxAwal; j < idxAwal+jumlahTemuan; j++ {
			fmt.Printf("| %-3s | %-15s | %-5s | %-5s | %-5s |\n", "No", "Negara", "Emas", "Perak", "Perunggu")
			fmt.Printf("| %-3d | %-15s | %-6d | %-6d | %-6d |\n", j+1, N[j].Nama, N[j].Emas, N[j].Perak, N[j].Perunggu)
		}
	}

}

// fungsi untuk menganalisis kebutuhan mendali suatu negara untuk menjadi juara 1
func analisisKemenangan(N daftarNegara, nama string) {
	var juara1, target negara
	var targetIdx int = -1
	var i int = 0
	var goals, goalsPerak, goalsPerunggu int
	nama = strings.ToLower(nama)

	for i < jumlahNegara && targetIdx == -1 {
		if strings.ToLower(N[i].Nama) == nama {
			targetIdx = i
		}
		i++
	}

	if targetIdx == -1 {
		fmt.Printf("\nGagal: Negara '%s' Negara tidak ditemukan dalam klasemen.\n", nama)
	} else if targetIdx == 0 {
		fmt.Printf("%s sudah berada di posisi utama (Juara 1)!\n", N[0].Nama)
	} else if targetIdx > 0 {
		juara1 = N[0]
		target = N[targetIdx]

		fmt.Println()
		fmt.Printf("ANALISIS KEBUTUHAN MENDALI %s UNTUK MENJADI JUARA 1\n", strings.ToUpper(target.Nama))
		fmt.Printf("Peringkat Saat Ini : %d\n", targetIdx+1)
		fmt.Println("---------------------------------------------------------")

		if target.Emas < juara1.Emas {
			goals = juara1.Emas - target.Emas
			fmt.Printf("Untuk menggeser %s, %s membutuhkan minimal :\n", juara1.Nama, target.Nama)
			fmt.Printf("> %+d Mendali Emas (Tanpa Memperdulikan Perak / Perunggu)\n", goals)
			fmt.Printf("Sehingga Total Mendali Emas Menjadi %d (Menyamai %s, Namun Unggul Urutan / Kriteria lain\n)", juara1.Emas, juara1.Nama)
			fmt.Printf("> Atau %+d Mendali Emas untuk Menang Mutlak.\n", goals+1)
		} else if target.Emas == juara1.Emas {
			if target.Perak < juara1.Perak {
				goalsPerak = juara1.Perak - target.Perak
				fmt.Printf("Jumlah Emas sudah sama dengan %s. %s membutuhkan:\n", juara1.Nama, target.Nama)
				fmt.Printf("> %+d Medali Perak tambahan untuk menyalip.\n", goalsPerak+1)
			} else if target.Perak == juara1.Perak {
				goalsPerunggu = juara1.Perunggu - target.Perunggu
				fmt.Printf("Jumlah Emas dan Perak sudah sama dengan %s. %s membutuhkan:\n", juara1.Nama, target.Nama)
				fmt.Printf("> %+d Medali Perunggu tambahan untuk menyalip.\n", goalsPerunggu+1)
			}
		}
		fmt.Println("---------------------------------------------------------")
		fmt.Println()
	}
}

// fungsi untuk menambahkan negara baru ke dalam daftar
func tambahNegara(N *daftarNegara, nama string) {
	var emas, perak, perunggu int

	if jumlahNegara >= NMAX {
		fmt.Println("\nGagal: Kapasitas maksimum negara peserta telah penuh!")
	} else if cariNegaraIndex(*N, nama) != -1 {
		fmt.Printf("\nGagal: Negara '%s' sudah terdaftar.\n", nama)
	} else {
		fmt.Println("Masukkan jumlah medali emas, perak, dan perunggu:")
		fmt.Print("> ")
		fmt.Scan(&emas, &perak, &perunggu)
		N[jumlahNegara] = negara{Nama: nama, Emas: emas, Perak: perak, Perunggu: perunggu}
		jumlahNegara++
		fmt.Printf("\nBerhasil: Negara '%s' berhasil ditambahkan.\n", nama)
	}
}

// fungsi untuk mengubah data negara
func ubahDataNegara(N *daftarNegara, nama string, index int) {
	var emasBaru, perakBaru, perungguBaru int

	fmt.Println("Masukkan jumlah medali emas, perak, dan perunggu yang baru:")
	fmt.Print("Jumlah medali emas: ")
	fmt.Scan(&emasBaru)
	N[index].Emas = emasBaru
	fmt.Print("Jumlah medali perak: ")
	fmt.Scan(&perakBaru)
	N[index].Perak = perakBaru
	fmt.Print("Jumlah medali perunggu: ")
	fmt.Scan(&perungguBaru)
	N[index].Perunggu = perungguBaru
	fmt.Printf("\nBerhasil: Data medali untuk negara '%s' berhasil diubah.\n", nama)

}

// Fungsi untuk mengubah data mendali yang spesifik
func ubahDataMedali(N *daftarNegara, nama string, index int) {
	var jumlahBaru int
	var mendaliDiCari string

	fmt.Println("Masukkan jenis medali yang ingin diubah (emas/perak/perunggu):")
	fmt.Print("> ")
	fmt.Scan(&mendaliDiCari)
	fmt.Println("Masukkan jumlah medali yang baru:")
	fmt.Print("> ")
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
func hapusNegara(N *daftarNegara, nama string) {
	var index int = cariNegaraIndex(*N, nama)
	var i int = index

	if index > -1 {
		fmt.Printf("\nGagal: Negara '%s' tidak ditemukan.\n", nama)

		for i < jumlahNegara-1 {
			N[i] = N[i+1]
			i++
		}

		N[jumlahNegara-1] = negara{}
		jumlahNegara--
		fmt.Printf("\nBerhasil: Negara '%s' berhasil dihapus.\n", nama)
	} else {
		fmt.Printf("\nGagal: Negara '%s' tidak ditemukan.\n", nama)
	}
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
	if jumlahnegara > 0 {
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
	} else {
		fmt.Println("\nKlasemen kosong. Tidak ada negara yang terdaftar.")
		fmt.Println()
	}

}
