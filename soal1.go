package main

import "fmt"

func hitungRataRata(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

func main() {
	// Data
	lumbaLumba := []int{96, 108, 89}
	koala := []int{88, 91, 110}

	// Data Bonus 1
	lumbaLumbaBonus1 := []int{97, 112, 101}
	koalaBonus1 := []int{109, 95, 123}

	// Data Bonus 2
	lumbaLumbaBonus2 := []int{97, 112, 101}
	koalaBonus2 := []int{109, 95, 106}

	// Menghitung skor rata-rata untuk setiap tim
	rataRataLumbaLumba := hitungRataRata(lumbaLumba)
	rataRataKoala := hitungRataRata(koala)

	// Menghitung skor rata-rata untuk setiap tim dengan Bonus 1
	rataRataLumbaLumbaBonus1 := hitungRataRata(lumbaLumbaBonus1)
	rataRataKoalaBonus1 := hitungRataRata(koalaBonus1)

	// Menghitung skor rata-rata untuk setiap tim dengan Bonus 2
	rataRataLumbaLumbaBonus2 := hitungRataRata(lumbaLumbaBonus2)
	rataRataKoalaBonus2 := hitungRataRata(koalaBonus2)

	// Membandingkan skor rata-rata untuk menentukan pemenang tanpa bonus
	fmt.Println("Skor rata-rata tanpa bonus :")
	fmt.Printf("Tim Lumba-lumba : %.2f\n", rataRataLumbaLumba)
	fmt.Printf("Tim Koala : %.2f\n", rataRataKoala)
	if rataRataLumbaLumba > rataRataKoala {
		fmt.Println("Pemenang: Tim Lumba-lumba")
	} else if rataRataKoala > rataRataLumbaLumba {
		fmt.Println("Pemenang: Tim Koala")
	} else {
		fmt.Println("Hasil Seri!")
	}

	// Membandingkan skor rata-rata untuk menentukan pemenang dengan bonus 1
	fmt.Println("\nSkor rata-rata dengan bonus 1:")
	fmt.Printf("Tim Lumba-lumba : %.2f\n", rataRataLumbaLumbaBonus1)
	fmt.Printf("Tim Koala : %.2f\n", rataRataKoalaBonus1)
	if rataRataLumbaLumbaBonus1 >= 100 && rataRataKoalaBonus1 >= 100 {
		if rataRataLumbaLumbaBonus1 > rataRataKoalaBonus1 {
			fmt.Println("Pemenang: Tim Lumba-lumba")
		} else if rataRataKoalaBonus1 > rataRataLumbaLumbaBonus1 {
			fmt.Println("Pemenang: Tim Koala")
		} else {
			fmt.Println("Hasil Seri!")
		}
	} else {
		fmt.Println("Tidak ada pemenang karena skor rata-rata di bawah 100.")
	}

	// Membandingkan skor rata-rata untuk menentukan pemenang dengan bonus 2
	fmt.Println("\nSkor rata-rata dengan bonus 2:")
	fmt.Printf("Tim Lumba-lumba : %.2f\n", rataRataLumbaLumbaBonus2)
	fmt.Printf("Tim Koala : %.2f\n", rataRataKoalaBonus2)
	if rataRataLumbaLumbaBonus2 >= 100 && rataRataKoalaBonus2 >= 100 {
		if rataRataLumbaLumbaBonus2 > rataRataKoalaBonus2 {
			fmt.Println("Pemenang: Tim Lumba-lumba")
		} else if rataRataKoalaBonus2 > rataRataLumbaLumbaBonus2 {
			fmt.Println("Pemenang: Tim Koala")
		} else {
			fmt.Println("Hasil Seri!")
		}
	} else {
		fmt.Println("Tidak ada pemenang karena skor rata-rata di bawah 100.")
	}
}
