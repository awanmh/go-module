# MODUL 17: SKEMA PEMROSESAN SEKUENSIAL

### Tujuan Pembelajaran

Mahasiswa mampu:

* Memahami konsep pemrosesan data secara sekuensial
* Memahami pola pembacaan data tanpa marker
* Memahami pola pembacaan data dengan marker
* Menangani kasus data kosong
* Memahami konsep elemen pertama sebagai kasus khusus
* Mengimplementasikan berbagai skema pemrosesan data dalam bahasa Go

---

## 17.1 Pengantar Skema Pemrosesan Sekuensial

### Apa itu Pemrosesan Sekuensial?

Pemrosesan sekuensial adalah:

> Proses membaca dan mengolah data satu per satu sesuai urutan kedatangannya.

---

### Mengapa Penting?

Banyak permasalahan komputasi menggunakan pola ini, seperti:

* Menghitung rata-rata
* Mencari nilai maksimum
* Menghitung jumlah data
* Menghitung frekuensi
* Mengolah data sensor
* Mengolah data transaksi

---

### Konsep Dasar

```text
Data Masuk
     │
     ▼
┌───────────┐
│ Dibaca    │
└─────┬─────┘
      ▼
┌───────────┐
│ Diproses  │
└─────┬─────┘
      ▼
┌───────────┐
│ Data Baru │
└─────┬─────┘
      ▼
   Ulangi
```

---

### Karakteristik Pemrosesan Sekuensial

* Data diproses satu per satu
* Tidak harus disimpan seluruhnya
* Cocok untuk data besar
* Umumnya menggunakan perulangan

---

## 17.2 Pembacaan Data Tanpa Marker

### Konsep

Pada pola ini:

> Semua data yang diberikan merupakan data valid yang harus diproses.

Jumlah data diketahui sejak awal.

---

### Alur Pemrosesan

```text
Input jumlah data (n)
        │
        ▼
Baca data ke-1
        │
        ▼
Proses
        │
        ▼
Baca data ke-2
        │
        ▼
Proses
        │
        ▼
      ...
        │
        ▼
Baca data ke-n
```

---

### Pseudocode

```text
input(n)

i = 1

selama i <= n:

    input(dat)

    proses(dat)

    i++
```

---

### Versi Go

```go
fmt.Scan(&n)

i := 0

for i < n {

    fmt.Scan(&dat)

    // proses data

    i++
}
```

---

### Contoh Kasus

#### Mencari Nilai Maksimum

Misal terdapat:

```text
5
7 3 15 9 12
```

---

#### Langkah

```text
max = sangat kecil

7  → max = 7
3  → tetap
15 → max = 15
9  → tetap
12 → tetap
```

---

#### Program Go

```go
package main

import "fmt"

func main() {

    var n, dat, max, i int

    fmt.Scan(&n)

    max = -999999

    i = 1

    for i <= n {

        fmt.Scan(&dat)

        if dat > max {
            max = dat
        }

        i++
    }

    fmt.Println("Data terbesar:", max)
}
```

---

### Insight Penting

Karena jumlah data diketahui:

```text
Tidak diperlukan marker
```

---

## 17.3 Pembacaan Data Dengan Marker

### Konsep

Pada pola ini:

> Jumlah data tidak diketahui sebelumnya.

Digunakan sebuah nilai khusus sebagai penanda berhenti.

Nilai khusus ini disebut:

```text
MARKER
```

---

### Ilustrasi

Misalnya:

```text
7 3 15 9 12 -1
```

Marker:

```text
-1
```

Data valid:

```text
7 3 15 9 12
```

Data:

```text
-1
```

Tidak diproses.

---

### Diagram

```text
Baca data
     │
     ▼
Apakah marker?
     │
 ┌───┴───┐
 │       │
Ya      Tidak
 │       │
 ▼       ▼
Selesai Proses
           │
           ▼
      Baca lagi
```

---

### Pseudocode

```text
input(dat)

selama dat != MARKER:

    proses(dat)

    input(dat)
```

---

### Versi Go

```go
fmt.Scan(&dat)

for dat != MARKER {

    // proses data

    fmt.Scan(&dat)
}
```

---

### Contoh Program

#### Mencari Nilai Maksimum

```go
package main

import "fmt"

func main() {

    var dat, max int

    max = -999999

    fmt.Scan(&dat)

    for dat != -1 {

        if dat > max {
            max = dat
        }

        fmt.Scan(&dat)
    }

    fmt.Println("Data terbesar:", max)
}
```

---

### Insight Penting

Marker:

```text
Bukan data valid
```

Sehingga:

```text
Tidak boleh ikut diproses
```

---

## 17.4 Kasus Data Kosong

### Permasalahan

Bagaimana jika data pertama langsung marker?

Contoh:

```text
-1
```

Artinya:

```text
Tidak ada data
```

---

### Mengapa Berbahaya?

Program seperti:

```go
max := -999999
```

akan menghasilkan:

```text
Data terbesar = -999999
```

Padahal:

```text
Tidak ada data
```

---

### Solusi

Periksa data pertama terlebih dahulu.

---

### Diagram

```text
Baca data pertama
        │
        ▼
 Apakah marker?
        │
 ┌──────┴──────┐
 │             │
Ya            Tidak
 │             │
 ▼             ▼
Data       Lanjut proses
Kosong
```

