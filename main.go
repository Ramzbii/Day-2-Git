package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

type person struct {
	name string
	age  int
}
type student struct {
	grade int
	person
}

func main() {
	//decimal
	var decimalNumber = 2.62
	fmt.Printf("Bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal : %.3f\n", decimalNumber)

	//boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	//string
	var message = " Nama saya Rama, salam kenal yak\n"
	fmt.Printf(message)

	//variable withouttype data
	nama := "Rama Pratama"
	fmt.Println("Nama saya : ", nama)

	// Variables Declaration
	var firstname string = "jon"
	var lastname string
	lastname = "wik"
	fmt.Printf("Halo %s %s!\n", firstname, lastname)

	/*Declaration Underscore Variables
	_ = "belajar Golang"
	_ = "golang itu mudah"
	name, _ := "jon", "wik"
	*/

	//constans
	const umur int = 20
	fmt.Print("Umurku ", umur, " tahun\n")

	//if else then

	var points = 8
	if points == 10 {
		fmt.Println("A")
	} else if points >= 7 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//switch case
	var point = 6
	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Wew")
	default:
		fmt.Println("WOW")
	}

	//for
	var fruits = [4]string{"apple", "Grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("angka", i)
	}
	//fungction
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	//struct
	var s1 = student{}
	s1.name = "ram"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("age :", s1.person.age)
	fmt.Println("grade :", s1.grade)

	var allstudents = []person{
		{name: "jamal", age: 20},
		{name: "ucup", age: 20},
		{name: "udin", age: 20},
	}

	for _, student := range allstudents {
		fmt.Println(student.name, "berumur", student.age)
	}

	//data structure map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	//slice
	var buahh = []string{"apple", "anggur", "pisang", "melon"}
	fmt.Println(buahh[0])

	//defer
	defer fmt.Println("ini defer pasti di akhir") // mengeksekusi di akhir
	fmt.Println("Selamat datang")

	//pointer
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
}
