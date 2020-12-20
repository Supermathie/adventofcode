package main

import (
	"testing"
)

func Test20FlipV1(t *testing.T) {
	data := "##....##....#.#.##.#.###....##....#..#.##.##.........#......###...#.##...##....#..#.#...#.##.#.#.#.."
	want := "##.#.#.#....#.#...#....##....####...#.##...#......#.##..........#..#.#.###....##..#.#.##.###....##.."

	result := flipV(data, 10)

	if result != want {
		t.Fatalf("day20 flipV(%s) → (%s), expected %s", data, result, want)
	}
}
func Test20FlipH1(t *testing.T) {
	data := "##....##....#.#.##.#.###....##....#..#.##.##.........#......###...#.##...##....#..#.#...#.##.#.#.#.."
	want := "..##....###.##.#.#..##....###.#.#..#..........##.#......#...##.#...####....##....#...#.#....#.#.#.##"

	result := flipH(data, 10)

	if result != want {
		t.Fatalf("day20 flipH(%s) → (%s), expected %s", data, result, want)
	}
}

// ..................#.
// #....##....##....###
// .#..#..#..#..#..#...
func TestFindNessie1(t *testing.T) {
	data := "..................#.#....##....##....###.#..#..#..#..#..#..."
	want := 1

	result := findNessie(data, 20)

	if result != want {
		t.Fatalf("day20 findNessie(%s) → (%d), expected %d", data, result, want)
	}
}

// ......................
// ...................#..
// .#....##....##....###.
// ..#..#..#..#..#..#....
// ......................

func TestFindNessie2(t *testing.T) {
	data := ".........................................#...#....##....##....###...#..#..#..#..#..#.........................."
	want := 1

	result := findNessie(data, 22)

	if result != want {
		t.Fatalf("day20 findNessie(%s) → (%d), expected %d", data, result, want)
	}
}
