package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getDataFromFile() []string {
	path := "input/data.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data []string

	re := *regexp.MustCompile(`(?mi)mul\(([0-9]+,[0-9]+)\)|(don)'t\(\)|(do)\(\)`)

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for i := range matches {
			if matches[i][0] == "do()" {
				data = append(data, "on")
			} else if matches[i][0] == "don't()" {
				data = append(data, "off")
			} else {
				data = append(data, matches[i][1])
			}
		}
	}

	return data
}

func doWork(data []string) int {
	var total int
	disabled := false

	for i := 0; i < len(data); i++ {
		if data[i] == "on" {
			disabled = false
			continue
		}

		if data[i] == "off" || disabled {
			disabled = true
			continue
		}

		nums := strings.Split(data[i], ",")

		if len(nums) != 2 {
			log.Fatal("invalid data")
		}

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		total += num1 * num2
	}

	return total
}

func main() {

	data := getDataFromFile()

	total := doWork(data)

	fmt.Println("answer:", total)
}
