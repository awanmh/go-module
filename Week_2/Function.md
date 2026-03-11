
# FUNGSI (FUNCTION)

## 3.1 Definisi Function

### Pengertian Function

**Function (fungsi)** adalah sekumpulan instruksi atau perintah dalam program yang dirancang untuk **mengolah suatu input dan menghasilkan sebuah nilai keluaran (output)**.

Secara sederhana, fungsi dapat dianggap seperti **mesin pengolah**:

```
Input  →  Proses di dalam fungsi  →  Output
```

Contoh dalam kehidupan sehari-hari:

| Input               | Proses              | Output          |
| ------------------- | ------------------- | --------------- |
| angka 5             | dihitung kuadratnya | 25              |
| jari-jari lingkaran | dihitung luasnya    | luas lingkaran  |
| angka mahasiswa     | dihitung rata-rata  | nilai rata-rata |

Dengan kata lain, **fungsi memetakan suatu nilai ke nilai lain**.

Misalnya:

```
f(x) = x²
```

Jika
x = 3 → maka f(x) = 9

---

### Ciri-ciri Function

Sebuah subprogram disebut **function** apabila memiliki dua karakteristik utama:

1. **Memiliki tipe nilai yang dikembalikan (return type)**
   Artinya fungsi harus menghasilkan suatu nilai tertentu.

2. **Memiliki perintah `return`**
   Perintah ini digunakan untuk mengembalikan hasil perhitungan dari fungsi ke program yang memanggilnya.

---

### Kapan Menggunakan Function?

Fungsi digunakan ketika **hasil perhitungan diperlukan kembali di program**.

Contoh penggunaan fungsi:

1. **Mengisi nilai variabel**

```
luas = hitungLuasLingkaran(7)
```

2. **Menjadi bagian dari ekspresi**

```
total = hitungDiskon(harga) + pajak
```

3. **Menjadi argumen fungsi lain**

```
fmt.Println(hitungRataRata(10,20,30))
```

---

### Aturan Penamaan Function

Nama fungsi sebaiknya **menggambarkan nilai yang dihasilkan**.

Biasanya berupa:

* **kata benda**
* **kata sifat**
* **frasa yang jelas**

Contoh nama fungsi yang baik:

```
median
rerata
nilaiTerbesar
volumeTabung
luasLingkaran
cekKelulusan
```

Contoh nama fungsi yang **kurang baik**:

```
proses1
fungsiA
hitung123
```

---

### Contoh Ilustrasi Function

Misalnya kita membuat fungsi untuk menghitung **kuadrat suatu bilangan**.

#### Pseudocode

```
function kuadrat(x : integer) -> integer
algoritma
return x * x
endfunction
```

#### Contoh Pemanggilan

```
hasil = kuadrat(5)
```

Output:

```
25
```

---

### Contoh Program Sederhana (Go)

```go
package main
import "fmt"

func kuadrat(x int) int {
    return x * x
}

func main() {
    var hasil int
    hasil = kuadrat(5)

    fmt.Println("Hasil kuadrat:", hasil)
}
```

Output

```
Hasil kuadrat: 25
```

---

# 3.2 Deklarasi Function

### Pengertian Deklarasi Function

**Deklarasi fungsi** adalah proses mendefinisikan fungsi agar dapat digunakan dalam program.

Biasanya fungsi diletakkan **terpisah dari program utama (main)** agar program lebih rapi dan mudah dipelihara.

---

### Struktur Function (Algoritma)

```
function <nama function> (<parameter>) -> <tipe data>

kamus
    {deklarasi variabel lokal}

algoritma
    {proses perhitungan}

return <nilai yang dikembalikan>

endfunction
```

Penjelasan:

| Bagian        | Fungsi             |
| ------------- | ------------------ |
| nama function | nama fungsi        |
| parameter     | nilai input        |
| tipe data     | tipe nilai output  |
| kamus         | variabel lokal     |
| algoritma     | proses perhitungan |
| return        | nilai hasil        |

---

### Struktur Function dalam Go

```
func namaFungsi(parameter tipeData) tipeDataReturn {
    // deklarasi variabel
    // proses
    return nilai
}
```

---

### Contoh Function: Menghitung Volume Tabung

Rumus volume tabung:

```
Volume = π × r² × t
```

Dimana:

* r = jari-jari
* t = tinggi

---

### Pseudocode

```
function volumeTabung(jari_jari, tinggi : integer) -> real

kamus
luasAlas, volume : real

algoritma
luasAlas <- 3.14 * (jari_jari * jari_jari)
volume <- luasAlas * tinggi

return volume

endfunction
```

---

### Implementasi Go

