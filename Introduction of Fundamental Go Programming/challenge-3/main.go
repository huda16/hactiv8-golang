package main

import (
    "fmt"
    "strings"
)

func main() {
    sentence := "selamat malam"

	parsingChar(sentence)
}

func parsingChar(sentence string) {
	// Pecah kalimat menjadi karakter-karakter dengan fungsi strings.Split
    chars := strings.Split(sentence, "")

    // Membuat map kosong untuk menghitung kemunculan karakter
    charCount := make(map[string]int)

    // Looping setiap karakter
    for _, char := range chars {
		// Tampilkan setiap karakter
		fmt.Printf("%s \n", char)

        // Jika karakter sudah ada di map, tambahkan jumlah kemunculan karakter
        if _, ok := charCount[char]; ok {
            charCount[char]++
        } else { // Jika karakter belum ada di map, buat entry baru dengan jumlah 1
            charCount[char] = 1
        }
    }

    // Tampilkan hasil perhitungan kemunculan karakter
	fmt.Println(charCount)
}