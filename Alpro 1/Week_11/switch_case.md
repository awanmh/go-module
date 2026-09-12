# MODUL 11: STRUKTUR KONTROL - SWITCH-CASE

Modul ini membahas **struktur percabangan `switch-case`**, pemilihan multi-kondisi berbasis nilai diskrit, perbandingannya dengan `if-else`, pemanfaatan klausa `default`, multiple values per case, dan kata kunci `fallthrough` di Go.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami konsep dasar dan kegunaan struktur percabangan `switch-case`
* Mengimplementasikan `switch-case` dengan pencocokan nilai ekspresi maupun kondisi logika
* Menggunakan klausa `default` sebagai penanganan kasus cadangan (*fallback*)
* Memahami mekanisme `fallthrough` pada bahasa Go
* Memilih antara `switch-case` dan `if-else` sesuai karakteristik persoalan

---

## 11.1 Karakteristik Switch-Case di Go

`switch-case` adalah bentuk percabangan alternatif yang sering kali **lebih bersih dan lebih mudah dibaca** dibandingkan rantai `if-else if` yang panjang ketika menguji satu variabel dengan banyak nilai kemungkinan.

### Keistimewaan Switch di Go:
1. **Tidak Memerlukan `break`**: Di bahasa C/Java, setiap case wajib diakhiri `break`. Di Go, program **secara otomatis berhenti** setelah mengeksekusi case yang cocok.
2. **Multiple Values**: Satu case dapat menampung beberapa nilai sekaligus yang dipisahkan koma (`case 1, 3, 5:`).
3. **Keyword `fallthrough`**: Jika ingin eksekusi tetap berlanjut ke case berikutnya tanpa henti, Go menyediakan keyword eksplisit `fallthrough`.

```text
                     ┌──────────────────┐
                     │ Evaluasi Variabel│
                     └────────┬─────────┘
                              │
               ┌──────────────┼──────────────┐
       Nilai 1 │      Nilai 2 │      Lainnya │
               ▼              ▼              ▼
         ┌───────────┐  ┌───────────┐  ┌───────────┐
         │ Blok C1   │  │ Blok C2   │  │  Default  │
         └─────┬─────┘  └─────┬─────┘  └─────┬─────┘
               │              │              │
               └──────────────┼──────────────┘
                              ▼
                 ┌─────────────────────────┐
                 │  Lanjut ke Baris Akhir  │
                 └─────────────────────────┘
```

---

## 11.2 Sintaks Switch di Go

```go
switch variabel {
case nilai_1:
    // Aksi 1
case nilai_2, nilai_3:
    // Aksi 2
default:
    // Aksi jika tidak ada yang cocok
}
```

---

## 11.3 Contoh Program Lengkap

### Program 1 – Konversi Format Waktu 24 Jam ke 12 Jam (AM / PM)
Membaca jam dalam format 24 jam (0 s.d. 23) lalu mengonversi ke format 12 jam:

```go
package main

import "fmt"

func main() {
    var jam int

    fmt.Print("Masukkan jam (0-23): ")
    fmt.Scan(&jam)

    switch jam {
    case 0:
        fmt.Println("12 AM")
    case 12:
        fmt.Println("12 PM")
    default:
        if jam > 0 && jam < 12 {
            fmt.Printf("%d AM\n", jam)
        } else if jam > 12 && jam < 24 {
            fmt.Printf("%d PM\n", jam - 12)
        } else {
            fmt.Println("Jam tidak valid")
        }
    }
}
```

#### Contoh Eksekusi:
```text
Masukan : 13  -> Keluaran: 1 PM
Masukan : 0   -> Keluaran: 12 AM
Masukan : 12  -> Keluaran: 12 PM
```

---

### Program 2 – Menu Pilihan Operasi Aritmatika
Program kalkulator sederhana menggunakan `switch-case`:

```go
package main

import "fmt"

func main() {
    var op string
    var a, b float64

    fmt.Print("Masukkan dua angka: ")
    fmt.Scan(&a, &b)

    fmt.Print("Pilih operator (+, -, *, /): ")
    fmt.Scan(&op)

    switch op {
    case "+":
        fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
    case "-":
        fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
    case "*":
        fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
    case "/":
        if b != 0 {
            fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
        } else {
            fmt.Println("Error: Pembagian dengan nol!")
        }
    default:
        fmt.Println("Operator tidak dikenali!")
    }
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Uji Kelayakan Kadar pH Air Minum
Buatlah program Go untuk menguji kelayakan kadar pH air yang diinputkan:
* Rentang valid pH adalah $0$ sampai $14$. Jika input $< 0$ atau $> 14$, cetak: `"Nilai pH tidak valid. Nilai pH harus antara 0 dan 14."`
* Jika $6.5 \le \text{pH} \le 8.6$, cetak: `"Air layak minum"`
* Jika di luar rentang tersebut (tapi masih dalam $0 - 14$), cetak: `"Air tidak layak minum"`

#### Contoh Uji:
| No | Masukan | Keluaran |
| :---: | :--- | :--- |
| 1 | `8.6` | `Air layak minum` |
| 2 | `9.0` | `Air tidak layak minum` |
| 3 | `16.0` | `Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.` |

---

### Soal 2 – Perhitungan Tarif Parkir Kendaraan
Sebuah tempat parkir menerapkan tarif berdasarkan jenis kendaraan dan lama durasi parkir (dalam jam):
* **Motor**: Rp 2.000,- per jam
* **Mobil**: Rp 5.000,- per jam
* **Truk**: Rp 8.000,- per jam
* Jika jenis kendaraan di luar ketiga jenis di atas, tampilkan pesan: `"Jenis kendaraan tidak dikenal!"`.

*Format Masukan*: String jenis kendaraan (`"motor"`, `"mobil"`, `"truk"`) dan bilangan bulat durasi parkir (jam).
*Format Keluaran*: Total biaya parkir yang harus dibayar.

---

## Kesimpulan Modul 11

* `switch-case` menyajikan sintaks yang bersih untuk percabangan berbasis kecocokan nilai atau menu pilihan.
* Fitur tanpa `break` otomatis di Go membuat kode lebih ringkas dan mengurangi risiko kesalahan logika.
