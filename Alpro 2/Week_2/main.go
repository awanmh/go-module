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

    var r, t int

    fmt.Print("Masukkan jari-jari: ")
    fmt.Scanln(&r)

    fmt.Print("Masukkan tinggi: ")
    fmt.Scanln(&t)

    fmt.Printf("Volume Tabung: %.1f", volumeTabung(r,t))

}