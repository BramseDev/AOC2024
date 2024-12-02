package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	data, err := getDataFromUrlAsFormattedString()
	if err != nil {
		fmt.Println("Fehler beim Abrufen der Daten:", err)
		return
	}

	parsedData, err := parseDataToIntSlices(data)
	if err != nil {
		fmt.Println("Fehler beim Verarbeiten der Daten:", err)
		return
	}

	fmt.Println(reportCalc(parsedData))
	fmt.Println(reportCalcPartTwo(parsedData))

}

func reportCalc(data [][]int) int {
	result := 0

	for i := 0; i < len(data); i++ {
		isIncreasing := true
		isDecreasing := true
		for j := 0; j < len(data[i])-1; j++ {
			diff := data[i][j+1] - data[i][j]

			if diff < -3 || diff > 3 || diff == 0 {
				isIncreasing = false
				isDecreasing = false
				break
			}

			if diff < 0 {
				isIncreasing = false
			}

			if diff > 0 {
				isDecreasing = false
			}
		}

		if isIncreasing || isDecreasing {
			result++
		}
	}

	return result
}

func reportCalcPartTwo(data [][]int) int {
	result := 0

	for i := 0; i < len(data); i++ {
		row := data[i]
		n := len(row)
		isSafe := false

		checkSafety := func(seq []int) bool {
			isIncreasing := true
			isDecreasing := true
			for j := 0; j < len(seq)-1; j++ {
				diff := seq[j+1] - seq[j]

				if diff < -3 || diff > 3 || diff == 0 {
					return false
				}
				if diff < 0 {
					isIncreasing = false
				}
				if diff > 0 {
					isDecreasing = false
				}
			}
			return isIncreasing || isDecreasing
		}
		if checkSafety(row) {
			result++
			continue
		}

		for j := 0; j < n; j++ {
			temp := append([]int{}, row[:j]...)
			temp = append(temp, row[j+1:]...)

			if checkSafety(temp) {
				isSafe = true
				break
			}
		}

		if isSafe {
			result++
		}
	}

	return result
}

func getDataFromUrlAsFormattedString() (string, error) {
	url := "https://adventofcode.com/2024/day/2/input"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Erstellen der Anfrage: %w", err)
	}

	req.Header.Set("Cookie", "session=<AOC-TOKEN")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Senden der Anfrage: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Lesen der Antwort: %w", err)
	}

	data := strings.TrimSpace(string(body))
	return data, nil
}

func parseDataToIntSlices(data string) ([][]int, error) {
	lines := strings.Split(data, "\n")
	result := [][]int{}

	for _, line := range lines {
		fields := strings.Fields(line)
		intLine := []int{}

		for _, str := range fields {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("Fehler beim Konvertieren von '%s' in eine Zahl: %w", str, err)
			}
			intLine = append(intLine, num)
		}

		result = append(result, intLine)
	}

	return result, nil
}
