# MODUL 4: PROSEDUR (Procedure)

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

- Memahami konsep prosedur
- Membuat prosedur di Golang
- Memanggil prosedur dengan benar
- Memahami parameter (formal & aktual)
- Memahami pass by value & pass by reference

---

## 4.1 Definisi Procedure

### Konsep Dasar

**Prosedur** adalah kumpulan instruksi yang dibungkus menjadi satu fungsi khusus **tanpa nilai kembalian (return)**.

Flownya :
```
        ┌───────────────┐
        │     INPUT     │
        │ (Scan / nilai)│
        └───────┬───────┘
                │ parameter aktual
                ▼
        ┌──────────────┐
        │   PROSEDUR   │
        │ (parameter   │
        │   formal)    │
        └───────┬──────┘
                │ proses (loop, hitung, print)
                ▼
        ┌───────────────┐
        │    OUTPUT     │
        │ (print / ubah │
        │   variabel)   │
        └───────────────┘

```

Tujuan utama:

- Mengurangi kompleksitas program
- Membuat kode lebih rapi & modular
- Mempermudah debugging & reuse code

### Ciri-Ciri Procedure

1. Tidak memiliki tipe return
2. Tidak menggunakan `return`
3. Biasanya menghasilkan efek langsung (output / perubahan data)

---

### Analogi

Bayangkan prosedur seperti:

> "Tombol di mesin kopi ☕"

Saat ditekan → langsung menghasilkan kopi
Tanpa mengembalikan “nilai”, tapi **langsung memberi efek**

---

### Contoh Sederhana Procedure

```go
package main
import "fmt"

// prosedur (tidak ada return)
func cetakHalo() {
    fmt.Println("Halo, Selamat Belajar!")
}

func main() {
    cetakHalo() // memanggil prosedur
}
```

Output:

```
Halo, Selamat Belajar!
```

---

## 4.2 Deklarasi Procedure

### Struktur Umum

#### Pseudocode

```
procedure namaProcedure(parameter)
    deklarasi
    algoritma
endprocedure
```

#### Golang

```go
func namaProcedure(parameter) {
    // isi prosedur
}
```

---

### Catatan Penting

- Ditulis **di luar func main()**
- Bisa sebelum atau sesudah main

---

### Contoh: Cetak Bilangan 1–N

```go
package main
import "fmt"

func cetakAngka(n int) {
    for i := 1; i <= n; i++ {
        fmt.Println(i)
    }
}

func main() {
    cetakAngka(5)
}
```

Output:

```
1
2
3
4
5
```

---

### Contoh: Deret Fibonacci

```go
package main

import "fmt"

func cetakNFibo(n int) {
    var f1, f2, f3 int
    f2 = 0
    f3 = 1

    for i := 1; i <= n; i++ {
        fmt.Println(f3)
        f1 = f2
        f2 = f3
        f3 = f1 + f2
    }
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah deret Fibonacci: ")
    fmt.Scan(&n)

    cetakNFibo(n)
}
```

---

## 4.3 Cara Pemanggilan Procedure

### Konsep

Prosedur **tidak akan berjalan jika tidak dipanggil**

---

### Sintaks Pemanggilan

```go
namaProcedure(argumen)
```

---

### Contoh

```go
package main
import "fmt"

func cetakPesan(nama string) {
    fmt.Println("Halo,", nama)
}

func main() {
    cetakPesan("Budi")   // pakai variabel langsung
    cetakPesan("Andi")   // pakai nilai langsung
}
```

Output:

```
Halo, Budi
Halo, Andi
```

---

## 4.4 Contoh Program Lengkap

### Studi Kasus

Menampilkan pesan berdasarkan flag:

- 0 = error
- 1 = warning
- 2 = informasi

---

### Program Lengkap

```go
package main
import "fmt"

func main() {
    var flag int
    var pesan string

    fmt.Print("Masukkan flag dan pesan: ")
    fmt.Scan(&flag, &pesan)

    cetakPesan(pesan, flag)
}

// prosedur
func cetakPesan(M string, flag int) {
    var jenis string

    if flag == 0 {
        jenis = "ERROR"
    } else if flag == 1 {
        jenis = "WARNING"
    } else if flag == 2 {
        jenis = "INFORMASI"
    } else {
        jenis = "TIDAK DIKETAHUI"
    }

    fmt.Println(jenis + ":", M)
}
```

---

### Contoh Input

```
1 Hello
```

### Output

```
WARNING: Hello
```

---

## 4.5 Parameter

### Konsep

Parameter = cara prosedur menerima data

---

### Jenis Parameter

#### 1. Parameter Formal

Yang ditulis saat deklarasi

```go
func tambah(a int, b int) {
```

---

#### 2. Parameter Aktual

Yang dikirim saat pemanggilan

```go
tambah(5, 3)
```

---

### Contoh Lengkap

```go
package main
import "fmt"

// parameter formal: r dan t
func volumeTabung(r int, t int) float64 {
    return 3.14 * float64(r*r) * float64(t)
}

func main() {
    var hasil float64

    // parameter aktual: 5 dan 10
    hasil = volumeTabung(5, 10)

    fmt.Println("Volume:", hasil)
}
```

---

## 4.6 Pass by Value vs Pass by Reference

---

### 1. Pass by Value

#### Konsep

- Data **disalin**
- Tidak mengubah data asli

---

#### Contoh

```go
package main
import "fmt"

func ubahNilai(x int) {
    x = 100
}

func main() {
    var a int = 10

    ubahNilai(a)

    fmt.Println("Nilai a:", a)
}
```

Output:

```
Nilai a: 10
```

Tidak berubah!

---

### 2. Pass by Reference (Pointer)

#### Konsep

- Menggunakan alamat memori
- Bisa mengubah nilai asli

---

#### Contoh

```go
package main
import "fmt"

func ubahNilai(x *int) {
    *x = 100
}

func main() {
    var a int = 10

    ubahNilai(&a)

    fmt.Println("Nilai a:", a)
}
```

Output:

```
Nilai a: 100
```

Berubah!

---

## Contoh Kombinasi (Function vs Procedure)

### Rumus:

`f(x, y) = 2x - 0.5y + 3`

---

### Implementasi

```go
package main
import "fmt"

// fungsi (return nilai)
func f1(x, y int) float64 {
    return float64(2*x) - 0.5*float64(y) + 3
}

// prosedur (pakai pointer)
func f2(x, y int, hasil *float64) {
    *hasil = float64(2*x) - 0.5*float64(y) + 3
}

func main() {
    var a, b int
    var c float64

    fmt.Print("Masukkan a dan b: ")
    fmt.Scan(&a, &b)

    f2(a, b, &c)

    fmt.Println("Hasil prosedur:", c)
    fmt.Println("Hasil fungsi:", f1(b, a))
}
```

---

## Kesimpulan

- Prosedur = fungsi tanpa return
- Digunakan untuk aksi langsung (print, ubah data)
- Parameter:
  - Formal → di fungsi
  - Aktual → saat dipanggil

- Pass by:
  - Value → aman (tidak ubah data asli)
  - Reference → powerful (bisa ubah data asli)
