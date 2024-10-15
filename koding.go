package main

import (
	"fmt"
)

type Buku struct {
	Nama   string
	Jumlah int
}

type Pengguna struct {
	Username string
	Password string
}

var daftarBuku = []Buku{
	{"Pemrograman", 10},
	{"Film", 5},
	{"Printing", 20},
}

var historiPeminjaman []string

func LihatInformasiPengguna() {
	fmt.Println("=== Informasi Pengguna Program ===")
	fmt.Println("Username              : Nina")
	fmt.Println("Password              : 2406499052")
	fmt.Println("Jenis Kelamin         : Perempuan")
	fmt.Println("Makanan Favorit       : Cookies")
	fmt.Println("Minuman Favorit  	   : Susu")
	fmt.Println("==================================")
}

func LihatDaftarBuku() {
	fmt.Println("=== Daftar Buku ===")
	for i, buku := range daftarBuku {
		fmt.Printf("%d. Nama Buku: %s\n   Jumlah: %d\n", i+1, buku.Nama, buku.Jumlah)
	}
	fmt.Println("===================")
}

func TambahDaftarBuku() {
	var nama string
	var jumlah int

	fmt.Println("=== Tambah Daftar Buku ===")
	fmt.Print("Nama Buku: ")
	fmt.Scan(&nama)
	fmt.Print("Jumlah: ")
	fmt.Scan(&jumlah)

	if jumlah > 0 {
		daftarBuku = append(daftarBuku, Buku{Nama: nama, Jumlah: jumlah})
		fmt.Println("Buku berhasil ditambahkan.")
	} else {
		fmt.Println("Jumlah buku tidak bisa kurang dari atau sama dengan 0.")
	}
}

func TambahPeminjamanBuku() {
	fmt.Println("=== Pinjam Buku ===")
	for i, buku := range daftarBuku {
		fmt.Printf("%d. %s (%d)\n", i+1, buku.Nama, buku.Jumlah)
	}

	var nomor, jumlah int
	fmt.Print("Masukkan Nomor Pinjaman Buku: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > len(daftarBuku) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	fmt.Print("Masukkan Jumlah Pinjaman: ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Jumlah pinjaman harus lebih dari 0.")
		return
	}

	buku := &daftarBuku[nomor-1]
	if jumlah > buku.Jumlah {
		fmt.Printf("Jumlah pinjaman melebihi stok. Tersedia: %d\n", buku.Jumlah)
		return
	}

	buku.Jumlah -= jumlah
	historiPeminjaman = append(historiPeminjaman, fmt.Sprintf("Dipinjam: %s, Jumlah: %d", buku.Nama, jumlah))
	fmt.Printf("Anda berhasil meminjam %d buku: %s\n", jumlah, buku.Nama)
}

func HistoriPeminjamanBuku() {
	fmt.Println("=== Histori Peminjaman Buku ===")
	if len(historiPeminjaman) == 0 {
		fmt.Println("Belum ada peminjaman.")
	} else {
		for _, histori := range historiPeminjaman {
			fmt.Println(histori)
		}
	}
	fmt.Println("===============================")
}

func KeluarDariProgram() {
	fmt.Println("Terima kasih telah menggunakan program peminjaman buku!")
}

func main() {
	fmt.Println("=== Selamat Datang di Perpustakaan Vokasi ===")

	var pengguna = Pengguna{"Nina", "2406499052"}
	var username, password string

	fmt.Print("Silahkan Input Username: ")
	fmt.Scan(&username)
	fmt.Print("Silahkan Input Password: ")
	fmt.Scan(&password)

	if username == pengguna.Username && password == pengguna.Password {
		fmt.Println("Login Sukses!")

		for {
			fmt.Println("=== Menu Program ===")
			fmt.Println("1. Lihat Informasi Pengguna Program")
			fmt.Println("2. Lihat Daftar Buku")
			fmt.Println("3. Tambah Daftar Buku")
			fmt.Println("4. Tambah Peminjaman Buku")
			fmt.Println("5. Histori Peminjaman Buku")
			fmt.Println("6. Keluar dari Program")
			fmt.Print("Input menu yang anda inginkan: ")

			var pilihan int
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				LihatInformasiPengguna()
			case 2:
				LihatDaftarBuku()
			case 3:
				TambahDaftarBuku()
			case 4:
				TambahPeminjamanBuku()
			case 5:
				HistoriPeminjamanBuku()
			case 6:
				KeluarDariProgram()
				return
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		}
	} else {
		fmt.Println("Username atau Password salah. Program dihentikan.")
	}
}
