# MODUL 12: STRUKTUR KONTROL - WHILE-LOOP

Modul ini membahas **struktur perulangan berbasis kondisi di awal (*pre-condition loop*)**, sintaks perulangan `while` idiomatis pada Go (`for condition {}`), penanganan kondisi terminasi, dan pencegahan *infinite loop*.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami perbedaan antara *counted loop* (`for i`) dan *conditional loop* (`while`)
* Mengimplementasikan perulangan kondisi di awal menggunakan sintaks `for kondisi {}` di Go
* Menentukan kondisi terminasi yang tepat agar program tidak mengalami *loop tak berhingga* (*infinite loop*)
* Memecahkan masalah yang jumlah iterasinya tidak dapat diprediksi sejak awal (misal: otentikasi login, dekomposisi digit angka)

---

## 12.1 Paradigma While-Loop

### Kapan Menggunakan While-Loop?

Pada modul sebelumnya (*for-loop*), kita selalu mengetahui batas pasti perulangan (misal: 10 kali, atau dari 1 s.d. $N$).
Namun, dalam banyak kasus nyata, **jumlah pengulangan tidak diketahui di awal**:
> *"Ulangi meminta kata sandi SELAMA kata sandi yang dimasukkan masih salah."*
> *"Ulangi membagi angka dengan 10 SELAMA angka tersebut masih lebih besar dari 0."*

### Diagram Alur While-Loop:

```text
               ┌────────────────────────┐
               │    Inisialisasi Nilai   │ (cth: input username & password)
               └───────────┬────────────┘
                           │
                           ▼
                 /───────────────────\
                < Apakah Kondisi Benar? >
                 \───────────────────/
                   │               │
             Ya    │               │ Tidak (Selesai)
                   ▼               ▼
        ┌──────────────────┐    ┌─────────────────────┐
        │ Eksekusi Badan   │    │ Keluar dari Loop    │
        │ Perulangan       │    └─────────────────────┘
        └──────────┬───────┘
                   ▼
        ┌──────────────────┐
        │ Perbarui Nilai   │ (Ambil input baru / ubah variabel)
        └──────────┬───────┘
                   │
                   └───────► (Kembali periksa kondisi)
```

> [!IMPORTANT]
> Karena kondisi diuji di **awal**, ada kemungkinan badan perulangan **tidak pernah dijalankan sama sekali (0 kali)** jika pada pemeriksaan pertama kondisi sudah bernilai `false`.

---

## 12.2 Sintaks While di Bahasa Go

Bahasa Go tidak memiliki kata kunci `while` terpisah. Perulangan `while` dituliskan menggunakan kata kunci `for` yang hanya diikuti oleh satu ekspresi kondisi:

```go
for kondisi {
    // Aksi yang diulang
    // WAJIB ada instruksi yang mengubah nilai variabel kondisi!
}
```

---

## 12.3 Contoh Program Lengkap

### Program 1 – Deret Perkalian Faktorial Menurun
Menampilkan deret perkalian faktorial dari suatu bilangan $N$ hingga mencapai 1:

```go
package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan bilangan n: ")
    fmt.Scan(&n)

    if n <= 1 {
        fmt.Println("1")
    } else {
        fmt.Print(n)
        n = n - 1
        for n >= 1 {
            fmt.Printf(" x %d", n)
            n = n - 1
        }
        fmt.Println()
    }
}
```

#### Contoh Eksekusi:
```text
Masukan : 5
Keluaran: 5 x 4 x 3 x 2 x 1
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Percobaan Gagal Login Pengguna
Buatlah program Go untuk menghitung berapa kali pengguna gagal melakukan login karena kesalahan kombinasi `username` dan `password`.
* **Ketentuan Akun Benar**: Username `"Admin"` dan Password `"Admin"` (tanpa tanda petik, peka huruf besar-kecil/*case sensitive*).
* **Alur Program**:
  1. Program membaca pasangan username dan password.
  2. Selama kredensial yang dimasukkan salah, program menghitung 1 kali kegagalan dan meminta input ulang.
  3. Ketika kredensial sudah benar, program berhenti dan menampilkan total kegagalan yang terjadi.

#### Contoh Sesi Uji:
```text
Masukan:
User123 user123
User admin
Admin admin
Admin Admin123
Admin Admin

Keluaran:
4 percobaan gagal login
```

---

### Soal 2 – Mencacah Digit Angka dari Kanan ke Kiri
Buatlah program yang membaca sebuah bilangan bulat positif sembarang $N$, kemudian mencacah dan menampilkan setiap digitnya satu per satu dimulai dari **digit terakhir (paling kanan) sampai digit pertama (paling kiri)**:
* *Petunjuk Logika*:
  * Digit paling kanan didapat dari: `n % 10`
  * Buang digit terakhir dengan: `n = n / 10`
  * Ulangi proses ini selama: `n > 0`

#### Contoh Uji:
```text
Masukan : 4829
Keluaran:
9
2
8
4
```

---

## Kesimpulan Modul 12

* While-loop di Go diwujudkan dengan sintaks `for kondisi {}`.
* Sangat cocok untuk perulangan berbasis kejadian (*event-driven*), validasi masukan berulang, dan pemrosesan algoritma numerik yang konvergen.
