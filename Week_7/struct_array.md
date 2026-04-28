# MODUL 7: STRUCT & ARRAY

### Tujuan Pembelajaran

Mahasiswa mampu:

* Memahami tipe bentukan (alias & struct)
* Menggunakan struct untuk data kompleks
* Menggunakan array, slice, dan map di Golang
* Memilih struktur data yang tepat

---

## 7.1 Tipe Bentukan (User Defined Type)

### Apa itu Tipe Bentukan?

Tipe bentukan adalah:

> Tipe data baru yang dibuat oleh programmer

---

### Kenapa Penting?

Karena:

* Membuat kode lebih **rapi & mudah dibaca**
* Memodelkan data dunia nyata (mahasiswa, waktu, dll)

---

### 1. Alias (Type)

#### Konsep

Alias = memberi **nama baru untuk tipe lama**

---

#### Analogi

```text
int → bilangan
float64 → pecahan
```

Sama saja, cuma beda nama

---

#### Contoh

```go
package main
import "fmt"

// alias
type bilangan int
type pecahan float64

func main() {
    var a, b bilangan
    var hasil pecahan

    a = 9
    b = 5

    hasil = pecahan(a) / pecahan(b)

    fmt.Println("Hasil:", hasil)
}
```

Output:

```
Hasil: 1.8
```

---

#### Insight Penting

* Alias **tidak membuat tipe baru secara struktur**
* Hanya mempermudah penamaan

---

### 2. Struct (Record)

#### Konsep

Struct = kumpulan beberapa data dalam 1 kesatuan

---

#### Alur Penggunaan Struct

```text
        ┌──────────────┐
        │ Definisikan  │
        │ Tipe Struct  │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Deklarasi    │
        │ Variabel     │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Isi Data ke  │
        │ Tiap Field   │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Akses Data   │
        │ dari Field   │
        └──────────────┘
```

---

#### Analogi

```text
Mahasiswa:
- nama
- umur
- nilai
```

Disimpan dalam 1 variabel

---

#### Struktur

```go
type NamaStruct struct {
    field1 tipe
    field2 tipe
}
```

---

#### Contoh Sederhana

```go
package main
import "fmt"

// struct
type Mahasiswa struct {
    nama  string
    umur  int
}

func main() {
    var m Mahasiswa

    m.nama = "Budi"
    m.umur = 20

    fmt.Println(m.nama, m.umur)
}
```

Output:

```
Budi 20
```

---

#### Contoh Real Case: Waktu Parkir

```go
package main
import "fmt"

type waktu struct {
    jam, menit, detik int
}

func main() {
    var masuk, keluar, durasi waktu
    var totalMasuk, totalKeluar, selisih int

    fmt.Println("Masukkan waktu masuk (jam menit detik):")
    fmt.Scan(&masuk.jam, &masuk.menit, &masuk.detik)

    fmt.Println("Masukkan waktu keluar (jam menit detik):")
    fmt.Scan(&keluar.jam, &keluar.menit, &keluar.detik)

    totalMasuk = masuk.jam*3600 + masuk.menit*60 + masuk.detik
    totalKeluar = keluar.jam*3600 + keluar.menit*60 + keluar.detik

    selisih = totalKeluar - totalMasuk

    durasi.jam = selisih / 3600
    durasi.menit = (selisih % 3600) / 60
    durasi.detik = selisih % 60

    fmt.Printf("Lama parkir: %d jam %d menit %d detik\n",
        durasi.jam, durasi.menit, durasi.detik)
}
```

---

## 7.2 ARRAY

### Konsep

Array = kumpulan data dengan:

* Tipe sama
* Ukuran tetap

---

### Alur Penggunaan Array

```text
        ┌──────────────┐
        │ Tentukan Tipe│
        │ & Ukuran     │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Deklarasi    │
        │ Variabel     │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Isi Data ke  │
        │ Tiap Index   │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Akses Data   │
        │ via Index    │
        └──────────────┘
```

