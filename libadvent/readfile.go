package libadvent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadFileLines reads an array of lines from a file and
// returns a channel that sends the results
func ReadFileLines(input string) (chan string, error) {
	buf, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	c := make(chan string)

	go func() {
		defer func() {
			close(c)
			err = buf.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		snl := bufio.NewScanner(buf)
		for snl.Scan() {
			c <- snl.Text()
		}
	}()
	return c, nil
}

// ReadFileInts reads an array of line-separated integers from a file
func ReadFileInts(input string) ([]int, error) {
	buf, err := os.Open(input)
	if err != nil {
		return nil, err
	}

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
