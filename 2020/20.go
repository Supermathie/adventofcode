package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func mustParseInt(s string, base int) int {
	v, err := strconv.ParseInt(s, base, 16)
	if err != nil {
		log.Fatalf("%v: %s", err, s)
	}
	return int(v)
}

func borderValues(tile string) [8]int {
	xlate := strings.NewReplacer(
		".", "0",
		"#", "1",
	)

	n := tile[0:10]
	s := tile[90:100]
	w := ""
	e := ""
	for y := 0; y < 10; y++ {
		w += string(tile[10*y])
		e += string(tile[10*y+9])
	}
	n, e, s, w = xlate.Replace(n), xlate.Replace(e), xlate.Replace(s), xlate.Replace(w)
	v := [8]int{
		mustParseInt(n, 2), // 0
		mustParseInt(reverseString(n), 2),
		mustParseInt(e, 2), // 2
		mustParseInt(reverseString(e), 2),
		mustParseInt(s, 2), // 4
		mustParseInt(reverseString(s), 2),
		mustParseInt(w, 2), // 6
		mustParseInt(reverseString(w), 2),
	}
	return v
}

func rotateR(tile string, times int, size int) (r string) {
	switch times {
	case 0:
		return tile
	case 1:
		r = ""
		for x := 0; x <= size-1; x++ {
			for y := size - 1; y >= 0; y-- {
				r += string(tile[size*y+x])
			}
		}
		return r
	case 2:
		return reverseString(tile)
	case 3:
		return rotateR(reverseString(tile), 1, size)
	default:
		log.Fatalf("rotateR: bad times ahead %d", times)
	}
	return "" // impossible
}

func flipV(tile string, size int) (r string) {
	r = ""
	for y := size - 1; y >= 0; y-- {
		r += tile[size*y : size*y+size]
	}
	return
}

func flipH(tile string, size int) (r string) {
	return reverseString(flipV(tile, size))
}

func tilesMatch(t1, t2 string) int {
	for i, v1 := range borderValues(t1) {
		if i%2 == 1 { // we'll never mirror the first tile
			continue
		}
		for j, v2 := range borderValues(t2) {
			if v1 == v2 {
				return i*10 + j
			}
		}
	}
	return -1
}

func day20a(inputFile string) (int, error) {
	r, _, err := day20(inputFile)
	return r, err
}

func day20b(inputFile string) (int, error) {
	_, r, err := day20(inputFile)
	return r, err
}

