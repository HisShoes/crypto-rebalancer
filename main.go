package main

func main() {
	dbType := "memory"

	switch dbType {
	case "memory":

	default:
		panic("Unknown database")
	}
}
