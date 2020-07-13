package utils

import (
	"math"
	"sort"
	"strconv"
)

func GetMean(dataset map[string][]string, label string) float64 {
	var mean float64
	sum := 0.0
	for _, value := range dataset[label] {
		if parsedValue, err := strconv.ParseFloat(value, 64); err == nil {
			sum += parsedValue
			continue
		}
	}
	mean = sum / float64(len(dataset[label]))
	return mean
}

func GetMedian(dataset map[string][]string, label string) float64 {
	sortedList := getSortedNumericStringsList(dataset[label])
	medianIndex := math.Ceil(float64(len(sortedList)) / 2)
	isDatasetLengthEven := len(sortedList) % 2 == 0
	if isDatasetLengthEven {
		parsedFloatElement := getParsedFloatElement(sortedList, medianIndex)
		return (parsedFloatElement * 2 + 1) / 2
	}
	return getParsedFloatElement(sortedList, medianIndex)
}

func GetMode(dataset map[string][]string, label string) []string {
	var modeHighest []string
	mode := make(map[string]int)
	modeHighest = append(modeHighest, "0")
	for _, value := range dataset[label] {
		mode[value]++
		for _, highest := range modeHighest {
			haveDifferentValues := highest != value
			if mode[value] > mode[highest] {
				modeHighest = append(modeHighest[0:0], value)
				break
			} else if haveDifferentValues && mode[value] == mode[highest] {
				modeHighest = append(modeHighest, value)
				break
			}
		}
	}
	return modeHighest
}

func GetRange(dataset map[string][]string, label string) float64 {
	var statisticalRange float64
	sortedList := getSortedNumericStringsList(dataset[label])
	lowest, _ := strconv.ParseFloat(sortedList[0], 64)
	highest, _ := strconv.ParseFloat(sortedList[len(sortedList) - 1], 64)
	statisticalRange = highest - lowest
	return statisticalRange
}


func GetHighest(dataset map[string][]string, label string) float64 {
	sortedList := getSortedNumericStringsList(dataset[label])
	highest, _ := strconv.ParseFloat(sortedList[len(sortedList) - 1], 64)
	return highest
}
func GetLowest(dataset map[string][]string, label string) float64 {
	sortedList := getSortedNumericStringsList(dataset[label])
	lowest, _ := strconv.ParseFloat(sortedList[0], 64)
	return lowest
}

func getSortedNumericStringsList(list []string) []string {
	sortedList := list
	sort.SliceStable(sortedList, func(i, j int) bool {
		first, _ := strconv.ParseFloat(sortedList[i], 64)
		second, _ := strconv.ParseFloat(sortedList[j], 64)
		return first < second
	})
	return sortedList
}

func getParsedFloatElement(list []string, index float64) float64 {
	parsedFloatElement, _ := strconv.ParseFloat(list[int(index)], 64)
	return parsedFloatElement
}