package main

import (
	"fmt"
)

func main() {
	pinjaman := 10000000.0
	bunga := (12 / 12) * 0.1 * pinjaman
	cicilan := (pinjaman + bunga) / 12

	fmt.Printf("Cicilan Apip : %.2f", cicilan)
}

// ===== Soal Satu =====
/*
	Apip meminjam uang dikoperasi sebesar 10.000.000, koperasi memberikan bunga sebesar 10% perbulan.
	Apip memilih cicilan selama 12 bulan, berapa cicilan perbulan?
	buatkan perhitungan kode dalam golang
*/
