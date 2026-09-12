# MODUL 1: RUNNING MODUL & STRUKTUR PROGRAM GO

Modul ini membahas **pengenalan praktikum Algoritma & Pemrograman 1**, tata tertib laboratorium informatika, aturan penulisan program dan submisi tugas, serta struktur anatomi dasar bahasa Go (Golang).

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Memahami aturan, etika, dan mekanisme praktikum di Laboratorium Informatika (IFLAB)
* Memahami alur penyampaian keluhan dan aturan submisi tugas/jurnal
* Mempersiapkan lingkungan kerja pemrograman Go (instalasi kompiler & teks editor)
* Memahami struktur dasar dan anatomi kode program dalam bahasa Go
* Melakukan kompilasi dan menjalankan program Go secara mandiri melalui terminal

---

## 1.1 Pengantar Praktikum Algoritma & Pemrograman 1

Praktikum Algoritma dan Pemrograman 1 (kode mata kuliah CAK1BAB3) merupakan kegiatan laboratorium terstruktur yang mendampingi perkuliahan teori di kelas.

### Komponen Penugasan Praktikum

1. **Tugas Pendahuluan (TP)**: Diberikan sebelum praktikum dimulai untuk menguji kesiapan materi awal (estimasi pengerjaan ~60 menit).
2. **Tugas Jurnal**: Tugas utama praktikum selama 100 menit di laboratorium yang dikerjakan dengan bimbingan asisten praktikum.
3. **Asesmen Praktikum (Ujian Mandiri)**: Ujian coding praktikum tanpa bantuan asisten untuk mengukur kompetensi individual praktikan.

---

## 1.2 Peraturan Praktikum Laboratorium Informatika

Untuk menjaga kelancaran dan ketertiban praktikum di Gedung Telkom University Landmark Tower (TULT):

* **Kehadiran**: Praktikan wajib hadir minimal 75% dari total pertemuan.
* **Keterlambatan**:
  * $\le 5$ menit: Diperbolehkan mengikuti praktikum tanpa tambahan waktu.
  * $> 30$ menit: Tidak diperbolehkan mengikuti praktikum (nilai modul = 0).
* **Tata Tertib di Ruang Praktikum**:
  * Wajib berpakaian seragam resmi sesuai ketentuan institusi.
  * Mematikan/mengondisikan nada dering ponsel.
  * Dilarang membawa makanan dan minuman ke dalam lab.
  * Menjaga integritas akademik: dilarang bekerja sama, memberikan contekan, atau menyebarkan soal praktikum.

---

## 1.3 Aturan Penulisan Kode dan Submisi Tugas

Semua tugas pemrograman wajib mematuhi standar koding (*clean code*) laboratorium:

1. **Ekstensi File**: Setiap program sumber wajib disimpan dengan ekstensi `.go`.
2. **Indentasi**: Gunakan indentasi konsisten (4 spasi atau 1 tab per blok indentasi).
   * Gunakan perkakas bawaan:
   ```bash
   go fmt namafile.go
   ```
3. **Single Entry Single Exit**:
   * Alur eksekusi dimulai dari baris instruksi pertama dan berakhir pada satu titik keluar yang terdefinisi.
4. **Konvensi Penamaan (Naming Convention)**:
   * Variabel umum: Gunakan nama deskriptif berbahasa Inggris atau Indonesia yang konsisten (misal: `jumlah`, `rerata`, `totalHarga`).
   * Iterator / indeks: Gunakan `i`, `j`, `k`.
   * Konstanta simbolik: Gunakan huruf kapital penuh (misal: `PI`, `MAX_N`, `TARIF_DASAR`).
   * Hindari penamaan yang diawali dengan tanda garis bawah (`_nama`).

---

## 1.4 Anatomi dan Struktur Dasar Program Go

Bahasa Go dirancang sederhana, efisien, dan memiliki struktur program yang sangat jelas.

