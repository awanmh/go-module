# MODUL 6: REKURSIF (Recursion)

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami konsep rekursif
* Menentukan base-case & recursive-case
* Membuat program rekursif di Golang
* Memahami alur forward & backward recursion

---

## 6.1 Pengantar Rekursif

### Apa itu Rekursif?

**Rekursif** adalah teknik pemrograman di mana:

> Sebuah fungsi/prosedur memanggil dirinya sendiri

flownya :
```
        ┌──────────────┐
        │    INPUT     │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │  Panggil     │
        │  Fungsi      │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │  Base Case?  │─── Ya ──► Kembalikan nilai
        └───────┬──────┘
          Tidak │
                ▼
        ┌──────────────┐
        │ Recursive    │
        │ Case         │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Forward Call │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Backward     │
        │ Return       │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │    OUTPUT    │
        └──────────────┘
```
---

### Konsep Inti

`Masalah besar → dipecah jadi masalah kecil → diselesaikan berulang`

---

### Analogi Sederhana

Bayangkan:

> 📦 Kotak di dalam kotak di dalam kotak…

Untuk membuka:

* Buka kotak luar
* Lanjut ke kotak di dalam
* Terus sampai kotak terakhir

Itulah rekursif

---

### Contoh Rekursif Tanpa Henti (SALAH)

```go
package main
import "fmt"

func cetak(x int) {
    fmt.Println(x)
    cetak(x + 1) // memanggil diri sendiri tanpa berhenti
}

func main() {
    cetak(5)
}
```

Output:

```
5 6 7 8 9 10 11 ...
```

Tidak berhenti → **Infinite Recursion**

---

## Masalah Utama Rekursif

Tanpa kondisi berhenti → program crash (stack overflow)

---

## Solusi: Base Case

### Base Case

Kondisi untuk **menghentikan rekursi**

---

### Contoh Perbaikan

```go
package main
import "fmt"

func cetak(x int) {
    if x == 10 { // base-case
        fmt.Println(x)
    } else {
        fmt.Println(x)
        cetak(x + 1) // recursive-case
    }
}

func main() {
    cetak(5)
}
```

Output:

```
5 6 7 8 9 10
```

---

## Komponen Rekursif

### 1. Base Case (Wajib!)

Kondisi berhenti

```go
if x == 10 {
    return
}
```

---

### 2. Recursive Case

Pemanggilan dirinya sendiri

```go
cetak(x + 1)
```

---

### Aturan Emas Rekursif

1. Harus ada **base-case**
2. Harus ada **perubahan menuju base-case**
3. Tidak boleh infinite loop

---

## 6.3 Forward & Backward Recursion

### Konsep Penting

#### Forward (Turun)

- Saat fungsi **memanggil dirinya sendiri**

#### Backward (Naik)

- Saat fungsi **selesai & kembali**

---

### Contoh 1: Forward Output

```go
package main
import "fmt"

func cetak(x int) {
    if x == 5 {
        fmt.Println(x)
    } else {
        fmt.Println(x) // cetak dulu
        cetak(x + 1)
    }
}

func main() {
    cetak(1)
}
```

Output:

```
1 2 3 4 5
```

---

### Contoh 2: Backward Output

```go
package main
import "fmt"

func cetak(x int) {
    if x == 5 {
        fmt.Println(x)
    } else {
        cetak(x + 1)
        fmt.Println(x) // cetak setelah rekursi
    }
}

func main() {
    cetak(1)
}
```

📌 Output:

```
5 4 3 2 1
```

---

### Insight Penting

| Posisi Print    | Hasil      |
| --------------- | ---------- |
| Sebelum rekursi | Urut naik  |
| Setelah rekursi | Urut turun |

---

## 6.4 Contoh Program Rekursif

---

## A. Menampilkan n hingga 1

#### Base-case: n == 1

```go
package main
import "fmt"

func baris(n int) {
    if n == 1 {
        fmt.Println(1)
    } else {
        fmt.Println(n)
        baris(n - 1)
    }
}

func main() {
    baris(5)
}
```

📌 Output:

```
5 4 3 2 1
```

---

### B. Penjumlahan 1 sampai n

### Base-case: n == 1

```go
package main
import "fmt"

func penjumlahan(n int) int {
    if n == 1 {
        return 1
    } else {
        return n + penjumlahan(n-1)
    }
}

func main() {
    fmt.Println(penjumlahan(5))
}
```

Output:

```
15
```

---

### C. Pangkat 2ⁿ

#### Base-case: n == 0

```go
package main
import "fmt"

func pangkat(n int) int {
    if n == 0 {
        return 1
    } else {
        return 2 * pangkat(n-1)
    }
}

func main() {
    fmt.Println(pangkat(4))
}
```

Output:

```
16
```

---

### D. Faktorial (n!)

#### Base-case: n == 0 atau 1

```go
package main
import "fmt"

func faktorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    } else {
        return n * faktorial(n-1)
    }
}

func main() {
    fmt.Println(faktorial(5))
}
```

Output:

```
120
```

---

### Visualisasi Alur Faktorial

Contoh:

```
faktorial(3)
= 3 * faktorial(2)
= 3 * (2 * faktorial(1))
= 3 * (2 * 1)
= 6
```

---

### Rekursif vs Iteratif

| Rekursif            | Iteratif           |
| ------------------- | ------------------ |
| Lebih sederhana     | Lebih cepat        |
| Lebih elegan        | Lebih hemat memori |
| Bisa stack overflow | Aman               |

---

## Kesimpulan

* Rekursif = fungsi memanggil dirinya sendiri
* WAJIB punya:

  * Base-case
  * Recursive-case
* Posisi kode menentukan hasil (naik / turun)
* Semua rekursif bisa diubah ke looping

### Gunakan rekursif untuk:

* Tree / Graph traversal
* Backtracking (DFS, A*)
* Divide & Conquer
