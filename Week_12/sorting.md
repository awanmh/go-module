# MODUL 12 & 13: PENGURUTAN DATA

### Tujuan Pembelajaran

Mahasiswa mampu:

* Memahami konsep dasar sorting (pengurutan data)
* Memahami algoritma Selection Sort
* Memahami algoritma Insertion Sort
* Mengimplementasikan sorting pada array sederhana
* Mengimplementasikan sorting pada struct
* Memahami perbedaan Selection Sort dan Insertion Sort
* Menggunakan sorting untuk menyelesaikan studi kasus

---

## 12.1 Konsep Dasar Pengurutan Data

### Apa itu Sorting?

Sorting adalah:

> Proses mengurutkan data berdasarkan aturan tertentu

---

### Jenis Pengurutan

| Jenis      | Contoh    |
| ---------- | --------- |
| Ascending  | 1 2 3 4 5 |
| Descending | 5 4 3 2 1 |

---

### Kenapa Sorting Penting?

Karena sorting digunakan hampir di semua sistem:

* Database
* Pencarian data
* Ranking
* Statistik
* Machine Learning
* Sistem rekomendasi

---

### Contoh Kehidupan Nyata

```text id="w42h3p"
- Mengurutkan nilai mahasiswa
- Mengurutkan harga produk
- Mengurutkan ranking game
- Mengurutkan buku perpustakaan
```

---

## 12.2 Ide Algoritma Selection Sort

### Konsep Selection Sort

Selection Sort adalah:

> Algoritma pengurutan dengan mencari nilai terkecil/terbesar lalu menempatkannya pada posisi yang benar

---

### Ide Dasar

Untuk ascending:

```text id="f20w7v"
1. Cari nilai terkecil
2. Tukar dengan posisi paling kiri
3. Ulangi untuk sisa data
```

---

### Analogi

```text id="1j7xpk"
Menyusun kartu angka:

[7 2 9 1 5]

Cari terkecil → 1
Tukar ke depan

[1 2 9 7 5]

Cari terkecil berikutnya
```

---

### Diagram Selection Sort

```text id="3t7jgo"
        ┌──────────────────┐
        │ Cari nilai min   │
        │ pada sisa data   │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Tukar dengan     │
        │ posisi depan     │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Geser area       │
        │ pencarian        │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Ulangi sampai    │
        │ selesai          │
        └──────────────────┘
```

---

### Pseudocode Selection Sort

```text id="vrf8j2"
i = 1

selama i <= n-1:

    idx_min = i-1
    j = i

    selama j < n:

        jika data[idx_min] > data[j]:
            idx_min = j

        j++

    swap(data[idx_min], data[i-1])

    i++
```

---

### Versi Go

```go id="bjyz53"
i := 1

for i <= n-1 {

    idx_min := i - 1
    j := i

    for j < n {

        if a[idx_min] > a[j] {
            idx_min = j
        }

        j++
    }

    t := a[idx_min]
    a[idx_min] = a[i-1]
    a[i-1] = t

    i++
}
```

---

## 12.3 Algoritma Selection Sort

### Contoh pada Array Integer

```go id="8tkv27"
package main
import "fmt"

type arrInt [4321]int

func selectionSort(T *arrInt, n int) {

    var t, i, j, idx_min int

    i = 1

    for i <= n-1 {

        idx_min = i - 1
        j = i

        for j < n {

            if T[idx_min] > T[j] {
                idx_min = j
            }

            j++
        }

        t = T[idx_min]
        T[idx_min] = T[i-1]
        T[i-1] = t

        i++
    }
}

func main() {

    data := arrInt{7, 2, 9, 1, 5}

    selectionSort(&data, 5)

    fmt.Println(data)
}
```

---

### Cara Kerja Selection Sort

#### Iterasi 1

```text id="yl3d43"
[7 2 9 1 5]

Cari terkecil → 1

Tukar dengan depan

[1 2 9 7 5]
```

---

#### Iterasi 2

```text id="1oqk7f"
[1 2 9 7 5]

Cari terkecil dari:
[2 9 7 5]

→ 2
```

---

#### Iterasi 3

```text id="zmxr4q"
[1 2 9 7 5]

Cari terkecil:
→ 5

[1 2 5 7 9]
```

---

### Insight Penting

* Setiap iterasi menempatkan satu data pada posisi benar
* Menggunakan proses:

  * pencarian minimum
  * swap

---

## 12.4 Selection Sort pada Struct

### Konsep

Sorting juga dapat dilakukan pada:

```text id="snr9t4"
Struct
```

---

### Contoh Struct Mahasiswa

```go id="87ynbt"
type mahasiswa struct {
    nama, nim, kelas, jurusan string
    ipk float64
}
```

---

### Array Struct

```go id="b4z93k"
type arrMhs [2023]mahasiswa
```

