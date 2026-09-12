# MODUL 6: STRUKTUR KONTROL - FOR-LOOP (BAGIAN 2: AKUMULATOR & POLA)

Modul ini membahas **pemanfaatan lanjut perulangan `for`**, teknik variabel akumulator (penjumlahan dan perkalian beruntun), penghitungan nilai faktorial, pemangkatan manual, serta pemrosesan $N$ data masukan berulang.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami dan mengimplementasikan konsep **akumulator** (penampung akumulatif)
* Menghitung nilai faktorial ($n!$) dan pemangkatan ($a^b$) menggunakan struktur iterasi
* Memproses sekumpulan masukan majemuk sebanyak $N$ kali
* Menganalisis kompleksitas iterasi dan mencegah *off-by-one errors*
* Memecahkan persoalan komputasi geometri dan barisan bilangan berulang

---

## 6.1 Konsep Akumulator

### Apa itu Akumulator?

Akumulator adalah:

> Variabel yang digunakan untuk **mengumpulkan atau menghimpun hasil perhitungan** dari setiap iterasi perulangan secara bertahap.

```text
Iterasi 1: akumulator = nilai_awal + data_1
Iterasi 2: akumulator = akumulator + data_2
Iterasi 3: akumulator = akumulator + data_3
...
Iterasi N: akumulator menyimpan total keseluruhan
```

### Aturan Inisialisasi Nilai Awal Akumulator:
* **Penjumlahan / Pencacah**: Inisialisasi awal dengan nilai netral penjumlahan, yaitu `0`.
  ```go
  var total int = 0
  ```
* **Perkalian / Faktorial / Pangkat**: Inisialisasi awal dengan nilai netral perkalian, yaitu `1`.
  ```go
  var hasilKali int = 1
  ```

---

## 6.2 Studi Kasus 1 – Perhitungan Faktorial ($n!$)

Definisi faktorial dari bilangan bulat non-negatif $n$:
$$n! = n \times (n - 1) \times (n - 2) \times ... \times 2 \times 1$$
Dengan kasus khusus: $0! = 1$.

```go
package main

import "fmt"

func main() {
    var n int
    var faktorial int = 1

    fmt.Print("Masukkan bilangan bulat non-negatif n: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        faktorial = faktorial * i
    }

    fmt.Printf("%d! = %d\n", n, faktorial)
}
```

#### Trace Perhitungan ($n = 5$):
| Nilai $i$ | `faktorial` Sebelum | Operasi (`faktorial * i`) | `faktorial` Sesudah |
| :---: | :---: | :---: | :---: |
| 1 | 1 | 1 * 1 | 1 |
| 2 | 1 | 1 * 2 | 2 |
| 3 | 2 | 2 * 3 | 6 |
| 4 | 6 | 6 * 4 | 24 |
| 5 | 24 | 24 * 5 | 120 |

---

## 6.3 Studi Kasus 2 – Pemrosesan Data Majemuk Sebanyak $N$ Kali

Pola umum membaca $N$ buah data di mana pengguna memasukkan jumlah data terlebih dahulu:

```text
Input N
  │
  ▼
Perulangan i dari 1 s.d. N:
  ├── Baca data ke-i
  └── Proses data ke-i (hitung, jumlahkan, dsb)
```

### Contoh Program: Menghitung Rata-rata $N$ Nilai
```go
package main

import "fmt"

func main() {
    var n int
    var nilai, total float64

    fmt.Print("Masukkan banyak data: ")
    fmt.Scan(&n)

    total = 0.0
    for i := 1; i <= n; i++ {
        fmt.Printf("Masukkan data ke-%d: ", i)
        fmt.Scan(&nilai)
        total = total + nilai
    }

    rerata := total / float64(n)
    fmt.Printf("Total = %.2f, Rata-rata = %.2f\n", total, rerata)
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Penjumlahan Sekumpulan Bilangan
Buatlah program untuk menjumlahkan sekumpulan bilangan bulat dari 1 sampai dengan $n$.
* **Masukan**: Bilangan bulat positif $n$.
* **Keluaran**: Hasil penjumlahan $1 + 2 + ... + n$.

#### Contoh Uji:
| No | Masukan | Keluaran |
| :--- | :--- | :--- |
| 1 | `3` | `6` |
| 2 | `1` | `1` |
| 3 | `5` | `15` |

---

### Soal 2 – Volume Sejumlah $N$ Kerucut
Buatlah program yang digunakan untuk menghitung volume dari $n$ buah kerucut jika diketahui jari-jari alas ($r$) dan tinggi ($t$) masing-masing kerucut.
Rumus volume kerucut:
$$V = \frac{1}{3} \pi r^2 t$$
*Gunakan:* `math.Pi`.
* **Masukan**: Baris pertama adalah bilangan bulat $n$ (jumlah kerucut). $n$ baris berikutnya masing-masing berisi dua bilangan desimal $r$ dan $t$.
* **Keluaran**: $n$ baris yang menyatakan volume masing-masing kerucut.

#### Contoh Uji:
```text
Masukan:
3
1 1
2 2
3 3

Keluaran:
1.0471975511965976
8.377580409572781
28.274333882308138
```

---

### Soal 3 – Pemangkatan Manual ($a^b$)
Buatlah program untuk menghitung hasil perpangkatan $a^b$ menggunakan operator perkalian berulang dan struktur perulangan `for` (tanpa menggunakan fungsi `math.Pow`).
* **Masukan**: Dua buah bilangan bulat positif $a$ dan $b$.
* **Keluaran**: Satu bilangan bulat hasil dari $a^b$.

#### Contoh Uji:
| Masukan | Keluaran |
| :--- | :--- |
| `2 3` | `8` |
| `5 4` | `625` |

---

## Kesimpulan Modul 6

* Akumulator adalah teknik fundamental dalam pemrograman untuk mereduksi sekumpulan data menjadi satu nilai ringkasan (jumlah, perkalian, rata-rata).
* Pola pembacaan $N$ data masukan memanfaatkan variabel pencacah $n$ sebagai batas atas perulangan `for`.
