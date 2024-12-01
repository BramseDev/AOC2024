package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	dataArray, err := getDataFromUrlAsIntArray()
	if err != nil {
		fmt.Println("Fehler beim Abrufen der Daten:", err)
		return
	}

	leftArr, rightArr := setUpArrays(dataArray)
	result := calcDifference(leftArr, rightArr)
	fmt.Println("result1 ", result)

	result2 := calcSecondTask(leftArr, rightArr)
	fmt.Println("result2 ", result2)
}

func calcSecondTask(leftArr []int, rightArr []int) int {
	result := 0

	for i := 0; i < len(leftArr); i++ {
		multiplier := 0
		for j := 0; j < len(leftArr); j++ {
			if leftArr[i] == rightArr[j] {
				multiplier++
			}
		}
		result += leftArr[i] * multiplier
	}
	return result
}

func calcDifference(leftArr []int, rightArr []int) int {
	var result int = 0
	for i := 0; i < len(leftArr); i++ {
		if leftArr[i] > rightArr[i] {
			result += (leftArr[i] - rightArr[i])
		} else if leftArr[i] < rightArr[i] {
			result += rightArr[i] - leftArr[i]
		} else {
			result += 0
		}

	}

	return result
}

func setUpArrays(dataArray []int) ([]int, []int) {
	leftArr := make([]int, (len(dataArray)+1)/2)
	rightArr := make([]int, len(dataArray)/2)

	for i := 0; i < len(dataArray); i += 2 {
		leftArr[i/2] = dataArray[i]
		if i+1 < len(dataArray) {
			rightArr[i/2] = dataArray[i+1]
		}
	}

	myArraySorter(leftArr)
	myArraySorter(rightArr)
	return leftArr, rightArr
}

func getDataFromUrlAsIntArray() ([]int, error) {
	url := "https://adventofcode.com/2024/day/1/input"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Erstellen der Anfrage: %w", err)
	}

	req.Header.Set("Cookie", "session=<AOC-TOKEN>")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Senden der Anfrage: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen der Antwort: %w", err)
	}

	data := strings.TrimSpace(string(body))
	fields := strings.Fields(data)

	intArray := make([]int, len(fields))
	for i, str := range fields {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("Fehler beim Konvertieren von '%s' in eine Zahl: %w", str, err)
		}
		intArray[i] = num
	}

	return intArray, nil
}

func myArraySorter(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivotIndex := len(arr) / 2

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	myArraySorter(arr[:left])
	myArraySorter(arr[left+1:])

	return arr
}
