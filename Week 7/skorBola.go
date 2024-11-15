package main // 2311102044
import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Masukkan nama klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Masukkan skor pertandingan %s vs %s : ", klubA, klubB)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			hasil = append(hasil, klubA) 
		} else if skorA < skorB {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
	}

	fmt.Println("Pertandingan selesai!")
	for i, h := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, h)
	}
}
