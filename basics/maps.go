package basics //import "github.com/workspace7/tourofgo/basics"

import "fmt"

//MapAssign - map assignment

func MapAssign() {
	var myFamily map[string]Person = make(map[string]Person)

	myFamily["me"] = Person{"Kamesh", "Sampath", 38}

	fmt.Println("My Family : %v", myFamily)
}

func MapLiterals() map[string]Person {
	var myFamily map[string]Person = map[string]Person{
		"me":   Person{"Kamesh", "Sampath", 38},
		"sree": {"Sreenidhi", "Kamesh", 34},
	}

	myFamily["keshav"] = Person{"Rithul", "K", 10}

	fmt.Printf("My Family : %v", myFamily)

	myFamily["dummy"] = Person{}

	return myFamily
}

// func main() {
// 	//map_assign()
// 	myFamily := map_literals()
// 	elem, ok := myFamily["keshav"]
// 	if ok {
// 		fmt.Printf("\n Keshav's First Name: %v", elem.firstName)
// 	}
// 	_, avbl := myFamily["dummy"]

// 	if avbl {
// 		fmt.Println("\nDeleting key \"dummy\"")
// 		delete(myFamily, "dummy")
// 	}

// 	el, avbl := myFamily["dummy"]

// 	if !avbl {
// 		fmt.Println("Succesfully deleted \"dummy\"")
// 	}
// }
