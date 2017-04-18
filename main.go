package main

func main() {

	pName := "arrays_slice" // will be repalced with commands
	switch pName {
	case "arrays_slice":
		{
			ChantWithArray()
			chants := make([]string, 2, 10)
			chants[0] = "Hare Krishna"
			ChantWithSlice(chants)
		}
	}
}
