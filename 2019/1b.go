package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	} else {
		return fuel + calculateFuel(fuel)
	}
}

func main() {
	input := "1.input"

	buf, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = buf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	totalFuel := 0

	for snl.Scan() {
		weight, err := strconv.Atoi(snl.Text())
		if err != nil {
			log.Fatal(err)
		}
		totalFuel += calculateFuel(weight)

	}
	fmt.Printf("%d\n", totalFuel)
}
