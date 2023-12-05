package main

import (
	"fmt"
	"sort"
)

func main() {
	umurAyu := 10

	// hitung umur cinta
	umurCinta := umurAyu + 5

	// hitung umur budi
	umurBudi := umurCinta + 2

	// hitung umur ibu
	umurIbu := 3*umurAyu + 2

	// hitung umur ayah
	umurAyah := umurCinta + umurBudi + umurAyu

	// cetak hasil
	fmt.Println("Umur Ayu adalah", umurAyu)
	fmt.Println("Umur Cinta:", umurCinta)
	fmt.Println("Umur Budi:", umurBudi)
	fmt.Println("Umur Ibu:", umurIbu)
	fmt.Println("Umur Ayah:", umurAyah)

	// membuat slice umur anak-anak
	umurAnak := []int{umurAyu, umurCinta, umurBudi}

	// Mengurutkan umur dari tertua hingga termuda
	sort.Sort(sort.Reverse(sort.IntSlice(umurAnak)))

	// Menampilkan hasil urutan
	fmt.Println("Urutan umur anak dari tertua hingga termuda:")
	for i, umur := range umurAnak {
		fmt.Printf("Anak ke-%d: %d tahun\n", i+1, umur)
	}
}

// ===== Soal Lima =====
/*
	Ayu, Budi dan Cinta merupakan kakak beradik. Budi 2 tahun diatas Cinta.
	sedangkan Ayu 5 tahun dibawah Cinta. Umur ayah dapat diketahui dengan menambahkan semua umur anaknya.
	sedangkan ibu berumur 3 kali lebih besar dari Ayu, setelah itu ditambah 2.
	Berapa umur cinta, jika ayu berumur 10 tahun?
	Berapa umur budi dan ibu, jika ayu berumur 10 tahun?
	Berapa umur ayah, jika Ayu berumur 10 tahun?
	Berdasarkan perhitungan diatas, urutkan berdasar anak tertua hingga termuda?
*/
