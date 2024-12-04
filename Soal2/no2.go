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

	var ketua, wakil int
	var suaraKetua, suaraWakil int

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > suaraKetua {
			wakil = ketua
			suaraWakil = suaraKetua
			ketua = i
			suaraKetua = hitungSuara[i]
		} else if hitungSuara[i] > suaraWakil {
			wakil = i
			suaraWakil = hitungSuara[i]
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil Ketua RT: %d\n", wakil)
}
