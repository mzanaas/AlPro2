package main
import "fmt"

func main() {
	var percobaan [5][4]string
	var ulangi string

	for {
		sama := true
		for i := 0; i < 5; i++ {
			fmt.Printf("Percobaan ke-%d: ", i+1)

			for j := 0; j < 4; j++ {
				fmt.Scan(&percobaan[i][j])
			}
		}

		for i := 1; i < 5; i++ {
			for j := 0; j < 4; j++ {
				if percobaan[i][j] != percobaan[0][j] {
					sama = false
					break
				}
			}
			if !sama {
				break
			}
		}

		if sama {
			fmt.Println("TRUE, semua urutan warna sama.")
		} else {
			fmt.Println("FALSE, ada perbedaan dalam urutan warna")
		}

		fmt.Print("Apakah ingin melakukan percobaan lagi? (y/n): ")
		fmt.Scan(&ulangi)

		if ulangi != "y" {
			fmt.Println("Program selesai.")
			break
		}
	}
}