---

### Analogi

```text
[10, 20, 30, 40]
```

1 variabel, banyak isi

---

### Deklarasi Array

```go
var arr [5]int
```

---

### Contoh

```go
package main
import "fmt"

func main() {
    var angka [5]int

    angka[0] = 10
    angka[1] = 20

    fmt.Println(angka)
}
```

📌 Output:

```
[10 20 0 0 0]
```

---

### Panjang Array

```go
len(arr)
```

---

### Contoh

```go
fmt.Println(len(angka))
```

---

### Indexing

* Dimulai dari 0
* Elemen terakhir = len(array) - 1

---

### Contoh Akses

```go
package main
import "fmt"

func main() {
    arr := [3]int{10, 20, 30}

    fmt.Println("Elemen pertama:", arr[0])
    fmt.Println("Elemen terakhir:", arr[len(arr)-1])
}
```

---

## SLICE (Array Dinamis)

### Konsep

Slice = array yang bisa berubah ukuran

---

### Alur Penggunaan Slice

```text
        ┌──────────────┐
        │ Deklarasi    │
        │ Slice        │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Tambah Data  │
        │ (append)     │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Kapasitas    │
        │ Penuh?       │─── Ya ──► Alokasi Array Baru
        └───────┬──────┘           (Otomatis)
          Tidak │
                ▼
        ┌──────────────┐
        │ Akses Data   │
        │ via Index    │
        └──────────────┘
```

---

### Deklarasi

```go
var s []int
```

---

### Contoh

```go
package main
import "fmt"

func main() {
    s := []int{1, 2, 3}

    s = append(s, 4)

    fmt.Println(s)
}
```

Output:

```
[1 2 3 4]
```

---

### Fungsi Penting

#### len() → jumlah data

```go
len(s)
```

#### cap() → kapasitas

```go
cap(s)
```

---

### Contoh

```go
fmt.Println(len(s), cap(s))
```

---

### Slice dari Array

```go
arr := [5]int{1,2,3,4,5}
s := arr[1:4]
```

Output:

```
[2 3 4]
```

---

## MAP (Array dengan Key)

### Konsep

Map = array dengan key (bukan index angka)

---

### Alur Penggunaan Map

```text
        ┌──────────────┐
        │ Inisialisasi │
        │ Map (make)   │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Tentukan     │
        │ Key & Value  │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Simpan Data  │
        │ (m[key]=val) │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Akses Data   │
        │ via Key      │
        └──────────────┘
```

---

### Analogi

```text
nama → nilai
```

---

### Deklarasi

```go
var m map[string]int
```

---

### Contoh

```go
package main
import "fmt"

func main() {
    nilai := map[string]int{
        "Budi": 90,
        "Ani":  85,
    }

    fmt.Println(nilai["Budi"])
}
```

Output:

```
90
```

---

### Operasi Map

#### Tambah / Update

```go
nilai["Budi"] = 95
```

#### Hapus

```go
delete(nilai, "Budi")
```

---

### Contoh Lengkap

```go
package main
import "fmt"

func main() {
    data := make(map[string]string)

    data["john"] = "hi"
    data["anne"] = "hello"

    fmt.Println(data["john"])

    data["anne"] = "updated"
    data["boy"] = "new"

    delete(data, "john")

    fmt.Println(data)
}
```

---

### Perbandingan Struktur Data

| Tipe   | Ukuran   | Akses | Kelebihan       |
| ------ | -------- | ----- | --------------- |
| Array  | Tetap    | Index | Cepat           |
| Slice  | Dinamis  | Index | Fleksibel       |
| Map    | Dinamis  | Key   | Powerful        |
| Struct | Kompleks | Field | Real-world data |

---

### Kesimpulan

* Alias → hanya nama baru
* Struct → gabungan data
* Array → ukuran tetap
* Slice → fleksibel
* Map → key-value

---