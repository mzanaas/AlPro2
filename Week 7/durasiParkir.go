package main
import "fmt"

// Struct untuk menyimpan
type waktu struct {
	jam, menit, detik int
}

func main () {
	// Deklarasi variabel
	var wParkir, wPulang, durasi waktu
    var dParkir, dPulang, lParkir int

	// Input
	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

	// Konversi waktu
	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600

    // Hitung lama parkir
	lParkir = dPulang - dParkir

	// Konversi durasi parkir
	durasi.jam = lParkir / 3600
	durasi.menit = (lParkir % 3600) / 60
	durasi.detik = lParkir % 60

	// Tampilkan hasil durasi parkir
	fmt.Printf("Lama parkir: %d jam %d menit %d detik\n", durasi.jam, durasi.menit, durasi.detik)
}