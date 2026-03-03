# REVIEW ALGORITMA DAN PEMROGRAMAN 1 (GOLANG)

Modul ini membahas **dasar-dasar bahasa Go (Golang)** mulai dari struktur program, tipe data, variabel, operator, hingga latihan logika dasar.

Target setelah belajar modul ini:

* Memahami struktur program Go
* Bisa membuat dan menjalankan program
* Memahami variabel, tipe data, operator
* Mengerti pointer dasar
* Bisa membuat program sederhana berbasis input-output

---

## 2.1 ANATOMI PROGRAM GO

Bayangkan program Go seperti **surat resmi**.
Setiap surat resmi punya bagian wajib. Begitu juga Go.

Struktur dasar program Go:

```go
package main
import "fmt"

func main() {
    fmt.Println("Halo Dunia")
}
```

Mari kita bedah satu per satu.

---

## 1️. package main → "Kop Surat"

```go
package main
```

Artinya:

* Ini adalah **program utama**
* File ini bisa langsung dijalankan

Jika bukan `main`, maka program tidak bisa dieksekusi langsung.

---

## 2️. import → "Meminjam Alat"

```go
import "fmt"
```

`fmt` adalah library bawaan Go untuk:

* Menampilkan teks (Output)
* Menerima input dari keyboard

Tanpa `import`, kita tidak bisa menggunakan alat tersebut.

---

## 3️. func main() → "Isi Surat"

```go
func main() {
    // semua perintah ditulis di sini
}
```

Semua kode yang ingin dijalankan harus berada di dalam `{ }`.

Program selalu mulai dari fungsi `main()`.

---

## Komentar (Catatan untuk Programmer)

Komentar adalah tulisan yang **tidak dibaca komputer**, hanya untuk manusia.

### Komentar satu baris

```go
// Ini komentar satu baris
```

### Komentar banyak baris

```go
/*
Ini komentar
yang panjang
*/
```

Komentar sangat penting saat mengajar atau bekerja tim.

---

## Koding, Kompilasi, dan Eksekusi

Go adalah bahasa **compiled**.

Artinya:

1. Kode diterjemahkan dulu menjadi aplikasi
2. Baru bisa dijalankan

## Langkah Menjalankan Program

### 1️. Tulis kode

Simpan dengan ekstensi `.go`
Contoh: `hello.go`

### 2️. Build (Kompilasi)

Buka Terminal:

```bash
go build hello.go
```

Jika sukses:

* Akan muncul file `hello.exe` (Windows)

### 3️. Jalankan Program

```bash
hello.exe
```

atau

```bash
./hello
```

Jangan klik dua kali file exe dari folder. Jalankan lewat terminal.

---

### Perintah Penting di Terminal

| Perintah | Fungsi                     |
| -------- | -------------------------- |
| go build | Mengubah .go jadi .exe     |
| go fmt   | Merapikan kode otomatis    |
| go clean | Menghapus file hasil build |

---

## 2.2 TIPE DATA DAN INSTRUKSI DASAR

Sekarang masuk ke konsep inti pemrograman.

---

## 1️. Variabel dan Tipe Data

Variabel adalah **kotak penyimpan data**.

Contoh:

```go
var umur int = 20
```

Artinya:

* Buat kotak bernama `umur`
* Isinya angka bulat

---

### Jenis Tipe Data Dasar

| Tipe      | Contoh | Fungsi           |
| --------- | ------ | ---------------- |
| int       | 10     | Bilangan bulat   |
| float64   | 3.14   | Bilangan desimal |
| bool      | true   | Benar/Salah      |
| string    | "Halo" | Teks             |
| byte/rune | 'A'    | Karakter         |

---

### Contoh Program

```go
package main
import "fmt"

func main() {
    var nama string = "Budi"
    var umur int = 20
    var lulus bool = true
    
    fmt.Println("Nama:", nama)
    fmt.Println("Umur:", umur)
    fmt.Println("Lulus:", lulus)
}
```

