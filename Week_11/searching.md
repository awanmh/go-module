# MODUL 11: PENCARIAN NILAI ACAK PADA HIMPUNAN DATA

### Tujuan Pembelajaran

Mahasiswa mampu:

* Memahami konsep pencarian data pada himpunan data
* Memahami algoritma Sequential Search
* Memahami algoritma Binary Search
* Mengimplementasikan pencarian pada array dan struct
* Memahami perbedaan pencarian linear dan pencarian berbasis keterurutan data
* Menentukan kapan menggunakan Sequential Search atau Binary Search

---

# 11.1 Konsep Dasar Pencarian Data

## Apa itu Pencarian Data?

Pencarian data adalah:

> Proses menemukan suatu data tertentu dari sekumpulan data yang tersedia

---

## Perbedaan dengan Pencarian Nilai Ekstrim

| Pencarian Ekstrim               | Pencarian Data            |
| ------------------------------- | ------------------------- |
| Selalu ada hasil                | Bisa tidak ditemukan      |
| Mencari max/min                 | Mencari data spesifik     |
| Fokus pada nilai terbesar/kecil | Fokus pada kecocokan data |

---

## Contoh Kasus Nyata

```text
- Mencari mahasiswa berdasarkan NIM
- Mencari nama pada daftar hadir
- Mencari produk berdasarkan kode
- Mencari alamat tertentu
```

---

## Kemungkinan dalam Pencarian

```text
Data ditemukan
        atau
Data tidak ditemukan
```

Karena itu biasanya digunakan:

```go
- boolean (true/false)
atau
- index (-1 jika tidak ditemukan)
```

---

# 11.2 Sequential Search

## Konsep Sequential Search

Sequential Search adalah:

> Teknik pencarian data dengan mengecek elemen satu per satu secara berurutan dari awal hingga akhir

---

## Analogi

```text
Mencari nama di daftar hadir:

Andi
Budi
Caca
Dina

→ cek dari atas satu-persatu
```

---

## Ciri Sequential Search

* Data dicek dari awal
* Tidak membutuhkan data terurut
* Berhenti jika data ditemukan
* Sederhana dan mudah dibuat

---

## Alur Sequential Search

```text
        ┌──────────────────┐
        │ Mulai dari index │
        │ pertama          │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Bandingkan data  │
        │ dengan target    │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Cocok?           │
        └───────┬──────────┘
                │
      Ya ◄──────┘──────► Tidak
      │                        │
      ▼                        ▼
┌─────────────┐      ┌────────────────┐
│ Data ketemu │      │ Lanjut ke data │
└─────────────┘      │ berikutnya     │
                     └──────┬─────────┘
                            ▼
                   ┌────────────────┐
                   │ Sampai akhir   │
                   └────────────────┘
```

---

## Pseudocode Sequential Search

```text
found = false
i = 0

selama i < n dan belum ditemukan:
    found = data[i] == X
    i++
```

---

## Versi Go

```go
found := false
i := 0

for i < n && !found {
    found = T[i] == X
    i++
}
```

---

## Contoh Sequential Search pada Array String

```go
package main
import "fmt"

type arrStr [1234]string

func SeqSearch(T arrStr, n int, X string) bool {

    var found bool = false
    var j int = 0

    for j < n && !found {
        found = T[j] == X
        j++
    }

    return found
}

func main() {

    data := arrStr{
        "Andi",
        "Budi",
        "Caca",
        "Dina",
    }

    fmt.Println(SeqSearch(data, 4, "Caca"))
}
```

---

## Sequential Search Mengembalikan Index

### Kenapa Index Penting?

Karena dengan index kita bisa:

* mengetahui posisi data
* mengambil seluruh data terkait

---

## Contoh Program

```go
func SeqSearchIndex(T arrStr, n int, X string) int {

    var found int = -1
    var j int = 0

    for j < n && found == -1 {

        if T[j] == X {
            found = j
        }

        j++
    }

    return found
}
```

---

## Penjelasan

| Kondisi    | Arti                        |
| ---------- | --------------------------- |
| found = -1 | data belum ditemukan        |
| found = j  | data ditemukan pada index j |

---

## Variasi Sequential Search

```go
func SeqSearchIndex(T arrStr, n int, X string) int {

    var j int = 0

    for j < n-1 && T[j] != X {
        j++
    }

    if T[j] == X {
        return j
    } else {
        return -1
    }
}
```

---

## Insight Penting

* Sequential Search cocok untuk data kecil
* Tidak memerlukan data terurut
* Semakin besar data → semakin lambat

---

# 11.3 Binary Search

## Konsep Binary Search

Binary Search adalah:

> Teknik pencarian dengan membagi area pencarian menjadi dua bagian

---

## Syarat Binary Search

### Data HARUS Terurut

```text
Ascending  : 1 2 3 4 5
Descending : 5 4 3 2 1
```

Jika data tidak terurut:

```text
Binary Search TIDAK bisa digunakan
```

---

## Ide Dasar Binary Search

```text
1. Ambil data tengah
2. Bandingkan dengan target
3. Jika target lebih kecil:
      cari ke kiri
4. Jika target lebih besar:
      cari ke kanan
5. Ulangi
```

---

## Analogi

```text
Mencari kata di kamus:

→ buka bagian tengah
→ lihat huruf
→ geser kiri/kanan
```

---

## Ilustrasi Binary Search

```text
Data:
[1 3 5 7 9 11 13]

Cari: 9

Tengah = 7
9 > 7 → geser kanan

[9 11 13]

Tengah = 11
9 < 11 → geser kiri

[9]

Ketemu
```

