package GoSmartSearch

import (
	"fmt"
	"sort"
	"strings"
)

type stringResult struct {
	value     string
	accurency float32
}

/*
Returns a map slice formed by the input elements ordered (based on the key) from most to least similar to the input term
*/
func SearchInMaps(elements []map[string]string, term string, key string, tolerance float32) ([]map[string]string, error) {
	if tolerance > 1 || tolerance < 0 {
		return nil, fmt.Errorf("validation error: %s", "'tolerance' must be a float32 from 0 to 1")
	}

	keyValues := make([]string, 0, len(elements))

	for _, item := range elements {
		keyValues = append(keyValues, item[key])
	}

	sortedKeyValues, err := SearchInStrings(keyValues, term, tolerance)

	if err != nil {
		return nil, err
	}

	result := make([]map[string]string, len(sortedKeyValues))

	for _, item := range sortedKeyValues {

		itemMap := findItemInMapSlice(elements, key, item)

		result = append(result, itemMap)
	}

	return result, nil

}

/*
Returns a slice formed by the input elements ordered from most to least similar to the input term
*/
func SearchInStrings(elements []string, term string, tolerance float32) ([]string, error) {

	if tolerance > 1 || tolerance < 0 {
		return nil, fmt.Errorf("validation error: %s", "'tolerance' must be a float32 from 0 to 1")
	}

	var tmpResult []stringResult

	for _, currentTerm := range elements {

		var resultObject stringResult
		resultObject.accurency = calculateAccurency(term, currentTerm)
		resultObject.value = currentTerm
		if resultObject.accurency >= tolerance {
			tmpResult = append(tmpResult, resultObject)
		}
	}

	sort.Slice(tmpResult, func(a, b int) bool {
		return tmpResult[a].accurency > tmpResult[b].accurency
	})

	result := make([]string, len(tmpResult))
	for i := range tmpResult {
		result[i] = tmpResult[i].value
	}

	return result, nil

}

func calculateAccurency(original string, current string) float32 {

	var hits, hitsExact float32
	var limit int

	if original == current {
		return 1
	}

	original, current = strings.ToLower(original), strings.ToLower(current)

	if original == current {
		return 1
	}

	if len(original) > len(current) {
		limit = len(current)
	} else {
		limit = len(original)
	}

	for i := 0; i < limit; i++ {
		if original[i] == current[i] {
			hitsExact++
		} else {
			for e := 0; e < limit; e++ {
				if (original[i] == current[e]) || (original[e] == current[i]) {
					hits += 0.25
				}
			}
		}
	}

	if int(hitsExact) == len(original) {
		return 1
	}

	hitsExact += hits

	return hitsExact / float32(len(original)) / 4
}

func findItemInMapSlice(elements []map[string]string, key string, value string) map[string]string {

	var result map[string]string
	for _, item := range elements {
		if item[key] == value {
			result = item
			break
		}
	}
	return result
}
