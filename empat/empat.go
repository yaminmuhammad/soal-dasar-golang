package main

import "fmt"

func main() {
	tinggi := 5

	for i := 0; i < tinggi; i++ {
		for j := 0; j < tinggi-i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}

// ===== Soal Empat =====
/*
	buatkan segitiga terbalik

*/
