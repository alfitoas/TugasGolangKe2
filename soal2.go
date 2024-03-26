package main

import (
	"fmt"
)

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func main() {
	// Data 1
	weightMark1 := 78.0
	heightMark1 := 1.69
	weightJohn1 := 92.0
	heightJohn1 := 1.95

	// Data 2
	weightMark2 := 95.0
	heightMark2 := 1.88
	weightJohn2 := 85.0
	heightJohn2 := 1.76

	// Menghitung BMI untuk Data 1
	bmiMark1 := calculateBMI(weightMark1, heightMark1)
	bmiJohn1 := calculateBMI(weightJohn1, heightJohn1)

	// Menghitung BMI untuk Data 2
	bmiMark2 := calculateBMI(weightMark2, heightMark2)
	bmiJohn2 := calculateBMI(weightJohn2, heightJohn2)

	// Menampilkan BMI
	fmt.Println("BMI untuk Data 1:")
	fmt.Printf("Mark: %.2f\n", bmiMark1)
	fmt.Printf("John: %.2f\n", bmiJohn1)

	fmt.Println("\nBMI untuk Data 2:")
	fmt.Printf("Mark: %.2f\n", bmiMark2)
	fmt.Printf("John: %.2f\n", bmiJohn2)

	// Membandingkan BMI Mark dan John
	markHigherBMI1 := bmiMark1 > bmiJohn1
	markHigherBMI2 := bmiMark2 > bmiJohn2

	// Menampilkan apakah Mark memiliki BMI lebih tinggi dari John
	fmt.Println("\nApakah Mark memiliki BMI lebih tinggi dari John pada Data 1?", markHigherBMI1)
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John pada Data 2?", markHigherBMI2)
}
