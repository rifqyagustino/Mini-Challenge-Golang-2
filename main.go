package main

import( 
	"fmt"
	"strings"
)

func main(){
	kalimat := "selamat malam"

	// Pecah kalimat menjadi huruf-huruf
	huruf := strings.Split(kalimat, "")

	// Buat map untuk menghitung kemunculan huruf-huruf dan spasi
	kemunculanHuruf := make(map[string]int)

	// Looping melalui setiap huruf
	for _, h := range huruf{
		fmt.Println(h)
		// Hitung kemunculan huruf dalam map
		kemunculanHuruf[h]++		
	}

	fmt.Println(kemunculanHuruf)
}