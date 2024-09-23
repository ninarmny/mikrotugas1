package main

import "fmt"

func main() {

	// deklarasi variabel
	var nama_customer, nama_barang string
	var harga float32 = 25000
	var quantity int32
	var hasil_diskon, total_harga float32

	// input nama customer
	fmt.Print("Input Nama Customer: ")
	fmt.Scanf("%s", &nama_customer)

	// input nama barang
	fmt.Print("Input Nama Barang: ")
	fmt.Scanf("%s", &nama_barang)

	// input quantity barang
	fmt.Print("Input Quantity Barang: ")
	fmt.Scanf("%d", &quantity)

	// kondisi diskon, kalau lebih dari 5 dapat 10%, selain itu 2%
	switch quantity {
	case 5: 
		hasil_diskon = 0.1 // 10%
	default:
		hasil_diskon = 0.2 // 2%
	}

	// proses
	sub_total := float32(quantity) * harga
	total_harga = sub_total - (hasil_diskon * sub_total)

	// tampilkan hasil harga
	fmt.Println("hasil_diskon: ", hasil dikson)
	fmt.Println("quantity: ", quantity)
	fmt.Println("harga: ", harga)
	fmt.Println("Total Harga adalah: ", total_harga)

}