package main

import "fmt"

func main() {
	var angka int
	isFound := false

	fmt.Print("Masukkan bilangan yg mau di cek prima = ")
	fmt.Scanln(&angka)

	if angka == 2 {
		fmt.Println(angka, "dua  bilangan prima")
	} else if angka%2 == 0 || angka <= 1 {
		fmt.Println(angka, "satu bukan bilangan prima")
	} else {
		for temp := 3; temp*temp <= angka; temp += 2 {
			if angka%temp == 0 {
				isFound = true
				break
			}
		}
		if isFound == false {
			fmt.Println(angka, " adalah bilangan prima")
		} else {
			fmt.Println(angka, " bukan bilangan prima")
		}
	}
}
