//Harris–Benedict equation
package main

import (
	"fmt"
)

const sedentary = 1.2
const slightly = 1.375
const moderate = 1.55
const very = 1.725
const extreme = 1.9

func main() {
	conf := readValues()
	c := bmrMan(conf.weight, conf.height, conf.age)
	fmt.Println("")
	fmt.Println("******************* MEN *******************************")
	fmt.Println("Weight:", conf.weight, "kg")
	fmt.Println("Eight:", conf.height, "cm")
	fmt.Println("Age:", conf.age, "years")
	fmt.Printf("Basal Metabolic Rate: %d kilocalorie\n", int(c*sedentary))
	fmt.Println(bmrWoman(65, 150, 30) * sedentary)
}

//weight in kg, height in cm and age in years
func bmrMan(weight, height, age int) float64 {
	result := 66.47 + (13.7 * float64(weight)) + (5 * float64(height)) - (6.8 * float64(age))
	return result
}

func bmrWoman(weight, height, age int) float64 {
	result := 655.1 + (9.563 * float64(weight)) + (1.853 * float64(height)) - (4.676 * float64(age))
	return result
}

type config struct {
	weight int
	height int
	age    int
	man    bool
	level  int
}

func readValues() config {
	conf := config{}
	for conf.weight <= 0 {
		var i int
		fmt.Println("Type weigth > 0 in kg")
		fmt.Scanln(&i)
		conf.weight = i
	}
	for conf.height <= 0 {
		var i int
		fmt.Println("Type heigth > 0 in cm")
		fmt.Scanln(&i)
		conf.height = i
	}
	for conf.age <= 0 {
		var i int
		fmt.Println("Type age > 0 in years")
		fmt.Scanln(&i)
		conf.age = i
	}
	// var b bool
	// fmt.Println("Are you man? t == true")
	// fmt.Scanln(&b)
	// conf.man = b
	// for conf.level <= 0 || conf.level >= 4 {
	// 	var i int
	// 	fmt.Println("Level activity")
	// 	fmt.Println("1 sedentary")
	// 	fmt.Println("2 moderate")
	// 	fmt.Println("3 vigorously")
	// 	fmt.Scanln(&i)
	// 	conf.level = i
	// }

	return conf
}
