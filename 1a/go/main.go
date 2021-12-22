package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers []int

	for fileScanner.Scan() {
		var num int
		num, err = strconv.Atoi(fileScanner.Text())
		if err != nil {
			log.Printf("failed to convert to integer: %s", err)
		} else {
			numbers = append(numbers, num)
		}
	}

	readFile.Close()

	var count = 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			count++
		}
	}
	log.Println(count)
}
