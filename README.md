[![Go Report Card](https://goreportcard.com/badge/github.com/DevCriss/GoSmartSearch)](https://goreportcard.com/report/github.com/DevCriss/GoSmartSearch)
[![Godoc](https://godoc.org/github.com/DevCriss/GoSmartSearch?status.svg)](https://pkg.go.dev/github.com/DevCriss/GoSmartSearch)

# GoSmartSearch
## What is this?
GoSmartSearch is an open source library that gives you the ability to search within your data collections. Forget about annoying nested conditions when you want to find a specific element in your data collection (Slice) using the GO language.

Find out which items on your list are most similar to the entry, the match doesn't have to be exact.
>Example:
Input: geoge.
>Result: George, Georgi, Georgy, etc.


## How can i use it?

Just install this library using your terminal
> go get github.com/ChristianDC13/GoSmartSearch

Import the library in your amazing project
```
import (
 ...
  gss "github.com/ChristianDC13/GoSmartSearch"
)
```

Now, start your search
```
elements := []map[string]string{
	{"name": "George", "age": "21", "gender": "M"},
	{"name": "Maria", "age": "38", "gender": "F"},
	{"name": "Victor", "age": "19", "gender": "M"},
	{"name": "Elizabeth", "age": "27", "gender": "F"},
	{"name": "Alexander", "age": "41", "gender": "M"},
	{"name": "Pamela", "age": "32", "gender": "F"},
}
result, _ := gss.SearchInMaps(elements, "Axel Rose", "name", 0)
```

## Functions
> #### SearchInStrings:
>  Receive a slice of strings
>  Returns a slice formed by the input elements ordered from most to least similar to the input term

> ### SearchInMaps:
>  Receive a slice of maps (key: string, value: string)
>  Returns a map slice formed by the input elements ordered (based on the key) from most to least similar to the input term

## Params

SmartyPants converts ASCII punctuation characters into "smart" typographic punctuation HTML entities. For example:

| Param               |Type                          |Description                         |
|----------------|-------------------------------|-----------------------------|
|elements|`[]string`or `[]map[string]string`            |Is the list (slice) of elements where you hope to find the element you are looking for.           |
|term          |`string`            |Is the world  or term you want to find.            |
|tolerance          |`float32`|It is a number between 0 and 1 that represents how much an element of the source list must look like to the input to be considered valuable represent how much look like to the input to be considered valuable and be returned to the output list.|
|key          |`string`|(Only applicable to map search) It is the key of the map by which you want to do the search. Example: In a list of employees, you could search by 'firstName', 'lastName', etc |

## More info in the examples folder
> https://github.com/ChristianDC13/GoSmartSearch/tree/main/examples

## Author
- [Christian De La Cruz](https://christiandelacruz.dev) 

