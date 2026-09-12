# MODUL 5: STRUKTUR KONTROL - FOR-LOOP (BAGIAN 1)

Modul ini membahas **paradigma perulangan (looping)**, struktur sintaks perulangan berbasis iterasi/pencacah (*counted loop*) menggunakan `for` di Go, perulangan naik (*ascending*), dan perulangan turun (*descending*).

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami konsep dasar dan urgensi struktur kontrol perulangan dalam algoritma
* Menyusun struktur `for` loop dengan 3 komponen utama (*init*, *condition*, *post*)
* Mengimplementasikan perulangan naik (*increment*) dan turun (*decrement*)
* Memahami alur eksekusi perulangan secara bertahap (*step-by-step tracing*)
* Menyelesaikan persoalan pencetakan deret dan barisan bilangan

---

## 5.1 Paradigma Perulangan

### Mengapa Perulangan Diperlukan?

Bayangkan Anda diminta mencetak teks `"Belajar Pemrograman"` sebanyak 1.000 kali.
Menuliskan `fmt.Println()` seribu kali akan sangat tidak efisien dan rentan kesalahan.

> **Perulangan (Loop)** memungkinkan satu blok perintah yang sama dijalankan berulang kali selama kondisi yang ditentukan masih terpenuhi (*true*).

---

## 5.2 Anatomi For-Loop di Go

Bahasa Go **hanya memiliki satu kata kunci perulangan**, yaitu `for`. Tidak ada kata kunci `while` atau `do-while` terpisah seperti di bahasa lain.

Struktur dasar *counted loop*:

```go
for inisialisasi; kondisi; post-statement {
    // Pernyataan / aksi yang diulang
}
```

```text
               ┌────────────────────────┐
               │    Inisialisasi        │ (Dijalankan sekali di awal, cth: i := 1)
               └───────────┬────────────┘
                           ▼
                 /───────────────────\
                < Apakah Kondisi Benar? >
                 \───────────────────/
                   │               │
             Ya    │               │ Tidak (Loop Selesai)
                   ▼               ▼
        ┌──────────────────┐    ┌─────────────────────┐
        │ Eksekusi Badan   │    │ Keluar dari Loop    │
        │ Perulangan       │    └─────────────────────┘
        └──────────┬───────┘
                   ▼
        ┌──────────────────┐
        │  Post Statement  │ (cth: i++)
        └──────────┬───────┘
                   │
                   └───────► (Kembali cek kondisi)
```

---

## 5.3 Pola Dasar Perulangan For

### 1. Perulangan Naik (*Ascending*)
Menghitung maju dari angka kecil ke angka besar:
```go
for i := 1; i <= 5; i++ {
    fmt.Print(i, " ")
}
// Output: 1 2 3 4 5
```

### 2. Perulangan Turun (*Descending*)
Menghitung mundur dari angka besar ke angka kecil:
```go
for i := 5; i >= 1; i-- {
    fmt.Print(i, " ")
}
// Output: 5 4 3 2 1
```

### 3. Perulangan dengan Step Tertentu
Pencacah bertambah dengan kelipatan tertentu:
```go
for i := 0; i <= 10; i += 2 {
    fmt.Print(i, " ")
}
// Output: 0 2 4 6 8 10
```

---

## 5.4 Contoh Program Lengkap

### Program 1 – Deret Angka 1 sampai N
Membaca sebuah bilangan bulat $N$, kemudian mencetak semua bilangan bulat dari 1 sampai $N$:

```go
package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan batas n: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}
```

#### Contoh Eksekusi:
```text
Masukan : 6
Keluaran: 1 2 3 4 5 6
```

---

### Program 2 – Penjumlahan Deret Bilangan ($1 + 2 + ... + N$)
Program menghitung jumlah total dari 1 hingga $N$:

```go
package main

import "fmt"

func main() {
    var n int
    var total int = 0

    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        total = total + i
    }

    fmt.Printf("Jumlah 1 s.d. %d adalah %d\n", n, total)
}
```

#### Trace Eksekusi ($N = 4$):
| Iterasi ($i$) | Nilai Awal `total` | Operasi (`total + i`) | Nilai Akhir `total` |
| :---: | :---: | :---: | :---: |
| 1 | 0 | 0 + 1 | 1 |
| 2 | 1 | 1 + 2 | 3 |
| 3 | 3 | 3 + 3 | 6 |
| 4 | 6 | 6 + 4 | 10 |

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Barisan Bilangan Kelipatan
Buatlah program yang menerima dua buah masukan bilangan bulat: bilangan pengali $k$ dan batas atas $n$.
Program kemudian menampilkan semua bilangan kelipatan $k$ yang berada pada rentang $1$ sampai $n$.
* **Masukan**: Dua bilangan bulat $k$ dan $n$ dipisahkan spasi.
* **Keluaran**: Barisan bilangan kelipatan $k$ dipisahkan spasi.

#### Contoh Uji:
```text
Masukan : 3 20
Keluaran: 3 6 9 12 15 18
```

---

### Soal 2 – Tabel Konversi Suhu Berjenjang
Buat program yang membaca batas awal suhu Celsius, batas akhir suhu Celsius, dan besar langkah kenaikan (*step*). Program kemudian mencetak tabel konversi ke Fahrenheit:
$$F = (C \times 9/5) + 32$$

#### Format Contoh:
```text
Masukan : 0 100 20
Keluaran:
Celsius | Fahrenheit
0.00    | 32.00
20.00   | 68.00
40.00   | 104.00
60.00   | 140.00
80.00   | 176.00
100.00  | 212.00
```

---

## Kesimpulan Modul 5

* For loop adalah struktur perulangan mendasar di Go yang menggabungkan inisialisasi, kondisi, dan post-statement dalam satu baris deklarasi yang rapi.
* Variabel pencacah (*iterator*) bertindak sebagai pengontrol berapa kali badan perulangan akan dieksekusi.
