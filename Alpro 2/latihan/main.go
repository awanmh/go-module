package main

import "fmt"

type arrInt [5]int

func selectionSort(T *arrInt, n int) {
    var t, i, j, idx_min int

    i = 1

    for i <= n-1 {
        idx_min = i - 1
        j = i

        for j < n {
            if T[idx_min] > T[j] {
                idx_min = j
            }
            j++
        }
        t = T[idx_min]
        T[idx_min] = T[i-1]
        T[i-1] = t

        i++
    }
}

func main() {
    data := arrInt{7, 2, 9, 1, 5}
    selectionSort(&data, 5)
    fmt.Println(data)
}