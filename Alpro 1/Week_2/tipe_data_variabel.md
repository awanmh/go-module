# MODUL 2: I/O, TIPE DATA & VARIABEL

Modul ini membahas **konsep dasar input-output (I/O)**, deklarasi variabel, tipe data primitif (integer, float, boolean, string, karakter/ASCII), konstanta simbolik, serta operasi aritmatika dasar dalam bahasa Go.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami mekanisme pembacaan input (`fmt.Scan`, `fmt.Scanln`, `fmt.Scanf`) dan pencetakan output (`fmt.Print`, `fmt.Println`, `fmt.Printf`)
* Membedakan jenis tipe data primitif pada Go (`int`, `float64`, `bool`, `string`, `byte`)
* Mendeklarasikan variabel menggunakan kata kunci `var` dan operator deklarasi ringkas `:=`
* Mendefinisikan konstanta menggunakan kata kunci `const`
* Memahami representasi karakter berbasis kode ASCII
* Menyelesaikan masalah komputasi aritmatika sederhana

---

## 2.1 Konsep Variabel & Memori

### Apa itu Variabel?

Variabel adalah:

> Tempat penyimpanan nilai sementara di dalam memori komputer (RAM) yang memiliki **nama**, **tipe data**, dan **nilai**.

### Analogi Wadah Bertanda

```text
┌──────────────┐
│  nama : "Budi"│ ◄── Tipe data: string (teks)
└──────────────┘
┌──────────────┐
│  umur : 19   │ ◄── Tipe data: int (bilangan bulat)
└──────────────┘
┌──────────────┐
│  ipk  : 3.85 │ ◄── Tipe data: float64 (bilangan pecahan)
└──────────────┘
```

---

## 2.2 Deklarasi Variabel di Go

Go menyediakan beberapa cara deklarasi variabel:

### 1. Deklarasi Eksplisit dengan Kata Kunci `var`
```go
var nilai int
nilai = 90

var nama string = "Anto"
var berat, tinggi float64 = 65.5, 172.0
```

### 2. Deklarasi Ringkas (*Short Variable Declaration*) `:=`
Digunakan di dalam fungsi untuk deklarasi sekaligus inisialisasi awal nilai (tipe data ditentukan secara otomatis oleh kompiler / *type inference*):
```go
umur := 20          // int
gaji := 4500000.50  // float64
aktif := true       // bool
pesan := "Halo"     // string
```

---

## 2.3 Tipe Data Primitif

| Tipe Data | Deskripsi | Rentang / Contoh |
| :--- | :--- | :--- |
| `int` | Bilangan bulat bertanda | $-2^{63}$ s.d. $2^{63}-1$ (arsitektur 64-bit) |
| `float64` | Bilangan desimal/pecahan | Presisi 64-bit IEEE-754 (contoh: `3.14`, `-0.05`) |
| `bool` | Logika kebenaran | `true` atau `false` |
| `string` | Kumpulan karakter teks diapit tanda petik ganda | `"Laboratorium Informatika"` |
| `byte` / `rune` | Karakter tunggal (representasi kode ASCII / Unicode) | `'A'`, `'7'`, `'#'` |

---

## 2.4 Konstanta Simbolik (`const`)

Konstanta adalah identifier yang nilainya **tetap** dan tidak dapat diubah setelah didefinisikan:

```go
const PI float64 = 3.1415926535
const GRAVITASI = 9.8
```

---

## 2.5 Input dan Output (I/O)

### Operasi Output
* `fmt.Print()`: Mencetak data tanpa baris baru.
* `fmt.Println()`: Mencetak data dan menambahkan baris baru (*newline*) di akhir.
* `fmt.Printf()`: Mencetak data dengan format tertentu menggunakan *format specifier*:
  * `%d` : Bilangan bulat (*integer*)
  * `%f` / `%.2f` : Bilangan pecahan (*float* dengan format 2 desimal)
  * `%s` : Teks (*string*)
  * `%c` : Karakter tunggal (*ASCII character*)
  * `%t` : Nilai boolean
  * `%v` : Nilai umum (*default representation*)

### Operasi Input
* `fmt.Scan(&var)`: Membaca masukan yang dipisahkan oleh spasi.
* `fmt.Scanln(&var)`: Membaca masukan hingga tombol Enter ditekan.
* `fmt.Scanf("%c", &var)`: Membaca masukan sesuai format tertentu (sangat berguna untuk membaca karakter tunggal).

> [!NOTE]
> Simbol tanda dan (`&`) di depan nama variabel saat pembacaan input menunjukkan alamat memori (*address-of*) tempat nilai masukan akan disimpan.

