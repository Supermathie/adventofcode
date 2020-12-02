package libadvent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadFileInts reads an array of line-separated integers from a file
func ReadFileInts(input string) ([]int, error) {
	buf, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = buf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)

	vals := make([]int, 0)

	for snl.Scan() {
		val, err := strconv.Atoi(snl.Text())
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

// ReadFileLines reads an array of lines from a file
func ReadFileLines(input string) (chan string, error) {
	buf, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	c := make(chan string)

	go func() {
		snl := bufio.NewScanner(buf)
		for snl.Scan() {
			c <- snl.Text()
		}
		close(c)
		err = buf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return c, nil
}
