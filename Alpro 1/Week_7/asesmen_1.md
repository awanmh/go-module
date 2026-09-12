# MODUL 7: ASESMEN PRAKTIKUM 1 (REVIEW & EVALUASI MANDIRI)

Modul ini berisi **evaluasi kompetensi mandiri tahap pertama**, mencakup seluruh materi dasar: I/O standar, tipe data, variabel, operator aritmatika/modulo/casting, dan perulangan `for`.

### Tujuan Pembelajaran

Mahasiswa mampu:

* Mengintegrasikan konsep I/O, variabel, operator aritmatika, modulo, dan perulangan secara simultan
* Menganalisis spesifikasi kebutuhan persoalan (*problem statement*) dan merumuskan algoritma solusi
* Menerapkan prinsip *clean code*, indentasi terstandar, dan efisiensi eksekusi
* Menyelesaikan masalah pemrograman di laboratorium secara mandiri dan tepat waktu

---

## 7.1 Karakteristik Asesmen Praktikum 1

Pada sesi Asesmen Praktikum 1:
1. Praktikan mengerjakan soal secara mandiri di komputer laboratorium.
2. Tidak ada bantuan ide algoritma dari asisten praktikum.
3. Pertanyaan hanya diperkenankan terkait kejelasan redaksional soal atau kendala teknis perangkat keras/lingkungan kompiler.

---

## 7.2 Pembahasan Contoh Kasus Asesmen 1

### Kasus 1 – Penjumlahan Rentang $[x, y]$
Buatlah program yang digunakan untuk menjumlahkan semua bilangan bulat dari $x$ sampai dengan $y$, dengan jaminan nilai $x \le y$.

* **Masukan**: Dua buah bilangan bulat positif $x$ dan $y$.
* **Keluaran**: Satu bilangan bulat hasil penjumlahan seluruh nilai dari $x$ sampai $y$ (inklusif).

#### Solusi 1: Pendekatan Iteratif (Perulangan For)
```go
package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    var jumlah int = 0
    for i := x; i <= y; i++ {
        jumlah = jumlah + i
    }

    fmt.Println(jumlah)
}
```

#### Solusi 2: Pendekatan Matematis Deret Aritmatika ($O(1)$)
Rumus deret aritmatika dengan suku pertama $x$, suku terakhir $y$, dan banyak suku $n = y - x + 1$:
$$S = \frac{n}{2} (x + y)$$

```go
package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    n := y - x + 1
    jumlah := (n * (x + y)) / 2

    fmt.Println(jumlah)
}
```

#### Contoh Uji:
| No | Masukan | Keluaran | Penjelasan |
| :--- | :--- | :--- | :--- |
| 1 | `4 5` | `9` | $4 + 5 = 9$ |
| 2 | `2 100` | `5049` | Penjumlahan 2 s.d. 100 |

---

### Kasus 2 – Program Kupon dan Diskon Belanja
Sebuah toko memberikan 1 kupon undian untuk setiap kelipatan belanja Rp 50.000,-. Jika total belanjaan mencapai Rp 100.000,- atau lebih, pembeli juga mendapatkan diskon langsung sebesar 10% dari total belanjanya.
* **Masukan**: Satu bilangan bulat positif menyatakan total belanja awal.
* **Keluaran**: Rincian total potongan diskon, harga akhir yang harus dibayar, dan jumlah kupon yang diperoleh.

```go
package main

import "fmt"

func main() {
    var belanja int
    fmt.Print("Total Belanja: ")
    fmt.Scan(&belanja)

    kupon := belanja / 50000
    diskon := 0
    if belanja >= 100000 {
        diskon = (belanja * 10) / 100
    }
    totalBayar := belanja - diskon

    fmt.Printf("Diskon      : Rp %d\n", diskon)
    fmt.Printf("Total Bayar : Rp %d\n", totalBayar)
    fmt.Printf("Kupon Undian: %d lembar\n", kupon)
}
```

---

## 7.3 Tips Menghadapi Asesmen Praktikum

1. **Baca Masukan & Keluaran dengan Seksama**: Perhatikan tipe data (apakah integer atau pecahan) dan format spasi/baris baru.
2. **Uji Kasus Ekstrim (*Edge Cases*)**: Uji nilai terkecil (misal 0 atau 1), nilai batas, dan nilai besar.
3. **Cek Kompilasi Sebelum Submisi**: Pastikan kode dapat dikompilasi dengan `go run` atau `go build` tanpa kesalahan sintaks.
4. **Perhatikan Batas Waktu**: Alokasikan waktu untuk memahami soal, merancang kode, dan melakukan pengujian.

---

## Kesimpulan Modul 7

* Asesmen 1 menguji kematangan logika dasar dan kecepatan penulisan program Go.
* Pemahaman mendalam tentang pembagian bulat, sisa bagi, dan perulangan `for` merupakan kunci keberhasilan menyelesaikan soal asesmen.
