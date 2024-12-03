package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func getDataFromFile() [][]int {
	path := "input/data.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var arr [][]int
	lineNum := -1
	var num int

	for scanner.Scan() {
		lineNum++

		line := strings.Split(scanner.Text(), " ")
		arr = append(arr, []int{})

		for i := 0; i < len(line); i++ {
			num, _ = strconv.Atoi(line[i])
			arr[lineNum] = append(arr[lineNum], num)
		}
	}

	return arr
}

func isValidDiff(a, b int, descending bool) bool {
	diff := a - b

	if descending {
		return diff < 0 && diff > -4
	}

	return diff > 0 && diff < 4
}

func isValidList(list []int) bool {
	valid := true
	descending := false

	for i := 0; i < len(list); i++ {
		if i == 0 {
			if list[i] == list[i+1] {
				valid = false
				break
			}

			descending = list[i] > list[i+1]
			continue
		}

		if !isValidDiff(list[i], list[i-1], descending) {
			valid = false
			break
		}
	}

	return valid
}

func validateByRemoving(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		arr1 := append([]int{}, arr[:i]...)
		arr1 = append(arr1, arr[i+1:]...)
		if isValidList(arr1) {
			return true
		}
	}
	return false
}

func findDiffs(arr [][]int) int {
	validCount := 0

	for i := 0; i < len(arr); i++ {
		if  isValidList(arr[i]) {
			validCount++
		}
	}

	return validCount
}

func findDiffsWithTolerance(arr [][]int) int {
	validCount := 0
	// invalidCount := 0

	for i := 0; i < len(arr); i++ {
		if isValidList(arr[i]) || validateByRemoving(arr[i]) {
			validCount++
			continue
		}
	}

	return validCount
}

func day2() {
	start := time.Now()

	arr := getDataFromFile()

	elapsed := time.Since(start)

	count := findDiffs(arr)
	countWithTolerance := findDiffsWithTolerance(arr)

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintf(w, "Day\t2\n")
	fmt.Fprintf(w, "Part\t1\t%d\n", count)
	fmt.Fprintf(w, "Part\t2\t%d\n", countWithTolerance)
	fmt.Fprintf(w, "Time\t\t%s\n", elapsed)
	w.Flush()
}

	func main() {
		day2()
	}