```go
package main
import "fmt"

func volumeTabung(jari_jari, tinggi int) float64 {

    var luasAlas float64
    var volume float64

    luasAlas = 3.14 * float64(jari_jari*jari_jari)
    volume = luasAlas * float64(tinggi)

    return volume
}

func main() {

    hasil := volumeTabung(7,10)

    fmt.Println("Volume Tabung:", hasil)

}
```

Output:

```
Volume Tabung: 1538.6
```

---

### Ilustrasi Cara Kerja Fungsi

```
Program Main
     |
     | memanggil
     v
volumeTabung(7,10)
     |
     | proses perhitungan
     v
1538.6
     |
     | return
     v
Program Main menerima hasil
```

---

# 3.3 Cara Pemanggilan Function

### Pengertian Pemanggilan Function

**Pemanggilan fungsi** adalah proses menjalankan fungsi yang telah dibuat dengan memberikan nilai parameter.

Penulisan pemanggilan fungsi:

```
namaFungsi(argumen)
```

---

### Perbedaan Function dan Prosedur

| Prosedur                   | Function              |
| -------------------------- | --------------------- |
| tidak mengembalikan nilai  | mengembalikan nilai   |
| dipanggil sebagai perintah | bisa menjadi ekspresi |

---

### Cara Pemanggilan Function

Fungsi dapat digunakan dalam beberapa cara:

---

### 1. Disimpan ke variabel

Pseudocode

```
v1 <- volumeTabung(r,t)
```

Go

```go
v1 = volumeTabung(r,t)
```

---

### 2. Digunakan dalam ekspresi

Pseudocode

```
v2 <- volumeTabung(r,t) + volumeTabung(15,t)
```

Go

```go
v2 = volumeTabung(r,t) + volumeTabung(15,t)
```

---

### 3. Langsung ditampilkan

Pseudocode

```
output(volumeTabung(14,100))
```

Go

```go
fmt.Println(volumeTabung(14,100))
```

---

### Contoh Program Lengkap

```go
package main
import "fmt"

func volumeTabung(r,t int) float64 {

    return 3.14 * float64(r*r) * float64(t)

}

func main() {

    var r,t int
    var v1,v2 float64

    r = 5
    t = 10

    v1 = volumeTabung(r,t)

    v2 = volumeTabung(r,t) + volumeTabung(15,t)

    fmt.Println("Volume 1 =", v1)
    fmt.Println("Volume 2 =", v2)

    fmt.Println("Volume langsung =", volumeTabung(14,100))

}
```

---

# 3.4 Contoh Program dengan Function

Pada bagian ini kita akan membuat program yang menggunakan dua fungsi:

1. **fungsi faktorial**
2. **fungsi permutasi**

---

## Konsep Faktorial

Faktorial dilambangkan dengan **!**

```
n! = n × (n-1) × (n-2) × ... × 1
```

Contoh:

```
5! = 5 × 4 × 3 × 2 × 1 = 120
```

---

## Konsep Permutasi

Permutasi digunakan untuk menghitung **banyaknya susunan objek**.

Rumus:

```
P(n,r) = n! / (n-r)!
```

Contoh:

```
P(5,2) = 5! / (3!) = 20
```

---

## Struktur Program

Program akan:

1. meminta input **a dan b**
2. jika **a ≥ b** maka hitung P(a,b)
3. jika **b > a** maka hitung P(b,a)

---

## Program Lengkap

```go
package main

import "fmt"

func main(){

    var a,b int

    fmt.Print("Masukkan dua bilangan: ")
    fmt.Scan(&a, &b)

    if a >= b {
        fmt.Println("Hasil Permutasi:", permutasi(a,b))
    }else{
        fmt.Println("Hasil Permutasi:", permutasi(b,a))
    }

}

func faktorial(n int) int{

    var hasil int = 1

    for i := 1; i <= n; i++ {

        hasil = hasil * i

    }

    return hasil
}

func permutasi(n,r int) int {

    return faktorial(n) / faktorial(n-r)

}
```

---

## Contoh Eksekusi Program

Input

```
2 5
```

Karena

```
5 ≥ 2
```

Program menghitung

```
P(5,2)
```

Perhitungan

```
5! = 120
3! = 6

120 / 6 = 20
```

Output

```
20
```

---

## Alur Pemanggilan Fungsi

```
main()
  |
  | memanggil
  v
permutasi(n,r)
  |
  | memanggil
  v
faktorial(n)
faktorial(n-r)
  |
  | return
  v
permutasi menghitung hasil
  |
  | return
  v
main menampilkan hasil
```

---

# Kesimpulan Modul

Function merupakan konsep penting dalam pemrograman karena:

1. **membuat program lebih terstruktur**
2. **menghindari pengulangan kode**
3. **memudahkan pemeliharaan program**
4. **memecah program besar menjadi bagian kecil**

Prinsip utama fungsi:

```
Input → Proses → Output
```

Jika suatu bagian program **menghasilkan nilai**, maka sangat disarankan menggunakan **function**.

---
