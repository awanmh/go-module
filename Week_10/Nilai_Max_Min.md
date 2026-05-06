# MODUL 10: PENCARIAN NILAI EKSTRIM

### Tujuan Pembelajaran

Mahasiswa mampu:

* Memahami konsep pencarian nilai maksimum dan minimum
* Mengimplementasikan algoritma pencarian ekstrim
* Mencari nilai ekstrim pada array sederhana
* Mencari nilai ekstrim pada data terstruktur (struct)
* Memahami perbedaan pencarian nilai vs indeks

---

## 10.1 Ide Pencarian Nilai Max/Min

### Apa itu Pencarian Nilai Ekstrim?

Pencarian nilai ekstrim adalah:

> Proses mencari nilai **terbesar (max)** atau **terkecil (min)** dari sekumpulan data

---

### Kenapa Penting?

Karena:

* Digunakan di banyak kasus nyata (nilai tertinggi, harga termurah, dll)
* Dasar dari banyak algoritma lanjutan
* Melatih logika iterasi dan perbandingan

---

### Analogi

```text
Nilai ujian:
[70, 85, 90, 60, 88]

→ nilai tertinggi = 90
→ nilai terendah = 60
```

---

### Ide Dasar Algoritma

```text
        ┌────────────────────┐
        │ Ambil data pertama │
        │ sebagai nilai awal │
        └─────────┬──────────┘
                  ▼
        ┌────────────────────┐
        │ Bandingkan dengan  │
        │ data berikutnya    │
        └─────────┬──────────┘
                  ▼
        ┌────────────────────┐
        │ Jika lebih besar / │
        │ kecil → update     │
        └─────────┬──────────┘
                  ▼
        ┌────────────────────┐
        │ Ulangi sampai akhir│
        └─────────┬──────────┘
                  ▼
        ┌────────────────────┐
        │ Dapat nilai ekstrim│
        └────────────────────┘
```

---

### Pseudocode

```text
max = indeks pertama
i = indeks berikutnya

selama i < n:
    jika data[i] > data[max]:
        max = i
    i++
```

---

### Versi Go

```go
max := 0
i := 1

for i < n {
    if a[i] > a[max] {
        max = i
    }
    i++
}
```

---

### Insight Penting

* Selalu mulai dari data pertama
* Bandingkan satu per satu (sekuensial)
* Bisa mencari:

  * nilai (value)
  * atau posisi (index)

---

## 10.2 Pencarian Nilai Ekstrim pada Array

### Konsep

Array digunakan untuk:

> Menyimpan kumpulan data yang akan dicari nilai ekstrimnya

---

### Alur Penggunaan

```text
        ┌──────────────┐
        │ Siapkan Data │
        │ (Array)      │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Ambil Nilai  │
        │ Awal         │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Iterasi      │
        │ (Loop)       │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Bandingkan   │
        │ & Update     │
        └───────┬──────┘
                ▼
        ┌──────────────┐
        │ Hasil Final  │
        └──────────────┘
```

---

### Contoh: Mencari Nilai Minimum

```go
package main
import "fmt"

type arrInt [2023]int

func terkecil(tabInt arrInt, n int) int {
    var min int = tabInt[0]
    var j int = 1

    for j < n {
        if min > tabInt[j] {
            min = tabInt[j]
        }
        j++
    }
    return min
}

func main() {
    data := arrInt{10, 5, 8, 3, 12}
    fmt.Println("Nilai terkecil:", terkecil(data, 5))
}
```

---

### Contoh: Mencari Index Nilai Minimum

```go
func terkecilIndex(tabInt arrInt, n int) int {
    var idx int = 0
    var j int = 1

    for j < n {
        if tabInt[idx] > tabInt[j] {
            idx = j
        }
        j++
    }
    return idx
}
```

---

### Perbedaan Penting

| Jenis | Output     |
| ----- | ---------- |
| Nilai | 3          |
| Index | 3 (posisi) |

---

### Insight Penting

* Nilai → langsung angka
* Index → bisa akses data lengkap

---

## 10.3 Pencarian Nilai Ekstrim pada Struct

### Konsep

Pencarian juga bisa dilakukan pada:

> Data kompleks (struct)

---

### Analogi

```text
Mahasiswa:
- nama
- nim
- ipk

→ cari IPK tertinggi
```

---

### Struktur Data

```go
type mahasiswa struct {
    nama, nim, kelas, jurusan string
    ipk float64
}
```

---

### Array Struct

```go
type arrMhs [2023]mahasiswa
```

---

### Contoh: Mencari IPK Tertinggi

```go
func IPKMax(T arrMhs, n int) float64 {
    var tertinggi float64 = T[0].ipk
    var j int = 1

    for j < n {
        if tertinggi < T[j].ipk {
            tertinggi = T[j].ipk
        }
        j++
    }
    return tertinggi
}
```

---

### Contoh: Mencari Index Mahasiswa IPK Tertinggi

```go
func IPKIndex(T arrMhs, n int) int {
    var idx int = 0
    var j int = 1

    for j < n {
        if T[idx].ipk < T[j].ipk {
            idx = j
        }
        j++
    }
    return idx
}
```

---

### Mengambil Data Lengkap

```go
idx := IPKIndex(data, n)

fmt.Println("Nama:", data[idx].nama)
fmt.Println("IPK:", data[idx].ipk)
```

---

### Insight Penting

* Struct memungkinkan data lebih kompleks
* Gunakan index jika ingin:

  * akses seluruh data
  * bukan hanya nilai

---

## Perbandingan Pendekatan

| Pendekatan | Kelebihan | Kekurangan           |
| ---------- | --------- | -------------------- |
| Nilai      | Sederhana | Tidak tahu posisi    |
| Index      | Fleksibel | Perlu akses tambahan |

---

## Kesimpulan

* Pencarian ekstrim = mencari nilai max/min
* Algoritma dasar:

  * ambil data pertama
  * bandingkan semua data
* Bisa diterapkan pada:

  * Array sederhana
  * Struct kompleks
* Gunakan index jika butuh data lengkap

---