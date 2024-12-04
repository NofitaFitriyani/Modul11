package main

import "fmt"

func main() {
	var suara = []int{7, 19, 3, 2, 78, 3, 1, -3, 18, 19, 0}
	var totalSuaraMasuk, totalSuaraSah int

	var hitungSuara [21]int

	for _, s := range suara {
		if s == 0 {
			break
		}
		totalSuaraMasuk++ 

		if s >= 1 && s <= 20 {
			hitungSuara[s]++
			totalSuaraSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitungSuara[i])
		}
	}
}
