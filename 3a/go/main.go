package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type power struct {
	gamma   int64
	epsilon int64
	power   int64
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	onesCounter := make([]int, 12) // not as dynamic as the TS solution
	totalCount := 0

	// reset scanner to get to beginning
	for fileScanner.Scan() {
		binary := fileScanner.Text()
		if binary != "" {
			totalCount++
			for i := 0; i < len(binary); i++ {
				onesCounter[i] += int(binary[i] - '0')
			}
		}
	}

	readFile.Close()

	pow := computePower(onesCounter, totalCount)
	log.Printf("%v", pow)
}

func computePower(ones []int, total int) power {
	p := power{gamma: 0, epsilon: 0, power: 0}
	gString := ""
	eString := ""
	for i := 0; i < len(ones); i++ {
		if ones[i] > total/2 {
			gString += "1"
			eString += "0"
		} else {
			gString += "0"
			eString += "1"
		}
	}

	var err error
	if p.gamma, err = strconv.ParseInt(gString, 2, 64); err != nil {
		log.Println(err)
	}
	if p.epsilon, err = strconv.ParseInt(eString, 2, 64); err != nil {
		log.Println(err)
	}

	p.power = p.gamma * p.epsilon
	return p
}
