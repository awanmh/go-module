# MODUL 13: STRUKTUR KONTROL - REPEAT-UNTIL

Modul ini membahas **struktur perulangan berbasis kondisi di akhir (*post-condition loop*)**, karakteristik eksekusi minimal 1 kali, dan implementasi idiomatis pola *Repeat-Until* di Go menggunakan `for { ... if kondisi { break } }`.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami karakteristik perulangan *post-condition* (evaluasi kondisi dilakukan di akhir blok)
* Membedakan perilaku perulangan `while` (kondisi di awal) vs `repeat-until` (kondisi di akhir)
* Mengimplementasikan struktur *Repeat-Until* idiomatis di bahasa Go dengan aman
* Menggunakan pernyataan `break` untuk keluar dari badan perulangan
* Menerapkan validasi input pengguna dan akumulasi donasi hingga target tercapai

---

## 13.1 Karakteristik Repeat-Until

Pada perulangan konvensional *Repeat-Until*:
> Instruksi di dalam badan perulangan **selalu dieksekusi minimal satu kali**, baru kemudian kondisinya dievaluasi di akhir.

```text
                  ┌────────────────────────┐
                  │ Eksekusi Badan Loop    │ ◄── Selalu jalan minimal 1x
                  └───────────┬────────────┘
                              ▼
                    /───────────────────\
                   <  Apakah Kondisi Stop? >
                    \───────────────────/
                      │               │
                Ya    │               │ Tidak (Ulangi lagi)
                      ▼               │
           ┌──────────────────────┐   │
           │  Keluar dari Loop    │   │
           └──────────────────────┘   │
                      ▲               │
                      └───────────────┘
```

---

## 13.2 Implementasi Idiomatis di Go

Bahasa Go mengimplementasikan perulangan dengan evaluasi di akhir menggunakan kombinasi perulangan tak berhingga `for { ... }` bersama dengan pernyataan percabangan penghenti `if kondisi_berhenti { break }`:

```go
for {
    // Instruksi badan perulangan (pasti dijalankan minimal 1 kali)

    if kondisi_terminasi {
        break // Berhenti dan keluar dari perulangan
    }
}
```

---

## 13.3 Contoh Program Lengkap

### Program 1 – Cetak Kata Sebanyak Target Pengguna
Program membaca sebuah kata dan mencetaknya berulang kali hingga mencapai kuota yang diinginkan:

```go
package main

import "fmt"

func main() {
    var kata string
    var target int

    fmt.Print("Masukkan kata dan jumlah pengulangan: ")
    fmt.Scan(&kata, &target)

    pencacah := 0
    for {
        fmt.Println(kata)
        pencacah++

        if pencacah >= target {
            break
        }
    }
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Menghitung Banyaknya Digit Suatu Bilangan
Buatlah program yang digunakan untuk menghitung banyaknya digit dari suatu bilangan bulat positif.
* **Masukan**: Sebuah bilangan bulat positif.
* **Keluaran**: Satu bilangan bulat yang menyatakan banyaknya digit.

#### Contoh Uji:
| No | Masukan | Keluaran |
| :---: | :--- | :---: |
| 1 | `5` | `1` |
| 2 | `234` | `3` |
| 3 | `78787` | `5` |
| 4 | `1894256` | `7` |

---

### Soal 2 – Pembulatan Bertahap Menuju Bilangan Bulat
Buatlah program yang membaca sebuah bilangan desimal (pecahan), kemudian melakukan penambahan bertahap sebesar `0.1` pada setiap iterasi hingga mencapai pembulatan bilangan bulat di atasnya:
* **Masukan**: Bilangan desimal pecahan positif (contoh: `2.7`).
* **Keluaran**: Barisan bilangan hasil penambahan tiap iterasi sampai mencapai nilai bulat terdekat di atasnya.

#### Contoh Uji:
```text
Masukan : 2.7
Keluaran:
2.8
2.9
3.0
```

---

### Soal 3 – Penggalangan Donasi Sosial
Sebuah lembaga amal mengumpulkan donasi untuk kegiatan bencana.
1. Program pertama kali membaca nilai **target donasi** (dalam Rupiah).
2. Program kemudian berulang kali meminta pengguna memasukkan nominal sumbangan para donatur.
3. Setiap sumbangan ditambahkan ke total akumulasi, dan sisa kekurangan target diinformasikan ke layar.
4. Program berhenti ketika total donasi telah **mencapai atau melampaui** target yang ditetapkan.

---

## Kesimpulan Modul 13

* *Repeat-Until* menjamin kode di dalam badan perulangan setidaknya dieksekusi satu kali sebelum kondisi berhenti diperiksa.
* Di Go, pola ini diwujudkan secara bersih menggunakan `for { ... if kondisi { break } }`.
