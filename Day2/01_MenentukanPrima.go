package main

import "fmt"

func main() {
	var angka, i int
	i = 2

	fmt.Print("Masukkan bilangan yg mau di cek prima = ")
	fmt.Scanln(&angka)

	for temp := 2; temp < angka; temp++ {
		if angka%temp == 0 {
			break
		}
		i++
	}

	if angka == i {
		fmt.Println(angka, " adalah bilangan prima")
	} else {
		fmt.Println(angka, " bukan bilangan prima")
	}

}