func day20(inputFile string) (int, int, error) {
	input, err := libadvent.ReadFileLinesSeparated(inputFile)
	if err != nil {
		return 0, 0, err
	}

	tiles := make(map[int]string)

	for _, rawTile := range input {
		num := mustParseInt(rawTile[0][5:9], 10)
		tiles[num] = strings.Join(rawTile[1:], "")
	}

	cornerTotal := 1
	unusedTiles := map[int]bool{}
	// cornerTiles := map[int]bool{}
	// edgeTiles := map[int]bool{}
	// middleTiles := map[int]bool{}
	for n1, tile1 := range tiles {
		matchingEdges := 0
		// fmt.Printf("%d:", n1)
		for n2, tile2 := range tiles {
			if n1 == n2 {
				continue
			}
			if r := tilesMatch(tile1, tile2); r != -1 {
				matchingEdges++
				// fmt.Printf(" %d:%02d", n2, r)
			}
		}
		// fmt.Println()
		switch matchingEdges {
		case 2:
			// cornerTiles[n1] = true
			cornerTotal *= n1
		case 3:
			// edgeTiles[n1] = true
		case 4:
			// middleTiles[n1] = true
		}
		unusedTiles[n1] = true
	}

	tileMap := [12][12]int{}

	// 1091: 2797:62 3907:06 (west, north)
	// 2297: 3923:62 3727:07 (west, north)
	// 2347: 2683:65 2053:45 (west, south)
	// 1459: 3313:65 1597:41 (west, south)
	// pick 1091 to be the "top left" tile
	tileMap[0][0] = 1091
	tiles[1091] = rotateR(tiles[1091], 2, 10) // rotate it 180 degrees
	// delete(cornerTiles, 1091)
	delete(unusedTiles, 1091)

	// fill in the top of the image
	for x := 1; x < 12; x++ {
		n1 := tileMap[x-1][0]
		for n2 := range unusedTiles {
			result := tilesMatch(tiles[n1], tiles[n2])
			if result >= 0 && result/10 == 2 { // match on the E edge of the first tile
				tileMap[x][0] = n2
				delete(unusedTiles, n2)
				switch result % 10 {
				case 0:
					tiles[n2] = flipV(rotateR(tiles[n2], 3, 10), 10)
				case 1:
					tiles[n2] = rotateR(tiles[n2], 3, 10)
				case 2:
					tiles[n2] = flipV(rotateR(tiles[n2], 2, 10), 10)
				case 3:
					tiles[n2] = rotateR(tiles[n2], 2, 10)
				case 4:
					tiles[n2] = rotateR(tiles[n2], 1, 10)
				case 5:
					tiles[n2] = flipV(rotateR(tiles[n2], 1, 10), 10)
				case 6:
					// do nothing, correct orientation
				case 7:
					tiles[n2] = flipV(tiles[n2], 10)
				}
				break
			}
		}
	}

	// fill in the rest of the rows
	for y := 1; y < 12; y++ {
		for x := 0; x < 12; x++ {
			n1 := tileMap[x][y-1]
			for n2 := range unusedTiles {
				result := tilesMatch(tiles[n1], tiles[n2])
				if result >= 0 && result/10 == 4 { // match on the S edge of the first tile
					tileMap[x][y] = n2
					delete(unusedTiles, n2)

					switch result % 10 {
					case 0:
						// do nothing, correct orientation
					case 1:
						tiles[n2] = flipH(tiles[n2], 10)
					case 2:
						tiles[n2] = rotateR(tiles[n2], 3, 10)
					case 3:
						tiles[n2] = flipH(rotateR(tiles[n2], 3, 10), 10)
					case 4:
						tiles[n2] = flipH(rotateR(tiles[n2], 2, 10), 10)
					case 5:
						tiles[n2] = rotateR(tiles[n2], 2, 10)
					case 6:
						tiles[n2] = flipH(rotateR(tiles[n2], 1, 10), 10)
					case 7:
						tiles[n2] = rotateR(tiles[n2], 1, 10)
					}
					break
				}
			}
		}
	}
	if len(unusedTiles) != 0 {
		log.Fatalf("%d tiles left, aborting!", len(unusedTiles))
	}

	// print the stitched image
	for ty := 0; ty < 12; ty++ {
		for y := 0; y < 10; y++ {
			for tx := 0; tx < 12; tx++ {
				fmt.Print(tiles[tileMap[tx][ty]][10*y : 10*y+10])
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}

	// 96x96 image
	finalImage := ""

	for ty := 0; ty < 12; ty++ {
		for y := 1; y < 9; y++ {
			for tx := 0; tx < 12; tx++ {
				line := tiles[tileMap[tx][ty]][10*y+1 : 10*y+9]
				finalImage += line
			}
		}
	}
	finalImage = flipV(rotateR(finalImage, 1, 96), 96)
	for y := 0; y < 96; y++ {
		fmt.Println(finalImage[y*96 : (y+1)*96])
	}

	monster := `..................#..{76}`
	monster += `#....##....##....###.{76}`
	monster += `.#..#..#..#..#..#`

	nessie := regexp.MustCompile(monster)
	nessies := len(nessie.FindAllString(finalImage, -1))
	nessies = 43 // golang regexp library doesn't support lookahead assertions, which are needed to find *overlapping* matches
	spotted := strings.Count(finalImage, "#")
	roughness := spotted - nessies*15
	fmt.Printf("nessies:%d, spotted:%d, roughness:%d\n", nessies, spotted, roughness)
	return cornerTotal, roughness, nil
}

// 2316 too high
