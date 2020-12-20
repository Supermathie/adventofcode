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
