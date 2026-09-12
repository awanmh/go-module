# MODUL 3: I/O, TIPE DATA & VARIABEL (LATIHAN 1 - DIV, MOD & CASTING)

Modul ini membahas **operasi pembagian bilangan bulat (*integer division*)**, operasi sisa bagi (*modulo*), konsep *type casting* (konversi tipe data), serta pemanfaatannya dalam pemecahan digit dan konversi satuan waktu.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami perbedaan mendasar antara pembagian riil (`/` pada float) dan pembagian bulat (`/` pada integer)
* Menggunakan operator modulo (`%`) untuk memperoleh sisa pembagian bilangan bulat
* Melakukan konversi tipe data (*type casting*) secara eksplisit dan aman
* Menyelesaikan masalah dekomposisi angka (pemecahan digit ratusan, puluhan, satuan)
* Mengimplementasikan konversi satuan waktu dan pecahan mata uang

---

## 3.1 Integer Division (`/`) dan Modulo (`%`)

Dalam ilmu pemrograman, pembagian bilangan bulat memiliki sifat khusus:

### 1. Integer Division (`/`)
Ketika dua bilangan bertipe `int` dibagi, hasilnya adalah **bilangan bulat**. Bagian pecahan/desimal langsung **dibuang (*truncated*)**, bukan dibulatkan ke nilai terdekat.

```text
7 / 2 = 3   (bukan 3.5)
19 / 5 = 3  (bukan 3.8)
```

### 2. Operator Modulo (`%`)
Operator modulo menghasilkan **sisa dari hasil pembagian bulat**.

```text
7 % 2 = 1   (karena 7 = (3 * 2) + 1)
19 % 5 = 4  (karena 19 = (3 * 5) + 4)
```

### Visualisasi Integer Division & Modulo:

```text
                ┌──────────────┐
                │ 17 dibagi 5  │
                └──────┬───────┘
                       │
         ┌─────────────┴─────────────┐
         ▼                           ▼
   17 / 5 = 3                 17 % 5 = 2
(Hasil Bagi Bulat)           (Sisa Pembagian)
```

---

## 3.2 Casting (Konversi Tipe Data)

Bahasa Go sangat ketat terhadap tipe data (*strongly typed*). Kompiler **tidak mengizinkan** operasi matematika antar tipe data yang berbeda secara implisit.

```go
var a int = 10
var b float64 = 2.5
// var c float64 = a * b // ERROR! Mismatched types int and float64
var c float64 = float64(a) * b // BENAR: a diubah secara eksplisit menjadi float64
```

### Jebakan Umum Pemula dalam Pembagian:
```go
var a, b int = 7, 2
var hasil float64 = float64(a / b)   // HASIL: 3.0 (salah! a/b diproses integer dulu = 3)
var hasilBenar float64 = float64(a) / float64(b) // HASIL: 3.5 (benar!)
```

---

## 3.3 Studi Kasus & Contoh Program

### Contoh 1 – Konversi Detik ke Jam, Menit, dan Detik
Diberikan input total waktu dalam satuan detik, ubah ke format: **X jam, Y menit, Z detik**.

```go
package main

import "fmt"

func main() {
    var totalDetik int

    fmt.Print("Masukkan total detik: ")
    fmt.Scan(&totalDetik)

    jam := totalDetik / 3600
    sisaDetik := totalDetik % 3600

    menit := sisaDetik / 60
    detik := sisaDetik % 60

    fmt.Printf("%d detik = %d jam, %d menit, %d detik\n", totalDetik, jam, menit, detik)
}
```

#### Contoh Eksekusi:
```text
Masukan : 3665
Keluaran: 3665 detik = 1 jam, 1 menit, 5 detik
```

---

### Contoh 2 – Memisahkan Digit Bilangan 3 Angka
Diberikan sebuah bilangan bulat 3 digit (100 - 999), pisahkan menjadi ratusan, puluhan, dan satuan:

```go
package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan bilangan 3 digit: ")
    fmt.Scan(&n)

    ratusan := n / 100
    puluhan := (n % 100) / 10
    satuan := n % 10

    fmt.Printf("Ratusan: %d, Puluhan: %d, Satuan: %d\n", ratusan, puluhan, satuan)
}
```

#### Contoh Eksekusi:
```text
Masukan : 749
Keluaran: Ratusan: 7, Puluhan: 4, Satuan: 9
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Pembagian Pecahan Uang
Buatlah program yang menerima masukan sejumlah uang kembalian (dalam Rupiah kelipatan seratus), kemudian menentukan lembaran uang yang harus diberikan untuk pecahan:
* Rp 100.000
* Rp 50.000
* Rp 20.000
* Rp 10.000
* Rp 5.000
* Rp 2.000
* Rp 1.000

*Petunjuk*: Gunakan pembagian bulat `/` untuk mencari jumlah lembar uang dan modulo `%` untuk meneruskan sisa uang ke pecahan berikutnya.

---

### Soal 2 – Membalik Angka 3 Digit
Buat program yang membaca satu bilangan bulat 3 digit (contoh: `841`), kemudian menampilkan angka tersebut dalam posisi terbalik (`148`).
* **Masukan**: Sebuah bilangan bulat antara 100 s.d. 999.
* **Keluaran**: Angka setelah dibalik.

---

### Soal 3 – Perhitungan Rata-Rata Presisi
Buat program yang membaca 3 bilangan bulat positif $a, b, c$, lalu menghitung nilai rata-rata dari ketiga bilangan tersebut dengan menampilkan 2 angka di belakang koma. Pastikan menggunakan *type casting* yang benar agar pecahan tidak hilang!

---

## Kesimpulan Modul 3

* Operator `/` pada integer selalu memotong angka desimal.
* Operator `%` (modulo) sangat bermanfaat untuk ekstraksi digit, konversi siklus waktu, dan klasifikasi kelipatan.
* Konversi tipe data (*type casting*) wajib dilakukan sebelum operasi aritmatika dieksekusi jika menginginkan hasil presisi desimal.
