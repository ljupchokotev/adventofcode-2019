package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var intcode []int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, x := range line {
			code, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}
			intcode = append(intcode, code)
		}
	}

	wantedResult := 19690720

	tryIntcode := make([]int, len(intcode))
	copy(tryIntcode, intcode)

	opCodePos := 0
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			tryIntcode[1] = noun
			tryIntcode[2] = verb

			for {
				opCode := tryIntcode[opCodePos]

				if opCode == 99 {
					if tryIntcode[0] == wantedResult {
						fmt.Println(100*noun + verb)
					}
					break
				}

				i1 := tryIntcode[opCodePos+1]
				i2 := tryIntcode[opCodePos+2]
				out := tryIntcode[opCodePos+3]
				if opCode == 1 {
					sum := tryIntcode[i1] + tryIntcode[i2]
					tryIntcode[out] = sum
				} else if opCode == 2 {
					mult := tryIntcode[i1] * tryIntcode[i2]
					tryIntcode[out] = mult
				}

				opCodePos += 4
			}
			copy(tryIntcode, intcode)
			opCodePos = 0
		}
	}
}
