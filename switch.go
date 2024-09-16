package main

import "fmt"

func main() {
	fmt.Println("Masukkan suhu awal (dalam derajat Celcius):")
	var suhuAwal float64
	var pilihan int
	fmt.Scanln(&suhuAwal)
	fmt.Println("Suhu awal:", suhuAwal, "°C")

	fmt.Println()

	fmt.Println("Pilih konversi suhu:")
	fmt.Println("1. Reamur")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Scanln(&pilihan)

	fmt.Println()

	switch pilihan {
	case 1:
		fmt.Println("Konversi ke Reamur :", suhuAwal*0.8, "°R")
	case 2:
		fmt.Println("Konversi ke Fahrenheit :", suhuAwal*1.8+32, "°F") 
	case 3:
		fmt.Println("Konversi ke Kelvin :", suhuAwal+273.15, " K")
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
