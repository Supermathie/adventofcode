package libadvent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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

	var vals []int

	for snl.Scan() {
		val, err := strconv.Atoi(snl.Text())
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}
