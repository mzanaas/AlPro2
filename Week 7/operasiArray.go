package main // 2311102044
import "fmt"

func hitungRataRata(arr []int) float64 {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return float64(total) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int, rataRata float64) float64 {
	var total float64 = 0
	for i := 0; i < len(arr); i++ {
		deviation := float64(arr[i]) - rataRata
		total += deviation * deviation
	}
	varian := total / float64(len(arr))
	standarDeviasi := varian
	for i := 0; i < 10; i++ { // Menghitung akar kuadrat secara manual
		standarDeviasi = (standarDeviasi + varian/standarDeviasi) / 2
	}
	return standarDeviasi
}

func hitungFrekuensi(arr []int, bilangan int) int {
	frekuensi := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == bilangan {
			frekuensi++
		}
	}
	return frekuensi
}

func hapusElemen(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}

func main() {
	var N, x, hapusIndeks, cariBilangan int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. tampil seluruh isi array
	fmt.Println("a. Seluruh isi array:", arr)

	// b. elemen indeks ganjil
	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. elemen indeks genap
	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. elemen indeks kelipatan bilangan x
	fmt.Print("Masukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// e. hapus elemen tertentu
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&hapusIndeks)
	arr = hapusElemen(arr, hapusIndeks)
	fmt.Println("e. Isi array setelah penghapusan:", arr)

	// f. hitung rata-rata
	rataRata := hitungRataRata(arr)
	fmt.Printf("f. Rata-rata elemen array: %.2f\n", rataRata)

	// g. hitung standar deviasi
	standarDeviasi := hitungStandarDeviasi(arr, rataRata)
	fmt.Printf("g. Standar deviasi elemen array: %.2f\n", standarDeviasi)

	// h. hitung frekuensi bilangan
	fmt.Print("Masukkan bilangan untuk mencari frekuensinya: ")
	fmt.Scan(&cariBilangan)
	frekuensi := hitungFrekuensi(arr, cariBilangan)
	fmt.Printf("h. Frekuensi bilangan %d dalam array: %d\n", cariBilangan, frekuensi)
}