---

## 2️. Cara Membuat Variabel

Ada 3 cara:

### Cara 1 (Panjang)

```go
var nama string
nama = "Budi"
```

### Cara 2 (Langsung Isi)

```go
var nama string = "Budi"
```

### Cara 3 (Paling Sering Dipakai)

```go
nama := "Budi"
```

Go otomatis menebak tipe datanya.

---

## 3️. Nilai Default (Jika Tidak Diisi)

Jika tidak diberi nilai:

| Tipe   | Nilai Awal |
| ------ | ---------- |
| int    | 0          |
| float  | 0.0        |
| bool   | false      |
| string | ""         |

Contoh:

```go
var dompet int
fmt.Println(dompet) // 0
```

---

## 4️. Operator

### Operator Matematika

| Simbol | Fungsi    |
| ------ | --------- |
| +      | Tambah    |
| -      | Kurang    |
| *      | Kali      |
| /      | Bagi      |
| %      | Sisa bagi |

Contoh:

```go
package main

import (
	"fmt"
)

func main() {
	var a, b int

	// Mengisi nilai a dan b
	a = 20
	b = 6

	// 1. Penambahan (+)
	hasilTambah := a + b
	fmt.Printf("%d + %d = %v\n", a, b, hasilTambah)

	// 2. Pengurangan (-)
	hasilKurang := a - b
	fmt.Printf("%d - %d = %v\n", a, b, hasilKurang)

	// 3. Perkalian (*)
	hasilKali := a * b
	fmt.Printf("%d * %d = %v\n", a, b, hasilKali)

	// 4. Pembagian (/)
	// Perhatikan: Karena a dan b adalah int, maka hasilnya juga int (dibulatkan ke bawah)
	hasilBagi := a / b
	fmt.Printf("%d / %d = %v\n", a, b, hasilBagi)

	// 5. Modulo / Sisa Bagi (%)
	// Sisa bagi dari 20 dibagi 6 adalah 2 (karena 6 * 3 = 18, 20 - 18 = 2)
	hasilSisaBagi := a % b
	fmt.Printf("%d %% %d = %v\n", a, b, hasilSisaBagi)
}
```

---

### Operator Perbandingan

| Simbol | Arti        |
| ------ | ----------- |
| ==     | Sama dengan |
| !=     | Tidak sama  |
| >      | Lebih besar |
| <      | Lebih kecil |

Hasilnya selalu `true` atau `false`.

contoh code : 
```go
package main

import "fmt"

func main() {
	var angka1, angka2, angka3, angka4, angka5 bool
	var a, b int

	a = 10
	b = 2

	angka1 = a != b
	angka2 = a == b
	angka3 = a > b
	angka4 = a < b
	angka5 = a >= b

	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(angka3)
	fmt.Println(angka4)
	fmt.Println(angka5)
}
```

---

### Operator Logika

| Simbol | Arti      |   |      |
| ------ | --------- | - | ---- |
| &&     | DAN       |   |      |
| !      | KEBALIKAN |   |      |

Contoh:

```go
true && false // hasilnya false
```

contoh code :
```go
package main

import "fmt"

func main() {
	var a, b bool
	var angka bool

	a = true
	b = false

	angka = a && b

	fmt.Println(angka)
}
```

---

## 5️. Casting (Aturan Ketat Go)

Go sangat disiplin.

Ini SALAH:

```go
10 + 3.14
```

Ini BENAR:

```go
float64(10) + 3.14
```

contoh code :
```go
package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 10.5
	var angka float64

	// INI AKAN ERROR (mismatched types int and float64)
	angka = a + b 

	// INI BARU BENAR
	angka = float64(a) + b

	fmt.Println(angka)
}
```

Tipe data harus sama sebelum dihitung.

---

## 6. Pointer (Alamat Memori Dasar)

Bayangkan variabel seperti rumah.

