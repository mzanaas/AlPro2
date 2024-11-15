package main // 2311102044
import "fmt"

type Titik struct {
    x, y float64
}

type Lingkaran struct {
    pusat Titik
    radius float64
}

func dalamLingkaran(l Lingkaran, t Titik) bool {
    jarakKuadrat := (t.x-l.pusat.x)*(t.x-l.pusat.x) + (t.y-l.pusat.y)*(t.y-l.pusat.y)
    return jarakKuadrat <= l.radius*l.radius
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

    fmt.Scan(&cx1, &cy1, &r1)
    fmt.Scan(&cx2, &cy2, &r2)
    fmt.Scan(&x, &y)

    lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{x, y}
    diLingkaran1 := dalamLingkaran(lingkaran1, titik)
	diLingkaran2 := dalamLingkaran(lingkaran2, titik)

    if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik berada di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}
}