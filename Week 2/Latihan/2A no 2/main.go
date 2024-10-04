package main
import "fmt"

func tahunKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} 
	if tahun%4 == 0 && tahun%100 != 0 {
		return true
	}
	return false
}

func main() {
	var tahun int

	for {
	    fmt.Print("Masukkan tahun (input 0 untuk keluar): ")
	    fmt.Scan(&tahun)

		if tahun == 0 {
			fmt.Println("Program selesai.")
			break
		}

	    if tahunKabisat(tahun) {
		    fmt.Println(tahun, ": tahun kabisat, TRUE")
	    } else {
		    fmt.Println(tahun, ": bukan tahun kabisat, FALSE")
	    }
    }
}