```text
┌────────────────────────────────────────┐
│             package main               │ ◄── Kop Surat / Penanda Program Utama
├────────────────────────────────────────┤
│             import "fmt"               │ ◄── Peminjaman Pustaka (Library)
├────────────────────────────────────────┤
│             func main() {              │
│                 // Perintah dieksekusi │ ◄── Badan Program Utama
│             }                          │
└────────────────────────────────────────┘
```

### 1. package main
Menandakan bahwa file ini adalah paket utama yang dapat dikompilasi menjadi berkas *executable* (dapat langsung dijalankan oleh sistem operasi).

### 2. import "fmt"
Mengimpor paket bawaan Go bernama `fmt` (*format*), yang menyediakan fungsi untuk operasi input (membaca ketikan keyboard) dan output (mencetak teks ke layar).

### 3. func main()
Fungsi utama program. Eksekusi kode Go selalu berawal dari baris pertama di dalam fungsi `main()` ini.

---

## 1.5 Komentar Program

Komentar adalah catatan dokumentasi yang **diabaikan oleh kompiler**, ditujukan bagi programmer:

```go
// Ini adalah komentar satu baris

/*
   Ini adalah komentar multibaris
   yang dapat memuat penjelasan panjang
*/
```

---

## 1.6 Contoh Program Lengkap

Berikut contoh program sederhana Go yang membaca dua angka dan mencetak hasil penjumlahannya:

```go
package main

import "fmt"

func main() {
    // Deklarasi variabel
    var nama string
    var a, b int

    fmt.Print("Masukkan nama Anda: ")
    fmt.Scanln(&nama)

    fmt.Println("Halo, selamat datang di Alpro 1,", nama)

    fmt.Print("Masukkan dua bilangan bulat: ")
    fmt.Scan(&a, &b)

    var hasil int = a + b
    fmt.Printf("Hasil penjumlahan %d + %d adalah %d\n", a, b, hasil)
}
```

---

## 1.7 Kompilasi dan Eksekusi Melalui Terminal

Bahasa Go merupakan bahasa bertipe *compiled language*. Terdapat dua perintah utama yang sering digunakan:

### 1. Menjalankan Langsung (Untuk Pengembangan / Uji Coba)
```bash
go run namafile.go
```
Perintah ini akan mengompilasi kode ke memori sementara dan langsung mengeksekusinya tanpa menghasilkan file `.exe` di folder kerja.

### 2. Mengompilasi Menjadi Berkas Biner (.exe)
```bash
go build namafile.go
```
Perintah ini menghasilkan berkas biner (misal: `namafile.exe` di Windows) yang dapat dijalankan langsung:
```bash
.\namafile.exe
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Latihan 1 – Menjalankan Program Pertama
1. Buat file bernama `halo.go` di folder kerja Anda.
2. Ketik program dasar yang mencetak identitas Anda (Nama, NIM, Kelas).
3. Jalankan dengan perintah `go run halo.go`.
4. Kompilasi dengan `go build halo.go`, amati terbentuknya file `.exe`, lalu jalankan file `.exe` tersebut.

### Latihan 2 – Eksplorasi Perintah `go fmt`
1. Tulis kode Go dengan indentasi berantakan (misal: spasi tidak teratur).
2. Jalankan perintah `go fmt halo.go`.
3. Perhatikan bagaimana Go secara otomatis merapikan tata letak baris kode Anda sesuai standar komunitas.

---

## Kesimpulan Modul 1

* Setiap program Go yang dapat dieksekusi mandiri wajib diawali dengan `package main` dan memiliki fungsi `func main()`.
* Library `fmt` adalah kunci utama untuk mencetak output ke terminal dan menerima masukan pengguna.
* Kebersihan kode (*code convention*) dan kedisiplinan praktikum adalah fondasi utama bagi calon praktikan sebelum mempelajari materi algoritma lanjutan.
