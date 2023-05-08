package main

import (
	"fmt"
	"math"
)

func multiplication_table() {
	var number int
	fmt.Printf("Quelle table de multiplication souhaitez-vous afficher ? ")
	fmt.Scanln(&number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}

func pythagoras_theorem() {
	var a, b float64
	fmt.Printf("Indiquer les valeurs des côtés connus du triangle : ")
	fmt.Scanln(&a, &b)
	fmt.Printf("La longueur du troisième côté est : %.1f", math.Sqrt(a*a+b*b))
}

func remarkable_identity() {
	var a, b float64
	fmt.Println("Indiquer deux nombres dont vous voulez connaître leur identité remarquable : ")
	fmt.Scanf("%f %f", &a, &b)
	fmt.Printf("(%.1f + %.1f)² = %.1f\n", a, b, (a+b)*(a+b))
	fmt.Printf("(%.1f - %.1f)² = %.1f\n", a, b, (a-b)*(a-b))
	fmt.Printf("(%.1f + %.1f)(%.1f - %.1f) = %.1f\n", a, b, a, b, (a+b)*(a-b))
}

func main() {
	var choice int
	for {
		fmt.Println(" ____________________________________________")
		fmt.Println("|           MATHBYPASS - v.1.0.1             |")
		fmt.Println("|       (bIg br0th€r I$ w@tchIng y0u)        |")
		fmt.Println("|____________________________________________|")
		fmt.Println("")
		fmt.Println("        1 > Table de multiplication")
		fmt.Println("        2 > Théorème de Pythagore")
		fmt.Println("        3 > Identité remarquable")
		fmt.Println("")
		fmt.Printf("Entrez votre choix : ")
		fmt.Scanln(&choice)
		fmt.Println("")
		switch choice {
		case 1:
			multiplication_table()
		case 2:
			pythagoras_theorem()
		case 3:
			remarkable_identity()
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
		fmt.Println("")
		fmt.Println("Appuyez sur Entrée pour continuer...")
		fmt.Scanln()
	}
}
