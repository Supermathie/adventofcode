package main

import (
	"log"
	"math"
	"math/cmplx"
	"strconv"

	"supermathie.net/libadvent"
)

func day12b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	pos := 0 + 0i
	dPos := 10 + 1i

	for dir := range c {
		action := dir[0]
		val, err := strconv.ParseFloat(dir[1:], 64)
		if err != nil {
			log.Fatalf("could not parse %s as a float: %v", dir[1:], err)
		}
		switch action {
		case 'N':
			dPos += complex(0, val)
		case 'S':
			dPos -= complex(0, val)
		case 'E':
			dPos += complex(val, 0)
		case 'W':
			dPos -= complex(val, 0)
		case 'L': // CCW rotation by θ (in radians) in the complex plane → multiply by e**(iθ)
			θ := float64(val) * math.Pi / 180
			dPos *= cmplx.Pow(complex(math.E, 0), complex(0, θ))
		case 'R':
			θ := float64(val) * math.Pi / 180
			dPos /= cmplx.Pow(complex(math.E, 0), complex(0, θ))
		case 'F':
			pos += dPos * complex(val, 0)
		default:
			log.Fatalf("bad action %v", action)
		}
	}

	return int(math.Abs(real(pos)) + math.Abs(imag(pos))), nil
}
