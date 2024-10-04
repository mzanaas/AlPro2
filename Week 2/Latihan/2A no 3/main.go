package main 
import "fmt"

const phi = 3.1415926535

func main() {
	var r int

	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * phi * float64(r*r*r)
	luas := 4.0 * phi * float64(r*r)

	fmt.Printf("Volume bola: %.2f\n", volume)
	fmt.Printf("Luas permukaan bola: %.2f\n", luas)

}