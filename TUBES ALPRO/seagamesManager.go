package main

import "fmt"

type negara struct {
	Nama string
	Emas int
	Perak int
	Perunggu int
}

var daftarNegara []negara

func menu() {
	fmt.Println("------------------------------------")
	fmt.Println("          SEAGAMES MANAGER          ")
	fmt.Println("------------------------------------")
	fmt.Println("1. Tambah atau Ubah Negara & Mendali")
	fmt.Println("2. Hapus Negara")
	fmt.Println("3. Tampilkan Pringkat Negara")
	fmt.Println("4. Keluar")
	fmt.Println("------------------------------------")
}


func main(){
	var pilihan int

	for{
	menu()
	fmt.Println("Pilih 1/2/3/4 ?")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
	case 2:
	case 3:
	case 4:
		fmt.Println("Terima kasih telah menggunakan SEAGAMES MANAGER")
		return
	default:
		fmt.Println("Pilihan tidak valid, silakan pilih 1/2/3/4\n")
	}

	if pilihan == 4 {
		break
	}
	}
	
}

func inputData() {
	var nama string
	var emas, perak, perunggu, foundIndex, i int

	fmt.Println("Masukkan nama negara:")
	fmt.Scan(&nama)
	fmt.Println("Masukkan jumlah medali emas:")
	fmt.Scan(&emas)
	fmt.Println("Masukkan jumlah medali perak:")
	fmt.Scan(&perak)
	fmt.Println("Masukkan jumlah medali perunggu:")
	fmt.Scan(&perunggu)

	foundIndex = -1
	for i = 0; i < len(daftarNegara); i++ {
		if daftarNegara[i].Nama == nama {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		daftarNegara[foundIndex].Emas = emas
		daftarNegara[foundIndex].Perak = perak
		daftarNegara[foundIndex].Perunggu = perunggu
		fmt.Println("Data negara berhasil diperbarui.")
	} else {
		daftarNegara = append(daftarNegara, negara{Nama: nama, Emas: emas, Perak: perak, Perunggu: perunggu})
		fmt.Println("Data negara berhasil ditambahkan.")
	}

}

func hapusNegara() {
	var nama string
	var i int
	fmt.Println("Masukkan nama negara yang ingin dihapus:")
	fmt.Scan(&nama)

	for i = 0; i < len(daftarNegara); i++ {
		if daftarNegara[i].Nama == nama {
			daftarNegara = append(daftarNegara[:i], daftarNegara[i+1:]...)
			fmt.Println("Data negara berhasil dihapus.")
			return
		}
	}
	fmt.Println("Negara tidak ditemukan.")
}

func peringkat(){
	var n, i, j int
	n = len(daftarNegara)

	if n == 0 {
		fmt.Println("Belum ada data negara yang dimasukkan.")
		return
	}

	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			
		}
	}
}