---

### Pseudocode

```text
input(dat)

jika dat == MARKER

    tampilkan:
    "tidak ada data"

selain itu

    proses data
```

---

### Versi Go

```go
fmt.Scan(&dat)

if dat == MARKER {

    fmt.Println("tidak ada data")

} else {

    for dat != MARKER {

        // proses

        fmt.Scan(&dat)
    }
}
```

---

### Contoh Program

```go
package main

import "fmt"

func main() {

    var dat, max int

    fmt.Scan(&dat)

    if dat == -1 {

        fmt.Println("tidak ada data")

    } else {

        max = -999999

        for dat != -1 {

            if dat > max {
                max = dat
            }

            fmt.Scan(&dat)
        }

        fmt.Println("Data terbesar:", max)
    }
}
```

---

### Insight Penting

Kasus kosong harus selalu diperiksa apabila:

```text
Jumlah data tidak diketahui
dan menggunakan marker
```

---

## 17.5 Elemen Pertama Sebagai Kasus Khusus

### Permasalahan

Pada pencarian maksimum biasanya digunakan:

```go
max = -999999
```

Namun bagaimana jika data:

```text
-1000000
```

Muncul?

Maka hasil menjadi salah.

---

### Solusi

Gunakan data pertama sebagai nilai awal.

---

### Ide Dasar

```text
Data pertama
↓
menjadi max awal
```

Tidak perlu lagi:

```text
BILANGAN_KECIL
```

---

### Diagram

```text
Baca data pertama
       │
       ▼
 max = data pertama
       │
       ▼
 Baca data berikutnya
       │
       ▼
 Bandingkan
       │
       ▼
 Update max
```

---

### Pseudocode

```text
input(dat)

jika marker:

    data kosong

selain itu:

    max = dat

    input(dat)

    selama dat != marker:

        jika dat > max:

            max = dat

        input(dat)
```

---

### Versi Go

```go
fmt.Scan(&dat)

if dat == MARKER {

    fmt.Println("tidak ada data")

} else {

    max = dat

    fmt.Scan(&dat)

    for dat != MARKER {

        if dat > max {
            max = dat
        }

        fmt.Scan(&dat)
    }

    fmt.Println(max)
}
```

---

### Contoh

#### Input

```text
-5 -10 -3 -8 -1
```

Marker:

```text
-1
```

---

#### Proses

```text
max = -5

-10 → tetap
-3  → max = -3
-8  → tetap
```

---

### Hasil

```text
-3
```

---

### Kelebihan Pendekatan Ini

#### Tidak perlu angka ekstrem

```text
Tidak perlu:
-999999
-999999999
```

---

#### Lebih Aman

Bekerja untuk:

```text
Bilangan positif
Bilangan negatif
Bilangan campuran
```

---

#### Lebih Umum

Digunakan dalam:

* Maksimum
* Minimum
* Rata-rata
* Statistik
* Pengolahan data real-time

---

## Perbandingan Skema Pemrosesan Sekuensial

| Skema          | Jumlah Data Diketahui | Marker | Kasus Kosong         |
| -------------- | --------------------- | ------ | -------------------- |
| Tanpa Marker   | Ya                    | Tidak  | Tidak                |
| Dengan Marker  | Tidak                 | Ya     | Bisa terjadi         |
| Kasus Khusus   | Tidak                 | Ya     | Ditangani            |
| Elemen Pertama | Tidak                 | Ya     | Ditangani lebih aman |

---

## Kesalahan Umum

### 1. Marker Ikut Diproses

Salah:

```go
if dat > max {
    max = dat
}

fmt.Scan(&dat)
```

Marker bisa ikut dihitung.

---

### 2. Tidak Menangani Data Kosong

Input:

```text
-1
```

Program menghasilkan nilai sampah.

---

### 3. Menggunakan Bilangan Kecil Tetap

```go
max = -999
```

Belum tentu lebih kecil dari semua data.

---

### 4. Lupa Membaca Data Berikutnya

```go
for dat != MARKER {

    proses(dat)

}
```

Mengakibatkan:

```text
Infinite Loop
```

---

## Ringkasan Pola Pemrosesan Sekuensial

### Pola 1

Jumlah data diketahui

```go
for i < n
```

---

### Pola 2

Jumlah data tidak diketahui

```go
for dat != MARKER
```

---

### Pola 3

Kasus kosong

```go
if dat == MARKER
```

---

### Pola 4

Elemen pertama sebagai nilai awal

```go
max = dat
```

---

## Kesimpulan

* Pemrosesan sekuensial adalah teknik membaca dan mengolah data satu per satu.
* Terdapat empat pola utama:

  * Tanpa marker
  * Dengan marker
  * Kasus data kosong
  * Elemen pertama sebagai kasus khusus
* Marker digunakan ketika jumlah data tidak diketahui.
* Data kosong harus selalu ditangani pada skema dengan marker.
* Menggunakan elemen pertama sebagai nilai awal lebih aman dibanding menggunakan konstanta bilangan kecil atau besar.
* Skema ini menjadi dasar berbagai algoritma seperti:

  * pencarian maksimum/minimum
  * rata-rata
  * statistik data
  * pengolahan streaming data
  * analisis data real-time

---
