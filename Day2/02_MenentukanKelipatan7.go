package main

import "fmt"

func main() {
	var angka, i int
	i = 7

	fmt.Print("Masukkan bilangan yg mau di cek kelipatan 7 = ")
	fmt.Scanln(&angka)

	if angka%i == 0 {
		fmt.Println(angka, " merupakan kelipatan dari 7")
	} else {
		fmt.Println(angka, " bukan kelipatan dari 7")
	}

}