---

### Selection Sort Berdasarkan IPK

```go id="91g7t8"
func selectionSortMhs(T *arrMhs, n int) {

    var i, j, idx_min int
    var t mahasiswa

    i = 1

    for i <= n-1 {

        idx_min = i - 1
        j = i

        for j < n {

            if T[idx_min].ipk > T[j].ipk {
                idx_min = j
            }

            j++
        }

        t = T[idx_min]
        T[idx_min] = T[i-1]
        T[i-1] = t

        i++
    }
}
```

---

### Insight Penting

Pada struct:

```text id="7sztq4"
Yang dibandingkan adalah field tertentu
```

Contoh:

* ipk
* nama
* nim
* tahun

---

## 12.6 Ide Algoritma Insertion Sort

### Konsep Insertion Sort

Insertion Sort adalah:

> Algoritma pengurutan dengan menyisipkan data ke posisi yang sesuai

---

### Ide Dasar

```text id="34k7qi"
1. Ambil satu data
2. Cari posisi yang sesuai
3. Geser data lain
4. Sisipkan data
```

---

### Analogi

```text id="q9z1oc"
Menyusun kartu di tangan:

Ambil satu kartu
Cari posisi yang tepat
Geser kartu lain
Masukkan kartu
```

---

### Diagram Insertion Sort

```text id="3f3l0h"
        ┌──────────────────┐
        │ Ambil satu data  │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Cari posisi      │
        │ yang sesuai      │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Geser data       │
        │ sebelumnya       │
        └────────┬─────────┘
                 ▼
        ┌──────────────────┐
        │ Sisipkan data    │
        └──────────────────┘
```

---

### Pseudocode Insertion Sort

```text id="rzqz7m"
i = 1

selama i <= n-1:

    j = i
    temp = data[j]

    selama j > 0 dan temp > data[j-1]:

        data[j] = data[j-1]
        j--

    data[j] = temp

    i++
```

---

### Versi Go

```go id="0k1o5t"
i := 1

for i <= n-1 {

    j := i
    temp := a[j]

    for j > 0 && temp > a[j-1] {

        a[j] = a[j-1]
        j--
    }

    a[j] = temp

    i++
}
```

---

## 12.7 Algoritma Insertion Sort

### Contoh pada Array Integer

```go id="cfd4wh"
package main
import "fmt"

type arrInt [4321]int

func insertionSort(T *arrInt, n int) {

    var temp, i, j int

    i = 1

    for i <= n-1 {

        j = i
        temp = T[j]

        for j > 0 && temp > T[j-1] {

            T[j] = T[j-1]
            j--
        }

        T[j] = temp

        i++
    }
}

func main() {

    data := arrInt{7, 2, 9, 1, 5}

    insertionSort(&data, 5)

    fmt.Println(data)
}
```

---

### Cara Kerja Insertion Sort

#### Langkah 1

```text id="4h8hhy"
[7 2 9 1 5]

Ambil 2

Bandingkan dengan 7
```

---

#### Langkah 2

```text id="djlwm0"
Geser 7

[7 7 9 1 5]

Masukkan 2

[2 7 9 1 5]
```

---

#### Langkah 3

```text id="tn6jpk"
Ambil 9

Masuk ke posisi benar

[2 7 9 1 5]
```

---

## 12.8 Insertion Sort pada Struct

### Contoh

```go id="0l5j80"
func insertionSortMhs(T *arrMhs, n int) {

    var i, j int
    var temp mahasiswa

    i = 1

    for i <= n-1 {

        j = i
        temp = T[j]

        for j > 0 && temp.nama > T[j-1].nama {

            T[j] = T[j-1]
            j--
        }

        T[j] = temp

        i++
    }
}
```

---

### Insight Penting

Pada insertion sort:

```text id="jlwm6h"
Tidak ada swap
```

Yang dilakukan adalah:

```text id="r5q3jt"
Menggeser data
```

---

## 12.9 Perbandingan Selection Sort vs Insertion Sort

| Aspek          | Selection Sort | Insertion Sort                |
| -------------- | -------------- | ----------------------------- |
| Konsep         | Cari minimum   | Sisipkan data                 |
| Operasi utama  | Swap           | Shift                         |
| Mudah dipahami | Ya             | Ya                            |
| Cocok untuk    | Data kecil     | Data hampir terurut           |
| Performa       | Stabil standar | Lebih baik pada data tertentu |

---

## Kesimpulan

* Sorting digunakan untuk menyusun data
* Selection Sort:

  * mencari nilai minimum/maksimum
  * menggunakan swap
* Insertion Sort:

  * menyisipkan data
  * menggunakan shift
* Sorting dapat diterapkan pada:

  * integer
  * string
  * struct
* Sorting sangat penting untuk optimasi pencarian data

---
