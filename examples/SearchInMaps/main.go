package main

import (
	"fmt"
	"time"

	gss "github.com/DevCriss/GoSmartSearch"
)

func main() {
	initTime := time.Now()
	elements := []map[string]string{
		{"name": "George", "age": "21", "gender": "M"},
		{"name": "Maria", "age": "38", "gender": "F"},
		{"name": "Victor", "age": "19", "gender": "M"},
		{"name": "Elizabeth", "age": "27", "gender": "F"},
		{"name": "Alexander", "age": "41", "gender": "M"},
		{"name": "Pamela", "age": "32", "gender": "F"},
	}
	result, _ := gss.SearchInMaps(elements, "Axel Rose", "name", 0)

	fmt.Println("Results:")
	for i, value := range result {
		fmt.Printf("#%d: %s\n", i, value)
	}
	fmt.Println("Time elapsed:", time.Since(initTime))
	fmt.Println("Slice size: ", len(elements))

}
