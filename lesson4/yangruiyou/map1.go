package main

import "fmt"

func main() {
	populationForCity := map[string]int{
		"Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}

	for city, population := range populationForCity {
		fmt.Printf("%-10s %8d \n", city, population)
	}

}