* `nama` → isi rumah
* `&nama` → alamat rumah
* `*alamat` → isi dari alamat itu

Contoh:

```go
kotak := "Sepatu"

fmt.Println(kotak)     // Isi
fmt.Println(&kotak)    // Alamat memori
```

Pointer sangat penting saat belajar struktur data nanti.

---

## 7️. Input dan Output

### Output

```go
fmt.Println("Halo")
fmt.Printf("%v", nilai)
```

---

### Input

WAJIB pakai `&`

```go
var angka int
fmt.Scanln(&angka)
```

Kenapa pakai `&`?

Karena kita memberi alamat memori agar data bisa disimpan ke variabel tersebut.

contoh code :
```go
package main

import (
	"fmt"
)

func main() {
	// 1. OUPUT (Menampilkan teks)
	fmt.Println("Halo, selamat datang di program biodata!")

	// Menyiapkan variabel untuk menampung input pengguna
	var nama string
	var umur int

	// 2. INPUT NAMA
	fmt.Print("Masukkan nama Anda: ")
	// WAJIB pakai & karena kita memasukkan data ke alamat memori variabel 'nama'
	fmt.Scanln(&nama) 

	// 3. INPUT UMUR
	fmt.Print("Masukkan umur Anda: ")
	// WAJIB pakai & untuk variabel 'umur'
	fmt.Scanln(&umur)

	// 4. OUTPUT HASIL 
	// Kita bisa menangkap nilainya kembali untuk dicetak
	fmt.Println("====================")
	fmt.Printf("Halo %s!\n", nama)
	fmt.Printf("Wah, usiamu sekarang %v tahun ya.\n", umur)
}
```
---

## 8️. Konstanta

Nilai yang tidak boleh berubah.

```go
const PI = 3.14
```

Digunakan untuk nilai tetap seperti:

* π
* pajak
* diskon tetap

---

# CONTOH PROGRAM LENGKAP

```go
package main
import "fmt"

func main() {
    const PI = 3.14
    var jariJari float64

    fmt.Print("Masukkan jari-jari: ")
    fmt.Scanln(&jariJari)

    luas := PI * jariJari * jariJari

    fmt.Println("Luas lingkaran:", luas)
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

## Bagian 1 – Riset

Cari tahu:

* Pascal → Compiled / Interpreted?
* C++ → ?
* Java → ?
* Python → ?

Tujuan: Memahami perbedaan jenis bahasa pemrograman.

---

## Bagian 2 – Latihan Logika

---

### Soal 1 – Analisis Program

Program ini melakukan apa?

```go
temp = satu
satu = dua
dua = tiga
tiga = temp
```

Petunjuk:
Perhatikan alur perpindahan data.

---

### Soal 2 – Tahun Kabisat

Syarat:

* Habis dibagi 400
  ATAU
* Habis dibagi 4 dan tidak habis dibagi 100

Contoh:

* 2016 → true
* 2018 → false

---

### Soal 3 – Hitung Bola

Rumus:

* Volume = 4/3 × π × r³
* Luas = 4 × π × r²

Gunakan:

```
PI = 3.1415926535
```

---

### Soal 4 – Konversi Suhu

Rumus:

* F = (C × 9/5) + 32
* R = C × 4/5
* K = C + 273.15

---

### Soal 5 – ASCII

1. Ubah angka menjadi huruf
2. Ubah huruf ke huruf setelahnya

Gunakan:

```go
fmt.Scanf("%c", &var)
fmt.Printf("%c", var)
```

---

## KESIMPULAN MODUL 2

Setelah memahami modul ini, kamu sudah menguasai:

* Struktur dasar program Go
* Cara menjalankan program
* Variabel dan tipe data
* Operator matematika & logika
* Casting
* Pointer dasar
* Input dan output

Ini adalah fondasi sebelum masuk ke:

* Percabangan (if)
* Perulangan (for)
* Array
* Struct
* Function lanjutan


---
