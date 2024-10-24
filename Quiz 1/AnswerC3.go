package main
import "fmt"

func perkalianRekursif(n int, m int) int {
    if m == 0 {
    return 0
    }
    return n + perkalianRekursif(n, m-1)
}

func main() {
    var bilangan_2311102044, bilanganM int
    fmt.Print("Masukkan bilangan n: ")
    fmt.Scan(&bilangan_2311102044)

    fmt.Print("Masukkan bilangan m: ")
    fmt.Scan(&bilanganM)
	
	hasil := perkalianRekursif(bilangan_2311102044, bilanganM) 
	fmt.Printf("Hasil dari %d x %d = %d\n", bilangan_2311102044, bilanganM, hasil)
}