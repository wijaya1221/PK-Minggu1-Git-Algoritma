package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	counts := make(map[int]int)

	// Menghitung kemunculan setiap digit dalam angka
	for _, char := range angka {
		digit, _ := strconv.Atoi(string(char))
		counts[digit]++
	}

	// Membuat slice untuk menyimpan angka yang muncul sekali
	var result []int
	for _, char := range angka {
		digit, _ := strconv.Atoi(string(char))
		if counts[digit] == 1 {
			result = append(result, digit)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
