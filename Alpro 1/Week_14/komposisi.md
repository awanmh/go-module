# MODUL 14: STRUKTUR KONTROL - KOMPOSISI (NESTED CONTROL STRUCTURES)

Modul ini membahas **komposisi struktur kontrol**, yaitu penggabungan (*nesting*) antara struktur percabangan dan perulangan (perulangan di dalam perulangan, percabangan di dalam perulangan, dan sebaliknya) untuk memecahkan persoalan yang kompleks.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami konsep komposisi dan penyarangan (*nesting*) struktur kontrol
* Mengimplementasikan perulangan bersarang (*nested loops*) untuk pola matriks dan bidang 2 dimensi
* Menggabungkan struktur percabangan di dalam perulangan untuk proses filtrasi data
* Mengimplementasikan algoritma pengujian sifat bilangan (seperti pengecekan bilangan prima)

---

## 14.1 Konsep Dasar Komposisi

Dalam pemrograman terstruktur:
> **Komposisi** adalah teknik meletakkan satu atau lebih struktur kontrol (percabangan atau perulangan) di dalam blok struktur kontrol lainnya.

### Variasi Komposisi yang Umum:
1. **Perulangan di dalam Perulangan (*Nested Loop*)**: Mengolah koordinat baris dan kolom (matriks, pola gambar).
2. **Percabangan di dalam Perulangan (*Branching inside Loop*)**: Memeriksa kondisi khusus pada setiap iterasi (filter ganjil/genap, mencari nilai batas).
3. **Perulangan di dalam Percabangan (*Loop inside Branching*)**: Menjalankan perulangan hanya jika kondisi prasyarat terpenuhi.

```text
for baris := 1; baris <= tinggi; baris++ {
    for kolom := 1; kolom <= lebar; kolom++ {
        // Dieksekusi sebanyak (tinggi * lebar) kali!
    }
}
```

---

## 14.2 Contoh Program Lengkap

### Program 1 – Menampilkan Barisan Bilangan Ganjil
Program membaca bilangan bulat positif $N$ dan menampilkan seluruh bilangan ganjil dari 1 sampai $N$:

```go
package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan batas n: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        // Komposisi: percabangan if di dalam perulangan for
        if i % 2 != 0 {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()
}
```

#### Contoh Eksekusi:
```text
Masukan : 7
Keluaran: 1 3 5 7
```

---

### Program 2 – Pola Persegi Bintang
Mencetak pola persegi berukuran $N \times N$ menggunakan *nested loop*:

```go
package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan ukuran n: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            fmt.Print("* ")
        }
        fmt.Println() // Pindah baris
    }
}
```

#### Contoh Eksekusi ($N = 3$):
```text
* * * 
* * * 
* * * 
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Menghitung Banyaknya Bilangan Ganjil
Buatlah program Go untuk menghitung berapa banyak bilangan ganjil yang terdapat dari $1$ hingga $n$.
* **Masukan**: Sebuah bilangan bulat positif $n$.
* **Keluaran**: Teks yang menyatakan banyaknya bilangan ganjil.
* *Catatan*: Gunakan perulangan dan percabangan untuk memeriksa bilangan (bukan rumus langsung).

#### Contoh Uji:
| Masukan ($n$) | Keluaran |
| :---: | :--- |
| `3` | `Terdapat 2 bilangan ganjil` |
| `2` | `Terdapat 1 bilangan ganjil` |
| `7` | `Terdapat 4 bilangan ganjil` |
| `10` | `Terdapat 5 bilangan ganjil` |

---

### Soal 2 – Pengujian Bilangan Prima
Sebuah bilangan bulat $n > 1$ dikatakan **prima** jika hanya memiliki tepat dua faktor pembagi positif, yaitu 1 dan bilangan itu sendiri (bilangan 1 bukan prima).
* **Masukan**: Sebuah bilangan bulat positif $n$.
* **Keluaran**: Teks `"prima"` jika bilangan tersebut prima, atau `"bukan prima"` jika bukan.

#### Contoh Uji:
| Masukan | Keluaran |
| :---: | :--- |
| `5` | `prima` |
| `12` | `bukan prima` |
| `19` | `prima` |
| `72` | `bukan prima` |

---

### Soal 3 – Pola Segitiga Bintang Siku-Siku
Buatlah program yang menerima input bilangan bulat positif $n$, kemudian mencetak segitiga siku-siku dengan tinggi $n$ baris:
```text
Masukan : 4
Keluaran:
*
* *
* * *
* * * *
```

---

## Kesimpulan Modul 14

* Komposisi adalah fondasi untuk menyelesaikan persoalan komputasi dua dimensi dan logika bertingkat.
* Perhatian ekstra diperlukan dalam memperhatikan *scope* variabel dan kondisi terminasi pada setiap lapis perulangan bersarang.