---

## Diagram Binary Search

```text
        ┌───────────────────┐
        │ Ambil data tengah │
        └─────────┬─────────┘
                  ▼
        ┌───────────────────┐
        │ Bandingkan target │
        └───────┬───────────┘
                │
      ┌─────────┼─────────┐
      ▼                   ▼
 Target kecil       Target besar
      │                   │
      ▼                   ▼
 Cari kiri          Cari kanan
      │                   │
      └─────────┬─────────┘
                ▼
         Ulangi proses
```

---

## Pseudocode Binary Search

```text
kr = 0
kn = n-1
found = false

selama kr <= kn dan belum ditemukan:

    med = (kr + kn) / 2

    jika data[med] < X:
        kr = med + 1

    jika data[med] > X:
        kn = med - 1

    selain itu:
        found = true
```

---

## Versi Go

```go
kr := 0
kn := n - 1
found := false

for kr <= kn && !found {

    med := (kr + kn) / 2

    if a[med] < X {
        kr = med + 1

    } else if a[med] > X {
        kn = med - 1

    } else {
        found = true
    }
}
```

---

## Binary Search Descending

## Contoh Program

```go
package main
import "fmt"

type arrInt [4321]int

func BinarySearch(T arrInt, n int, X int) bool {

    var found bool = false
    var med int

    var kr int = 0
    var kn int = n - 1

    for kr <= kn && !found {

        med = (kr + kn) / 2

        if X > T[med] {

            kn = med - 1

        } else if X < T[med] {

            kr = med + 1

        } else {

            found = true
        }
    }

    return found
}
```

---

## Binary Search Mengembalikan Index

```go
func BinarySearchIndex(T arrInt, n int, X int) int {

    var found int = -1
    var med int

    var kr int = 0
    var kn int = n - 1

    for kr <= kn && found == -1 {

        med = (kr + kn) / 2

        if X > T[med] {

            kn = med - 1

        } else if X < T[med] {

            kr = med + 1

        } else {

            found = med
        }
    }

    return found
}
```

---

## Insight Penting

| Sequential Search      | Binary Search  |
| ---------------------- | -------------- |
| Tidak perlu terurut    | Harus terurut  |
| Lambat pada data besar | Sangat cepat   |
| Mudah dibuat           | Lebih kompleks |

---

# 11.4 Pencarian pada Struct

## Konsep

Pencarian tidak hanya pada integer/string.

Tetapi juga bisa pada:

```text
Struct
```

---

## Contoh Struct Mahasiswa

```go
type mahasiswa struct {
    nama, nim, kelas, jurusan string
    ipk float64
}
```

---

## Array Struct

```go
type arrMhs [2023]mahasiswa
```

---

# Sequential Search pada Struct

## Contoh: Cari Berdasarkan Nama

```go
func SeqSearchMhs(T arrMhs, n int, X string) int {

    var found int = -1
    var j int = 0

    for j < n && found == -1 {

        if T[j].nama == X {
            found = j
        }

        j++
    }

    return found
}
```

---

## Binary Search pada Struct

### Syarat Penting

```text
Field pencarian HARUS sesuai
dengan field pengurutan
```

---

## Contoh

```text
Array diurutkan berdasarkan nim
↓
Binary search harus mencari nim
```

---

## Contoh Binary Search Struct

```go
func BinarySearchMhs(T arrMhs, n int, X string) int {

    var found int = -1
    var med int

    var kr int = 0
    var kn int = n - 1

    for kr <= kn && found == -1 {

        med = (kr + kn) / 2

        if X < T[med].nim {

            kn = med - 1

        } else if X > T[med].nim {

            kr = med + 1

        } else {

            found = med
        }
    }

    return found
}
```

---

## Mengambil Data Lengkap

```go
idx := BinarySearchMhs(data, n, "23111001")

fmt.Println(data[idx].nama)
fmt.Println(data[idx].nim)
fmt.Println(data[idx].ipk)
```

---

# Perbandingan Sequential dan Binary Search

| Aspek        | Sequential Search | Binary Search  |
| ------------ | ----------------- | -------------- |
| Data Terurut | Tidak wajib       | Wajib          |
| Kecepatan    | Lebih lambat      | Lebih cepat    |
| Implementasi | Mudah             | Lebih kompleks |
| Cocok Untuk  | Data kecil        | Data besar     |

---

# Kesalahan Umum

## 1. Binary Search pada Data Tidak Terurut

```text
Hasil pencarian bisa salah
```

---

## 2. Salah Urutan Ascending/Descending

```text
Algoritma ascending
dipakai pada data descending
```

---

## 3. Tidak Menangani Data Tidak Ditemukan

Gunakan:

```go
return -1
```

---

# Kesimpulan

* Pencarian data digunakan untuk mencari elemen tertentu
* Sequential Search:

  * sederhana
  * tidak perlu data terurut
* Binary Search:

  * cepat
  * membutuhkan data terurut
* Pencarian dapat dilakukan pada:

  * array sederhana
  * struct
* Gunakan Binary Search untuk data besar yang sudah terurut

---
# Penutup

Pada modul ini mahasiswa telah mempelajari:

* Sequential Search
* Binary Search
* Pencarian pada struct
* Perbedaan pencarian linear dan pencarian berbasis keterurutan data

Materi ini menjadi dasar penting untuk:

* Sorting
* Database searching
* Algoritma optimasi
* Struktur data lanjutan

---
