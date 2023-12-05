package main

import "fmt"

func hitungBiayaPerjalanan(jarak, konsumsiBensin, hargaBensin float64) float64 {

	// Menghitung jumlah liter bensin yang digunakan
	literBensin := jarak / konsumsiBensin

	// Menghitung biaya total perjalanan
	biayaTotal := literBensin * hargaBensin

	return biayaTotal
}

func main() {
	// Parameter perjalanan Yayat
	jarak := 63.0          // dalam kilometer
	konsumsiBensin := 17.0 // dalam kilometer per liter
	hargaBensin := 8350.0  // dalam rupiah per liter

	// Menghitung biaya perjalanan
	biayaPerjalanan := hitungBiayaPerjalanan(jarak, konsumsiBensin, hargaBensin)

	// Menampilkan hasil
	fmt.Printf("Biaya perjalanan Yayat: Rp%.f\n", biayaPerjalanan)
}

// ===== Soal Tiga =====
/*
	yayat akan pulang kampung ke daerah serang, dia akan menempuh jarak 63km
	dia menggunakan sepeda motor untuk sampai ke serang , motor yang ia gunakan
	menghabiskan bensin 17km/liter, berapakah biaya yang dikeluarkan oleh yayat
	jika harga bensin 8.350/liter? buatkan koding perhitungan dengan golang
*/
