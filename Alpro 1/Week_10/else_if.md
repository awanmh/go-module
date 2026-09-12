# MODUL 10: STRUKTUR KONTROL - ELSE-IF (PERCABANGAN GANDA & MAJEMUK)

Modul ini membahas **percabangan dua arah (`if-else`)** dan **percabangan banyak arah bertingkat (`else-if`)**, pemetaan rentang nilai (*range checking*), dan penyelesaian studi kasus tarif logistik bertingkat.

### Tujuan Pembelajaran

Setelah mempelajari modul ini, mahasiswa mampu:

* Mengimplementasikan percabangan dua kondisi (`if-else`)
* Mengimplementasikan percabangan majemuk bertingkat (`if-else if-else`)
* Menentukan urutan evaluasi kondisi logika yang efektif dan bebas tumpang tindih (*mutually exclusive*)
* Memecahkan permasalahan klasifikasi mutu, penentuan kategori usia, dan perhitungan tarif berjenjang

---

## 10.1 Percabangan Dua Kondisi (`if-else`)

Ketika ada **dua alternatif tindakan yang saling bertolak belakang**:
> Jika kondisi terpenuhi lakukan **Aksi 1**, jika tidak terpenuhi lakukan **Aksi 2**.

```text
                  /─────────────────────\
                 <  Apakah Kondisi Benar >
                  \─────────────────────/
                     │               │
               Ya    │               │ Tidak
                     ▼               ▼
           ┌──────────────────┐    ┌──────────────────┐
           │     Aksi 1       │    │     Aksi 2       │
           └─────────┬────────┘    └─────────┬────────┘
                     │                       │
                     └───────────┬───────────┘
                                 ▼
                    ┌─────────────────────────┐
                    │  Instruksi Selanjutnya  │
                    └─────────────────────────┘
```

### Sintaks Go:
```go
if kondisi {
    // Blok dieksekusi jika kondisi bernilai true
} else {
    // Blok dieksekusi jika kondisi bernilai false
}
```

> [!IMPORTANT]
> Di Go, kata kunci `else` **harus berada pada baris yang sama** dengan tanda kurung kurawal tutup `}` dari blok `if` sebelumnya.

---

## 10.2 Percabangan Majemuk (`if-else if-else`)

Digunakan ketika terdapat **tiga atau lebih kondisi** yang saling eksklusif:

```go
if kondisi_1 {
    // Aksi 1
} else if kondisi_2 {
    // Aksi 2
} else if kondisi_3 {
    // Aksi 3
} else {
    // Aksi default jika tidak ada kondisi yang terpenuhi
}
```

---

## 10.3 Contoh Program Lengkap

### Program 1 – Cek Kelayakan Pembuatan KTP
Seseorang dapat membuat KTP jika usianya sudah $\ge 17$ tahun dan memiliki berkas lengkap (status kewarganegaraan):

```go
package main

import "fmt"

func main() {
    var usia int
    var berkasLengkap bool

    fmt.Print("Masukkan usia dan status kelengkapan berkas (true/false): ")
    fmt.Scan(&usia, &berkasLengkap)

    if usia >= 17 && berkasLengkap {
        fmt.Println("Bisa membuat KTP")
    } else {
        fmt.Println("Belum bisa membuat KTP")
    }
}
```

---

### Program 2 – Konversi Nilai Angka ke Indeks Huruf
Program mengklasifikasikan nilai ujian mahasiswa ke dalam indeks huruf:

```go
package main

import "fmt"

func main() {
    var nilai float64

    fmt.Print("Masukkan nilai akhir: ")
    fmt.Scan(&nilai)

    var indeks string
    if nilai >= 80.0 {
        indeks = "A"
    } else if nilai >= 70.0 {
        indeks = "AB"
    } else if nilai >= 65.0 {
        indeks = "B"
    } else if nilai >= 60.0 {
        indeks = "BC"
    } else if nilai >= 50.0 {
        indeks = "C"
    } else if nilai >= 40.0 {
        indeks = "D"
    } else {
        indeks = "E"
    }

    fmt.Printf("Nilai: %.2f | Indeks Huruf: %s\n", nilai, indeks)
}
```

---

# TUGAS & LATIHAN PRAKTIK

---

### Soal 1 – Aplikasi Perhitungan Biaya Kirim Pos (PT POS)
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel.
**Ketentuan Perhitungan:**
1. Masukan berupa total berat parsel dalam satuan **gram**.
2. Hitung total berat dalam **kg** dan **sisa gram**.
3. Biaya dasar jasa pengiriman adalah **Rp 10.000,- per kg**.
4. Biaya sisa berat:
   * Jika sisa berat $\ge 500$ gram, biaya tambahan adalah **Rp 5,- per gram**.
   * Jika sisa berat $< 500$ gram, biaya tambahan adalah **Rp 15,- per gram**.
5. **Kebijakan Diskon Khusus**: Jika total berat parsel **lebih dari 10 kg** ($> 10000$ gram), maka seluruh biaya dari sisa berat (yang kurang dari 1 kg) **digratiskan (Rp 0)**.

#### Contoh Interaksi Program:
```text
Contoh 1:
Berat parsel (gram): 8500
Detail berat: 8 kg + 500 gr
Detail biaya: Rp. 80000 + Rp. 2500
Total biaya: Rp. 82500

Contoh 2:
Berat parsel (gram): 9250
Detail berat: 9 kg + 250 gr
Detail biaya: Rp. 90000 + Rp. 3750
Total biaya: Rp. 93750

Contoh 3:
Berat parsel (gram): 11750
Detail berat: 11 kg + 750 gr
Detail biaya: Rp. 110000 + Rp. 0
Total biaya: Rp. 110000
```

---

### Soal 2 – Klasifikasi Bentuk Segitiga
Diberikan 3 panjang sisi segitiga $a, b, c$. Periksa terlebih dahulu apakah ketiga sisi tersebut membentuk segitiga yang valid (jumlah dua sisi selalu lebih besar dari sisi ketiga). Jika valid, tentukan jenisnya:
* **Segitiga Sama Sisi**: jika ketiga sisi sama panjang ($a == b$ dan $b == c$).
* **Segitiga Sama Kaki**: jika tepat dua sisi sama panjang.
* **Segitiga Sembarang**: jika ketiga sisi berbeda panjang.
* Jika tidak memenuhi syarat segitiga: cetak `"Bukan Segitiga"`.

---

## Kesimpulan Modul 10

* Struktur `if-else` dan `else-if` memungkinkan program menangani banyak skenario alternatif secara terstruktur.
* Urutan pengecekan kondisi dalam `else-if` harus disusun secara logis dari kondisi paling spesifik ke kondisi yang lebih umum.
