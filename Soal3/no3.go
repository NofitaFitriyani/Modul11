package main

import "fmt"

const NMAX = 1000000
var data [NMAX]int

func posisiSequential(n, k int) int {
    for i := 0; i < n; i++ {
        if data[i] == k {
            return i
        }
    }
    return -1
}

func posisiBinary(n, k int) int {
    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        if data[mid] == k {
            return mid
        } else if data[mid] < k {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

func isiArray(n int) {
    for i := 0; i < n; i++ {
        fmt.Scan(&data[i])
    }
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    isiArray(n)

    posisi := posisiSequential(n, k)
    if posisi != -1 {
        fmt.Printf("Data berada pada posisi ke-%d\n", posisi)
    } else {
        fmt.Println("TIDAK ADA")
    }

    posisi = posisiBinary(n, k)
    if posisi != -1 {
        fmt.Printf("Data berada pada posisi ke-%d\n", posisi)
    } else {
        fmt.Println("TIDAK ADA")
    }
}
