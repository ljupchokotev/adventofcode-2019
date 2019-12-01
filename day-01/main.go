package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		moduleFuel := calculateFuel(mass)
		fuel += moduleFuel
	}
	fmt.Println(fuel)
}

func calculateFuel(mass int) int {
	fuel := mass/3 - 2

	if fuel < 6 {
		return fuel
	}

	return fuel + calculateFuel(fuel)
}
