package main

import "fmt"

// Fungsi mergeArrays menggabungkan dua array tanpa duplikasi.
func ArrayMerge(arrayA []string, arrayB []string) []string {
	result := []string{}

	// Menggunakan map untuk melacak keberadaan elemen
	seen := make(map[string]bool)

	// Menggabungkan elemen-elemen dari arrayA ke result
	for _, i := range arrayA {
		if !seen[i] {
			result = append(result, i)
			seen[i] = true
		}
	}

	// Menggabungkan elemen-elemen dari arrayB ke result
	for _, i := range arrayB {
		if !seen[i] {
			result = append(result, i)
			seen[i] = true
		}
	}

	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
