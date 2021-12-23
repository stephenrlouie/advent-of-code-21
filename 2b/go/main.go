package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x   int
	y   int
	aim int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var directions []string
	var magnitudes []int

	for fileScanner.Scan() {
		tokens := strings.Split(fileScanner.Text(), " ")
		if len(tokens) == 2 {
			if tokens[0] != "" {
				directions = append(directions, tokens[0])
			} else {
				log.Printf("Failed to parse the direction. %v", tokens)
			}

			if tokens[1] != "" {
				num, err := strconv.Atoi(tokens[1])
				if err != nil {
					log.Printf("failed to convert to integer: %s", err)
				} else {
					magnitudes = append(magnitudes, num)
				}
			} else {
				log.Printf("Failed to parse the magnitude")
			}
		}
	}

	readFile.Close()

	subPosition := position{}
	for i := 0; i < len(directions); i++ {
		subPosition.plotCourse(directions[i], magnitudes[i])
	}

	log.Printf("x: %d, y: %d, aim: %d. Mult: %d", subPosition.x, subPosition.y, subPosition.aim, subPosition.x*subPosition.y)
}

func (pos *position) plotCourse(dir string, mag int) {
	switch dir {
	case "up":
		pos.aim -= mag
	case "down":
		pos.aim += mag
	case "forward":
		pos.x += mag
		pos.y += mag * pos.aim
		if pos.y < 0 {
			pos.y = 0
		}
	default:
		log.Printf("direction unknown")
	}
}
