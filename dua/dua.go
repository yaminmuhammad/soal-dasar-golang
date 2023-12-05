package main

import "fmt"

func main() {
	// Bobot untuk masing-masing komponen
	bobotAbsensi := 0.1
	bobotTugasHarian := 0.2
	bobotUTS := 0.3
	bobotUAS := 0.4

	// Nilai untuk masing-masing komponen
	nilaiAbsensi := 16.0 * bobotAbsensi
	nilaiTugasHarian := 70.0 * bobotTugasHarian
	nilaiUTS := 80.0 * bobotUTS
	nilaiUAS := 70.0 * bobotUAS

	// Menghitung total nilai
	totalNilai := nilaiAbsensi + nilaiTugasHarian + nilaiUTS + nilaiUAS

	// Menampilkan hasil
	fmt.Printf("Nilai Apip: %.2f\n", totalNilai)
}

// ===== Soal Dua =====
/*
	Alan merupakan seorang dosen di STMIK IB, dia ingin memberikan nilai
	kepada apip, dengan nilai absensi 16xhadir, nilai tugas harian = 70
	nilai uts = 80 dan nilai uas = 70, dengan bobot nilai absensi
	10%(untuk 18 x hadir), nilai tugas harian=20%, nilai uts = 30%
	dan nilai uas = 40%. buatkan perhitungan kode dengan golang
*/
