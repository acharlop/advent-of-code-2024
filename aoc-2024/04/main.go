package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var path = "input/data.txt"
var data [][]string = [][]string{}
var letters = []string{"X", "M", "A", "S"}
var dirs = []int{-1, 0, 1}
var countWords = 0
var countXs = 0

func getDataFromFile() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		data = append(data, line)
	}
}

func isEdge(row, col int) bool {
	return row < 0 || col < 0 || row >= len(data) || col >= len(data[row])
}

func loop(row, col int) {
	if isEdge(row, col) {
		return
	}

	if findX(row, col) {
		countXs += 1
	}
	for _, i := range dirs {
		for _, j := range dirs {
			found := findWord(row, col, i, j, 0)
			if len(found) == len(letters) {
				countWords += 1
			}
		}
	}

	if col+1 == len(data[row]) {
		loop(row+1, 0)
	} else {
		loop(row, col+1)
	}
}

func findWord(row, col, addRow, addCol, idx int) [][]int {
	if isEdge(row, col) || data[row][col] != letters[idx] {
		return nil
	}

	if idx == len(letters)-1 {
		return [][]int{{row, col}}
	}

	return append(findWord(row+addRow, col+addCol, addRow, addCol, idx+1), []int{row, col})
}

func isValidX(a, b string) bool {
	str := a + b
	if str == "MS" || str == "SM" {
		return true
	}
	return false
}

func findX(row, col int) bool {
	if isEdge(row-1, col-1) || isEdge(row+1, col+1) || data[row][col] != "A" {
		return false
	}

	if isValidX(data[row-1][col-1], data[row+1][col+1]) &&
		isValidX(data[row-1][col+1], data[row+1][col-1]) {
		return true
	}

	return false
}

func main() {
	start := time.Now()

	getDataFromFile()

	loop(0, 0)

	elapsed := time.Since(start)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintf(w, "Day\t4\n")
	fmt.Fprintf(w, "Part\t1\t%d\n", countWords)
	fmt.Fprintf(w, "Part\t2\t%d\n", countXs)
	fmt.Fprintf(w, "Time\t\t%s\n", elapsed)
	w.Flush()
}
