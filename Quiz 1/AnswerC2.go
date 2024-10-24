package main
import "fmt"

func cekEsTebak(kartuNomor_2311102044 string) bool {
	digitPertama := kartuNomor_2311102044[0]
	for i := 0; i < len(kartuNomor_2311102044); i++ {
		digit := kartuNomor_2311102044[i]
		if digit%2 == 0 || digit != digitPertama {
			return false
		}
	}
	return true
}

func cekEsCendol(kartuNomor_2311102044 string) bool {
	for i := 0; i < len(kartuNomor_2311102044); i++ {
		if (kartuNomor_2311102044[i]-'0')%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	var N int
	var esTebakCount, esCendolCount, lamangCount int

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var kartuNomor_2311102044 string
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i+1)
		fmt.Scan(&kartuNomor_2311102044)

		if cekEsTebak(kartuNomor_2311102044) {
			fmt.Println("Es Tebak")
			esTebakCount++
		} else if cekEsCendol(kartuNomor_2311102044) {
			fmt.Println("Es Cendol")
			esCendolCount++
		} else {
			fmt.Println("Lamang")
			lamangCount++
		}
	}

	fmt.Printf("Jumlah yang memperoleh Es Tebak: %d\n", esTebakCount)
	fmt.Printf("Jumlah yang memperoleh Es Cendol: %d\n", esCendolCount)
	fmt.Printf("Jumlah yang memperoleh Lamang: %d\n", lamangCount)
}
