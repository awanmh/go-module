# MODUL 9: STRUKTUR KONTROL - IF-THEN (PERCABANGAN TUNGGAL)

Modul ini membahas **paradigma percabangan (selection/branching)** dalam algoritma, ekspresi logika boolean, operator relasional, dan implementasi struktur kontrol percabangan tunggal (*If-Then*) pada bahasa Go.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami konsep percabangan untuk pengambilan keputusan dalam program
* Menggunakan operator relasional (`==`, `!=`, `<`, `<=`, `>`, `>=`) dan operator logika (`&&`, `||`, `!`)
* Menyusun struktur percabangan tunggal `if` di bahasa Go
* Menerapkan logika filter kondisi untuk memecahkan permasalahan dunia nyata

---

## 9.1 Paradigma Percabangan

### Apa itu Percabangan?

Dalam kehidupan sehari-hari, kita sering mengambil keputusan berdasarkan kondisi tertentu:
> *"Jika hari ini hujan, maka bawa payung."*

Dalam pemrograman:
> **Percabangan (Branching)** adalah mekanisme yang memungkinkan program memilih untuk **mengeksekusi suatu blok instruksi atau melewatinya**, bergantung pada hasil evaluasi suatu kondisi boolean (*true* atau *false*).

### Diagram Alur If-Then:

```text
                    ┌──────────────────┐
                    │  Kondisi Diuji   │
                    └────────┬─────────┘
                             │
                  /─────────────────────\
                 <  Apakah Kondisi Benar >
                  \─────────────────────/
                     │               │
               Ya    │               │ Tidak
                     ▼               │
           ┌──────────────────┐      │
           │  Aksi Dieksekusi │      │
           └─────────┬────────┘      │
                     │               │
                     ▼               ▼
           ┌─────────────────────────────────┐
           │      Instruksi Selanjutnya      │
           └─────────────────────────────────┘
```

---

## 9.2 Operator Relasional dan Logika

### 1. Operator Relasional (Pembanding)
| Operator | Arti | Contoh | Hasil |
| :---: | :--- | :---: | :---: |
| `==` | Sama dengan | `5 == 5` | `true` |
| `!=` | Tidak sama dengan | `5 != 3` | `true` |
| `<` | Lebih kecil | `4 < 9` | `true` |
| `<=` | Lebih kecil atau sama dengan | `7 <= 7` | `true` |
| `>` | Lebih besar | `10 > 2` | `true` |
| `>=` | Lebih besar atau sama dengan | `3 >= 5` | `false` |

### 2. Operator Logika
* `&&` (AND): Bernilai `true` hanya jika **kedua** operan bernilai `true`.
* `||` (OR): Bernilai `true` jika **salah satu atau kedua** operan bernilai `true`.
* `!` (NOT): Membalik nilai kebenaran (`!true` menjadi `false`).

---

## 9.3 Sintaks If di Go

Di Go, penulisan kondisi `if` **tidak memerlukan tanda kurung biasa `()`**, tetapi kurung kurawal `{}` **wajib** digunakan:

```go
if kondisi {
    // Instruksi dijalankan jika kondisi bernilai true
}
```

---

## 9.4 Contoh Program Lengkap

### Program 1 – Nilai Mutlak (Absolute Value)
Menghitung nilai mutlak bilangan (jika bilangan negatif, kalikan dengan -1 agar menjadi positif):

```go
package main

import "fmt"

func main() {
    var x int

    fmt.Print("Masukkan sebuah bilangan bulat: ")
    fmt.Scan(&x)

    nilaiMutlak := x
    if nilaiMutlak < 0 {
        nilaiMutlak = -nilaiMutlak
    }

    fmt.Printf("|%d| = %d\n", x, nilaiMutlak)
}
```

---

### Program 2 – Penentuan Potongan Diskon
Seseorang berhak mendapat potongan harga Rp 15.000,- jika total belanjanya lebih dari Rp 100.000,-:

```go
package main

import "fmt"

func main() {
    var belanja int

    fmt.Print("Total Belanja (Rp): ")
    fmt.Scan(&belanja)

    potongan := 0
    if belanja > 100000 {
        potongan = 15000
    }

    totalAkhir := belanja - potongan
    fmt.Printf("Potongan   : Rp %d\n", potongan)
    fmt.Printf("Total Bayar: Rp %d\n", totalAkhir)
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Kapasitas Motor Peserta Touring
Sebuah program Go digunakan untuk menentukan jumlah motor yang diperlukan peserta touring.
Ketentuan:
* 1 motor muat maksimal 2 orang (1 pengemudi dan 1 boncengan).
* Setiap motor diprioritaskan terisi 2 orang.
* Jika ada 1 orang tersisa yang belum mendapat motor, wajib disediakan 1 motor tambahan untuknya.

*Contoh Masukan & Keluaran:*
| Masukan (Orang) | Keluaran (Motor) | Penjelasan |
| :---: | :---: | :--- |
| `10` | `5` | $10 / 2 = 5$ motor penuh |
| `1` | `1` | 1 motor untuk 1 orang |
| `25` | `13` | 12 motor @ 2 orang + 1 motor untuk sisa 1 orang |
| `9` | `5` | 4 motor @ 2 orang + 1 motor untuk sisa 1 orang |

---

### Soal 2 – Pengecekan Bilangan Genap Negatif
Buat program yang menerima sebuah bilangan bulat, lalu menentukan apakah bilangan tersebut merupakan **"genap negatif"** atau **"bukan"**.
* Suatu bilangan disebut genap negatif jika: nilainya $< 0$ dan habis dibagi 2 (`n % 2 == 0`).

*Contoh Masukan & Keluaran:*
| Masukan | Keluaran |
| :---: | :--- |
| `10` | `bukan` |
| `-4` | `genap negatif` |
| `0` | `bukan` |
| `-2` | `genap negatif` |

---

### Soal 3 – Faktor Bilangan
Buat program untuk menentukan apakah bilangan bulat $a$ merupakan faktor dari bilangan bulat $b$.
Suatu bilangan $a$ disebut faktor dari $b$ apabila $b$ habis dibagi oleh $a$ ($b \% a == 0$).
* **Masukan**: Dua buah bilangan bulat positif $a$ dan $b$.
* **Keluaran**: Teks boolean atau pesan `"true"` jika $a$ faktor dari $b$, atau `"false"` jika bukan.

---

## Kesimpulan Modul 9

* Percabangan tunggal `if` digunakan ketika suatu aksi hanya dieksekusi jika kondisi khusus terpenuhi.
* Penggunaan operator relasional dan logika yang tepat memungkinkan penanganan syarat matematis yang kompleks secara ringkas.
