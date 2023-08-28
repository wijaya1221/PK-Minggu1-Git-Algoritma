package main

import (
	"fmt"
	"strings"
)

func Mapping(slice []string) string {
	counts := make(map[string]int)

	for _, elem := range slice {
		counts[elem]++
	}

	// Mengubah format hasil langsung dalam fungsi Mapping
	var parts []string
	for key, value := range counts {
		parts = append(parts, fmt.Sprintf("%s:%d", key, value))
	}
	return "map[" + strings.Join(parts, " ") + "]"
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	fmt.Println(Mapping([]string{}))
}
