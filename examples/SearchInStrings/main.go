package main

import (
	"fmt"
	"time"

	gss "github.com/DevCriss/GoSmartSearch"
)

func main() {
	initTime := time.Now()
	elements := []string{
		"juan",
		"pepe",
		"victor",
		"helen",
	}

	result, err := gss.SearchInStrings(elements, "Axel Rose", 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Results:")
	for i, value := range result {
		fmt.Printf("#%d: %s\n", i, value)
	}
	fmt.Println("Time elapsed:", time.Since(initTime))
	fmt.Println("Slice size: ", len(elements))
}
