package basics //import "github.com/workspace7/tourofgo/basics"

import "fmt"

//BuildAndPrint the data built with struct Person
func BuildAndPrint() {
	sk := Person{"Kamesh", "Sampath", 38}
	fmt.Printf("First Name: %s , Last Name: %s , Age: %d\n", sk.firstName, sk.lastName, sk.age)

	sreek := Person{"Sreenidhi", "Kamesh", 39}
	fmt.Println(sreek)

	p2 := &sreek
	//normal pointer way cumbersome
	fmt.Printf("Ugly way:%d \n", (*p2).age)
	p2.age = 34
	// for struct alone golang allows direct deref with dot operator
	fmt.Println(p2.age)

	//Literals
	p3 := Person{firstName: "Rithul", lastName: "K", age: 10}
	fmt.Println(p3)

	//implicit values for uninit fields
	p4 := Person{firstName: "A"}
	fmt.Println(p4)

	// pointer to Struct
	p5 := &Person{"B", "C", 1}
	fmt.Println(*p5)
}
