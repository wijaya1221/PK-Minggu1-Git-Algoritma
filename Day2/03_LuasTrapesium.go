package main

import "fmt"

func main() {
	var luas, a, b, t int

	fmt.Print("Masukkan nilai A = ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan nilai B = ")
	fmt.Scanln(&b)

	fmt.Print("Masukkan nilai Tinggi = ")
	fmt.Scanln(&t)

	luas = (a + b) * t / 2
	fmt.Println(luas, " adalah luas trapesium")

}
