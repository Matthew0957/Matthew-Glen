package main

import "fmt"

func main() {
	var pilihan, jumlah int
	var bayar int
	total := 0

	barang := []string{"pensil", "Buku","penghapus"}
	harga := []int{2000,5000,3000,1500}

	fmt.println("===== POINT OF SALES =====")

	FOR {
		fmt.Println(
"Daftar Barang:")
		fmt.Println("------------------------")
		fmt.Println("No / Nama Barang / Harga")
		fmt.Println("------------------------")
		for i := 0; i < len(barang); i++ {
			fmt.Printf("%d. %-12s Rp %d/n", i+1, barang[i], harga[i])
		}
		fmt.Println("0 . selesai")
		fmt.Println("------------------------")

		fmt.Println("Pilih barang :")
		fmt.Scan(%pilihan)

		if pilihan == 0 {
			break
		}

		fmt.Print("Jumlah :")
		fmt.Scan(%jumlah)

		subtotal := harga[pilihan-1] * jumlah
		total += subtotal
		fmt.Println("Subtotal ditambahkan")
  }

     fmt.Println("
	 =====STRUK BELANJA=====")
	 fmt.Println("------------------------")
	 fmt.Printf("Total Belanja : Rp%d", total)

	 kembalian := bayar - total
	 fmt,Printf("Kembalian : Rp%d", kembalian)
	 fmt.Println("------------------------")
	 fmt.Println("Terima Kasih telah berbelanja")
}
