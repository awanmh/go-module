# MODUL 4: I/O, TIPE DATA & VARIABEL (LATIHAN 2 - STUDI KASUS TERAPAN)

Modul ini berfokus pada **penguatan pemecahan masalah (problem solving)** menggunakan kombinasi I/O, tipe data, variabel, rumus matematika terapan, dan geometri bidang datar.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Mengimplementasikan rumus geometri dan fisika matematika ke dalam kode program Go
* Menggunakan fungsi-fungsi paket `math` (seperti `math.Sqrt`, `math.Pow`)
* Menangani kasus masukan majemuk dengan tipe data bilangan riil (*floating point*)
* Menghasilkan keluaran dengan format desimal presisi

---

## 4.1 Menggunakan Paket Matematika (`math`)

Untuk perhitungan matematika tingkat lanjut, Go menyediakan paket bawaan `math`:

```go
import "math"
```

Beberapa fungsi esensial:
* `math.Sqrt(x)` : Menghitung akar kuadrat $\sqrt{x}$ (parameter & hasil bertipe `float64`).
* `math.Pow(x, y)` : Menghitung pemangkatan $x^y$.
* `math.Pi` : Nilai konstanta $\pi$ (presisi tinggi ~3.141592653589793).

---

## 4.2 Studi Kasus 1 – Jarak Euclidean Antara Dua Titik

Diberikan koordinat dua titik pada bidang dua dimensi: Titik $A(x_1, y_1)$ dan Titik $B(x_2, y_2)$.
Jarak antara kedua titik dihitung menggunakan rumus Euclidean:
$$d = \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2}$$

### Program Lengkap:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var x1, y1, x2, y2 float64

    fmt.Println("Masukkan koordinat titik A (x1 y1):")
    fmt.Scan(&x1, &y1)

    fmt.Println("Masukkan koordinat titik B (x2 y2):")
    fmt.Scan(&x2, &y2)

    jarak := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

    fmt.Printf("Jarak titik (%.2f, %.2f) ke (%.2f, %.2f) = %.2f\n", x1, y1, x2, y2, jarak)
}
```

#### Contoh Eksekusi:
```text
Masukan : 
1.0 1.0
4.0 5.0

Perhitungan:
dx = 4.0 - 1.0 = 3.0
dy = 5.0 - 1.0 = 4.0
jarak = sqrt(3^2 + 4^2) = sqrt(9 + 16) = sqrt(25) = 5.00

Keluaran: Jarak titik (1.00, 1.00) ke (4.00, 5.00) = 5.00
```

---

## 4.3 Studi Kasus 2 – Perhitungan Luas dan Keliling Persegi Panjang

Sebuah taman kota berbentuk persegi panjang memiliki panjang $P$ dan lebar $L$. Hitunglah luas dan keliling taman tersebut:

```go
package main

import "fmt"

func main() {
    var panjang, lebar float64

    fmt.Print("Masukkan panjang dan lebar taman: ")
    fmt.Scan(&panjang, &lebar)

    luas := panjang * lebar
    keliling := 2 * (panjang + lebar)

    fmt.Printf("Luas Taman     = %.2f meter persegi\n", luas)
    fmt.Printf("Keliling Taman = %.2f meter\n", keliling)
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Perhitungan Hipotenusa Segitiga Siku-Siku
Diberikan panjang dua sisi tegak sebuah segitiga siku-siku, yaitu sisi alas $a$ dan sisi tinggi $b$.
Hitung panjang sisi miring (hipotenusa) $c$ menggunakan teorema Pythagoras:
$$c = \sqrt{a^2 + b^2}$$

#### Format Masukan & Keluaran:
* **Masukan**: Dua buah bilangan desimal $a$ dan $b$.
* **Keluaran**: Satu bilangan desimal format 2 digit di belakang koma menyatakan panjang sisi miring $c$.

---

### Soal 2 – Perhitungan Tagihan Listrik Sederhana
Sebuah rumah tangga menggunakan daya listrik dengan pemakaian $K$ kiloWatt-hour (kWh).
Tarif dasar per kWh adalah Rp 1.500,- dengan biaya beban tetap bulanan sebesar Rp 25.000,-.
Selain itu, dikenakan pajak pertambahan nilai (PPN) sebesar 10% dari total tagihan (pemakaian + beban tetap).
* **Masukan**: Bilangan riil menyatakan total pemakaian kWh.
* **Keluaran**: Rincian biaya pemakaian, biaya beban, pajak, dan total tagihan bersih yang harus dibayar.

---

## Kesimpulan Modul 4

* Pemecahan masalah pemrograman membutuhkan pemodelan matematis yang runtut sebelum dituliskan ke dalam kode Go.
* Pustaka bawaan `math` menyediakan fungsi-fungsi trigonometri, akar, dan eksponensial yang siap pakai untuk berbagai perhitungan saintifik.
