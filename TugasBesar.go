package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Username  string
	Password  string
	Saldo     float64
	Transaksi []string
}

// Data mahasiswa
var users = []User{
	{Username: "Nina", Password: "2406499052", Saldo: 3500000},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selamat Datang di ATM!")
	user := login(reader)

	// Menu pilihan
	for {
		fmt.Println("\nPilih Menu:")
		fmt.Println("1. Lihat Informasi Akun")
		fmt.Println("2. Lihat Saldo")
		fmt.Println("3. Tambahkan Saldo")
		fmt.Println("4. Tarik Saldo")
		fmt.Println("5. Lihat Histori Transaksi")
		fmt.Println("6. Keluar dari Program")

		fmt.Print("Masukkan pilihan: ")
		inputMenu, _ := reader.ReadString('\n')
		inputMenu = strings.TrimSpace(inputMenu)

		switch inputMenu {
		case "1":
			lihatInformasiAkun(user)
		case "2":
			lihatSaldo(user)
		case "3":
			tambahSaldo(user, reader)
		case "4":
			tarikSaldo(user, reader)
		case "5":
			lihatHistoriTransaksi(user)
		case "6":
			fmt.Println("Terima kasih sudah menggunakan ATM!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// Fungsi login
func login(reader *bufio.Reader) *User {
	for {
		fmt.Print("Masukkan Username (Nama Depan): ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Masukkan Password (NPM): ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		for i := range users {
			if users[i].Username == username && users[i].Password == password {
				fmt.Println("Login berhasil!")
				return &users[i]
			}
		}

		fmt.Println("Username atau password salah. Coba lagi.")
	}
}

// Fungsi untuk melihat informasi akun
func lihatInformasiAkun(user *User) {
	fmt.Println("\nInformasi Akun:")
	fmt.Println("Username:", user.Username)
	fmt.Println("Password:", user.Password)
}

// Fungsi untuk melihat saldo
func lihatSaldo(user *User) {
	fmt.Println("\nSaldo Anda saat ini adalah:", user.Saldo)
}

// Fungsi untuk menambah saldo
func tambahSaldo(user *User, reader *bufio.Reader) {
	fmt.Print("\nMasukkan jumlah saldo yang ingin ditambahkan: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	jumlah, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid.")
		return
	}

	user.Saldo += jumlah
	user.Transaksi = append(user.Transaksi, "Tambah saldo: "+strconv.FormatFloat(jumlah, 'f', 2, 64))
	fmt.Println("Saldo berhasil ditambahkan!")
}

// Fungsi untuk menarik sald0
func tarikSaldo(user *User, reader *bufio.Reader) {
	fmt.Print("\nMasukkan jumlah saldo yang ingin ditarik: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	jumlah, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid.")
		return
	}

	if jumlah > user.Saldo {
		fmt.Println("Saldo tidak mencukupi.")
		return
	}

	user.Saldo -= jumlah
	user.Transaksi = append(user.Transaksi, "Tarik saldo: "+strconv.FormatFloat(jumlah, 'f', 2, 64))
	fmt.Println("Saldo berhasil ditarik!")
}

// Fungsi untuk melihat histori transaksi
func lihatHistoriTransaksi(user *User) {
	fmt.Println("\nHistori Transaksi:")
	if len(user.Transaksi) == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}
	for _, t := range user.Transaksi {
		fmt.Println(t)
	}
}