---

## 2.6 Contoh Program Lengkap

### Program 1 – Penjumlahan 5 Bilangan Bulat
Program membaca lima bilangan bulat dan mencetak hasil penjumlahannya:

```go
package main

import "fmt"

func main() {
    var a, b, c, d, e int
    var hasil int

    fmt.Println("Masukkan 5 bilangan bulat dipisahkan spasi:")
    fmt.Scan(&a, &b, &c, &d, &e)

    hasil = a + b + c + d + e

    fmt.Printf("Hasil penjumlahan %d + %d + %d + %d + %d = %d\n", a, b, c, d, e, hasil)
}
```

#### Contoh Eksekusi:
```text
Masukan : 3 2 7 10 2
Keluaran: Hasil penjumlahan 3 + 2 + 7 + 10 + 2 = 24
```

---

### Program 2 – Evaluasi Persamaan Matematika
Sebuah program menghitung nilai fungsi:
$$f(x) = \frac{2}{x + 5} + 5$$

```go
package main

import "fmt"

func main() {
    var x float64

    fmt.Print("Masukkan nilai x: ")
    fmt.Scan(&x)

    fx := (2.0 / (x + 5.0)) + 5.0

    fmt.Printf("f(%.2f) = %.6f\n", x, fx)
}
```

#### Contoh Eksekusi:
```text
Masukan : 5
Keluaran: f(5.00) = 5.200000
```

---

### Program 3 – Karakter dan Kode ASCII
Program membaca 5 nilai integer (kode ASCII) lalu mencetaknya sebagai huruf, kemudian membaca 3 karakter huruf dan mencetak karakter setelahnya:

```go
package main

import "fmt"

func main() {
    var c1, c2, c3, c4, c5 int
    var h1, h2, h3 byte

    // Membaca 5 kode ASCII
    fmt.Scan(&c1, &c2, &c3, &c4, &c5)

    // Membaca string 3 karakter
    var s string
    fmt.Scan(&s)
    h1 = s[0]
    h2 = s[1]
    h3 = s[2]

    // Menampilkan representasi karakter
    fmt.Printf("%c%c%c%c%c\n", c1, c2, c3, c4, c5)
    // Menampilkan 3 karakter setelahnya (geser +1)
    fmt.Printf("%c%c%c\n", h1+1, h2+1, h3+1)
}
```

#### Contoh Eksekusi:
```text
Masukan:
66 97 103 117 115
SNO

Keluaran:
Bagus
TOP
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Analisis Program Rotasi Nilai
Perhatikan instruksi pemindahan nilai berikut:
```go
temp := satu
satu = dua
dua = tiga
tiga = temp
```
**Tugas**:
Buatlah program Go lengkap untuk menguji cuplikan kode di atas. Masukkan tiga buah teks kata atau bilangan, lakukan pertukaran, dan terangkan alur pemindahan data yang terjadi!

---

### Soal 2 – Tahun Kabisat
Tahun kabisat adalah tahun yang habis dibagi 400 **atau** (habis dibagi 4 dan tidak habis dibagi 100).
* **Masukan**: Sebuah bilangan bulat positif menyatakan tahun (misal: 2016, 2018).
* **Keluaran**: Nilai boolean `true` jika tahun kabisat, atau `false` jika bukan kabisat.

#### Contoh Uji:
| Masukan | Keluaran |
| :--- | :--- |
| `2016` | `true` |
| `2018` | `false` |

---

### Soal 3 – Perhitungan Geometri Bola
Buatlah program untuk menghitung **Volume** dan **Luas Permukaan** bola berdasarkan input jari-jari $r$:
* $\text{Volume} = \frac{4}{3} \pi r^3$
* $\text{Luas} = 4 \pi r^2$
* Gunakan konstanta `PI = 3.1415926535` dan tipe data `float64`.

---

### Soal 4 – Konversi Suhu
Buat program yang membaca suhu dalam satuan Celsius ($C$), kemudian menghitung dan menampilkan konversinya ke:
* Fahrenheit ($F$): $(C \times 9/5) + 32$
* Reamur ($R$): $C \times 4/5$
* Kelvin ($K$): $C + 273.15$

---

## Kesimpulan Modul 2

* Variabel di Go memiliki tipe data yang tegas (*statically typed*).
* Pemilihan tipe data yang tepat (`int` untuk pencacah/jumlah, `float64` untuk perhitungan presisi, `byte` untuk karakter) krusial untuk akurasi komputasi.
* Karakter pada dasarnya adalah nilai integer yang dipetakan menurut tabel kode ASCII.
