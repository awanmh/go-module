# MODUL 15: ASESMEN PRAKTIKUM 2 (EVALUASI AKHIR PRAKTIKUM)

Modul ini berisi **evaluasi komprehensif tahap kedua**, mencakup seluruh materi yang telah dipelajari sepanjang semester: operasi I/O, percabangan (`if-then`, `else-if`, `switch-case`), perulangan (`for`, `while`, `repeat-until`), dan komposisi struktur kontrol.

### Tujuan Pembelajaran

Mahasiswa mampu:

* Menganalisis persoalan komputasi kompleks yang melibatkan kombinasi perulangan dan percabangan
* Mengidentifikasi struktur kontrol yang paling efisien dan tepat untuk suatu kebutuhan
* Mengimplementasikan program Go yang utuh, bersih, dan menangani seluruh kasus batas (*edge cases*)
* Menguji dan memvalidasi kebenaran keluaran program secara mandiri

---

## 15.1 Studi Kasus 1 – Klasifikasi Lengkap Bentuk Segitiga

Diberikan tiga bilangan bulat positif $a, b, c$ yang menyatakan panjang sisi-sisi suatu bangun.
Program harus menentukan apakah ketiga sisi tersebut membentuk:
1. `"Bukan segitiga"`: Jika tidak memenuhi syarat segitiga (jumlah dua sisi mana pun harus lebih besar dari sisi ketiga: $a+b > c$ dan $a+c > b$ dan $b+c > a$).
2. `"Segitiga sama sisi"`: Jika ketiga sisi sama panjang ($a == b$ dan $b == c$).
3. `"Segitiga sama kaki"`: Jika tepat dua sisi sama panjang.
4. `"Segitiga siku-siku"`: Jika memenuhi teorema Pythagoras ($a^2 + b^2 = c^2$ atau $a^2 + c^2 = b^2$ atau $b^2 + c^2 = a^2$).
5. `"Segitiga sembarang"`: Jika ketiga sisi berbeda panjang dan tidak siku-siku.

### Program Solusi:

```go
package main

import "fmt"

func main() {
    var a, b, c int

    fmt.Scan(&a, &b, &c)

    // 1. Validasi keberadaan segitiga
    if a+b <= c || a+c <= b || b+c <= a {
        fmt.Println("Bukan segitiga")
    } else if a == b && b == c {
        fmt.Println("Segitiga sama sisi")
    } else if a == b || a == c || b == c {
        fmt.Println("Segitiga sama kaki")
    } else if (a*a + b*b == c*c) || (a*a + c*c == b*b) || (b*b + c*c == a*a) {
        fmt.Println("Segitiga siku-siku")
    } else {
        fmt.Println("Segitiga sembarang")
    }
}
```

#### Contoh Uji:
| Masukan | Keluaran |
| :---: | :--- |
| `3 3 3` | `Segitiga sama sisi` |
| `5 5 8` | `Segitiga sama kaki` |
| `3 4 5` | `Segitiga siku-siku` |
| `4 5 6` | `Segitiga sembarang` |
| `1 2 3` | `Bukan segitiga` |

---

## 15.2 Studi Kasus 2 – Simulasi Mesin ATM Interaktif

Buatlah program simulasi transaksi perbankan ATM dengan saldo awal Rp 500.000,-:
1. Menampilkan menu interaktif berulang:
   * `1`: Cek Saldo
   * `2`: Setor Tunai
   * `3`: Tarik Tunai (pastikan saldo mencukupi)
   * `4`: Keluar
2. Program terus berjalan hingga pengguna memilih menu `4`.

```go
package main

import "fmt"

func main() {
    saldo := 500000
    pilihan := 0

    for pilihan != 4 {
        fmt.Println("=== MENU ATM ===")
        fmt.Println("1. Cek Saldo")
        fmt.Println("2. Setor Tunai")
        fmt.Println("3. Tarik Tunai")
        fmt.Println("4. Keluar")
        fmt.Print("Pilihan Anda: ")
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            fmt.Printf("Saldo Anda saat ini: Rp %d\n\n", saldo)
        case 2:
            var setor int
            fmt.Print("Jumlah setor tunai (Rp): ")
            fmt.Scan(&setor)
            saldo += setor
            fmt.Printf("Setor berhasil! Saldo baru: Rp %d\n\n", saldo)
        case 3:
            var tarik int
            fmt.Print("Jumlah tarik tunai (Rp): ")
            fmt.Scan(&tarik)
            if tarik <= saldo {
                saldo -= tarik
                fmt.Printf("Penarikan berhasil! Sisa saldo: Rp %d\n\n", saldo)
            } else {
                fmt.Println("Gagal: Saldo tidak mencukupi!\n")
            }
        case 4:
            fmt.Println("Terima kasih telah bertransaksi.")
        default:
            fmt.Println("Pilihan tidak valid, silakan coba lagi.\n")
        }
    }
}
```

---

## Kesimpulan Modul 15

* Asesmen Praktikum 2 menandai penguasaan lengkap atas struktur kontrol dasar dalam algoritma pemrograman.
* Keahlian memadukan percabangan dan perulangan secara modular mempersiapkan mahasiswa melangkah ke mata kuliah Algoritma & Pemrograman 2 (fungsi, prosedur, rekursif, struct, array, searching, dan sorting).
