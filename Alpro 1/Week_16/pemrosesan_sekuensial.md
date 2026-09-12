# MODUL 16: SKEMA PEMROSESAN SEKUENSIAL (PENGAYAAN)

Modul ini membahas **skema pemrosesan data secara sekuensial (berurutan)**, pola pembacaan data tanpa penanda (*without marker*), pembacaan dengan penanda akhir (*sentinel/marker*), penanganan deret data kosong, serta teknik pemrosesan khusus pada elemen pertama.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami konsep pemrosesan sekuensial dalam komputasi aliran data (*stream processing*)
* Mengimplementasikan pembacaan data tanpa marker (jumlah data $N$ telah diketahui)
* Mengimplementasikan pembacaan data dengan marker penanda akhir deret
* Menangani kasus khusus saat rangkaian data kosong
* Memahami teknik inisialisasi menggunakan elemen pertama data masukan

---

## 16.1 Pengantar Skema Pemrosesan Sekuensial

Pemrosesan sekuensial adalah:

> Metode pengolahan sekumpulan data di mana data dibaca dan diproses **satu per satu menurut urutan kedatangannya**, tanpa keharusan menyimpan seluruh data ke dalam array terlebih dahulu.

### Mengapa Sangat Efisien?
* **Hemat Memori**: Program tidak perlu mengalokasikan array besar di memori.
* **Skalabilitas**: Mampu mengolah aliran data berukuran sangat masif (jutaan data transaksi atau data sensor).

---

## 16.2 Pembacaan Data Tanpa Marker

Jumlah data diketahui sejak awal ($N$ buah data).

```text
Input N
  │
  ▼
Perulangan i = 1 s.d. N:
  ├── Baca data ke-i
  └── Proses data (misal: akumulasi)
```

---

## 16.3 Pembacaan Data dengan Marker (Sentinel)

Jumlah data tidak diketahui, tetapi diakhiri oleh satu nilai khusus yang disebut **marker** (penanda berhenti yang nilainya bukan data valid).

### Contoh Program: Menghitung Rata-rata Bilangan dengan Marker 9999
Program terus membaca bilangan riil dan berhenti ketika menemukan nilai `9999`:

```go
package main

import "fmt"

func main() {
    var bilangan float64
    var total float64 = 0.0
    var jumlahData int = 0

    fmt.Println("Masukkan bilangan (akhiri dengan 9999):")
    fmt.Scan(&bilangan)

    for bilangan != 9999 {
        total = total + bilangan
        jumlahData++
        fmt.Scan(&bilangan) // Baca data berikutnya
    }

    if jumlahData > 0 {
        rerata := total / float64(jumlahData)
        fmt.Printf("Banyak data: %d | Rata-rata: %.2f\n", jumlahData, rerata)
    } else {
        fmt.Println("Rangkaian data kosong (hanya ada marker)")
    }
}
```

---

## 16.4 Elemen Pertama Sebagai Kasus Khusus

Teknik ini sangat penting saat mencari nilai **Maksimum** atau **Minimum**:
> Nilai awal `max` atau `min` **tidak boleh diasumsikan 0**, melainkan harus diinisialisasi dari **data pertama yang dibaca**.

```go
package main

import "fmt"

func main() {
    var n, nilai int

    fmt.Print("Banyak data: ")
    fmt.Scan(&n)

    if n > 0 {
        // Baca data pertama sebagai kasus khusus
        fmt.Scan(&nilai)
        terbesar := nilai
        terkecil := nilai

        // Baca sisa data ke-2 sampai ke-n
        for i := 2; i <= n; i++ {
            fmt.Scan(&nilai)
            if nilai > terbesar {
                terbesar = nilai
            }
            if nilai < terkecil {
                terkecil = nilai
            }
        }

        fmt.Printf("Maksimum: %d | Minimum: %d\n", terbesar, terkecil)
    }
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Rata-rata dengan Marker 9999
Diberikan sejumlah bilangan riil yang diakhiri dengan marker `9999`. Buatlah program untuk menghitung nilai rata-rata dari bilangan-bilangan tersebut (marker `9999` tidak ikut dihitung).

---

### Soal 2 – Pencarian String $x$ dalam Sekumpulan $N$ String
Diberikan sebuah string $x$, bilangan bulat $n$, dan $n$ buah data string berikutnya. Buat program untuk menjawab:
1. Apakah string $x$ ada di dalam kumpulan data?
2. Pada posisi urutan ke berapa string $x$ pertama kali ditemukan?
3. Ada berapa kali string $x$ muncul dalam kumpulan data?

---

### Soal 3 – Pengukuran Curah Hujan 4 Daerah
Empat daerah $A, B, C,$ dan $D$ yang bersebelahan mengukur curah hujan. Tetesan air hujan jatuh secara acak pada koordinat bidang $2D$ dari $(0,0)$ sampai $(1,1)$.
* Pembagian Daerah:
  * Daerah A: $x < 0.5$ dan $y \ge 0.5$
  * Daerah B: $x \ge 0.5$ dan $y \ge 0.5$
  * Daerah C: $x < 0.5$ dan $y < 0.5$
  * Daerah D: $x \ge 0.5$ dan $y < 0.5$
* Setiap tetesan mewakili $0.0001$ ml curah hujan.
* Program membaca jumlah tetesan $N$, diikuti koordinat $(x, y)$ untuk setiap tetesan, lalu menghitung total curah hujan yang diterima masing-masing daerah.

---

## Kesimpulan Modul 16

* Skema pemrosesan sekuensial adalah fondasi pemrosesan data efisien sebelum mempelajari struktur data array pada mata kuliah Alpro 2.
* Penanganan kasus khusus seperti deret data kosong dan inisialisasi elemen pertama menjamin ketahanan program (*robustness*) terhadap berbagai variasi masukan.
