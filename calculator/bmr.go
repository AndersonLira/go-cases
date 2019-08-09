//Harris–Benedict equation
package main

import (
	"fmt"
)

func main() {
	conf := readValues()
	var calories float64
	var gender string
	if conf.man {
		calories = bmrMan(conf)
		gender = "MAN"
	} else {
		calories = bmrWoman(conf)
		gender = "WOMAN"
	}
	fmt.Println("")
	fmt.Printf("******************* %s *******************************\n", gender)
	fmt.Println("Weight:", conf.weight, "kg")
	fmt.Println("Eight:", conf.height, "cm")
	fmt.Println("Age:", conf.age, "years")
	fmt.Printf("Basal Metabolic Rate (level %s): %d kilocalorie\n", conf.level.name(), int(calories*conf.level.value()))
}

//weight in kg, height in cm and age in years
func bmrMan(conf config) float64 {
	result := 66.47 + (13.7 * float64(conf.weight)) + (5 * float64(conf.height)) - (6.8 * float64(conf.age))
	return result
}

func bmrWoman(conf config) float64 {
	result := 655.1 + (9.563 * float64(conf.weight)) + (1.853 * float64(conf.height)) - (4.676 * float64(conf.age))
	return result
}

type config struct {
	weight int
	height int
	age    int
	man    bool
	level  level
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
	var b bool
	fmt.Println("Are you man? t == true")
	fmt.Scanln(&b)
	conf.man = b
	for conf.level < 1 || conf.level > 5 {
		var i int
		fmt.Println("Level activity")
		fmt.Println("1 sedentary")
		fmt.Println("2 slightly")
		fmt.Println("3 moderate")
		fmt.Println("4 very")
		fmt.Println("5 extreme")
		fmt.Scanln(&i)
		conf.level = level(i)
	}

	return conf
}

type level int

const (
	sedentary level = 1 + iota
	slightly
	moderate
	very
	extreme
)

func (l level) value() float64 {
	switch l {
	case sedentary:
		return 1.2
	case slightly:
		return 1.375
	case moderate:
		return 1.55
	case very:
		return 1.725
	case extreme:
		return 1.9

	}
	panic("Invalid level value")
}

func (l level) name() string {
	switch l {
	case sedentary:
		return "sedentary"
	case slightly:
		return "slightly"
	case moderate:
		return "moderate"
	case very:
		return "very intense"
	case extreme:
		return "extreme"

	}
	panic("Invalid level value")

}
