package main

import "fmt"

type node struct {
	val  int
	next *node
}

func day23a(input string) (int, error) {
	cup1 := day23(input, len(input), 100)
	ans := 0
	for cur := cup1.next; cur.val != 1; cur = cur.next {
		ans = 10*ans + cur.val
	}
	return ans, nil
}
func day23b(input string) (int, error) {
	cup1 := day23(input, 1_000_000, 10_000_000)
	ans := cup1.next.val * cup1.next.next.val
	return ans, nil
}

func day23(input string, numCups int, moves int) *node {
	var cur, last *node
	index := map[int]*node{}
	{
		for i, c := range input {
			num := int(c - '0')
			new := &node{num, cur}
			index[num] = new
			if i == 0 {
				cur = new
				cur.next = new
				last = new
			} else {
				last.next = new
				last = new
			}
		}
	}

	for i := len(input) + 1; i <= numCups; i++ {
		new := &node{i, cur}
		index[i] = new
		last.next = new
		last = new
	}

	for t := 0; t < moves; t++ {
		if false {
			ans := 0
			for i := index[1].next; i.val != 1; i = i.next {
				ans = 10*ans + i.val
			}
			fmt.Printf("t:%d %d\n", t, ans)
		}

		selected := cur.next
		cur.next = selected.next.next.next

		destinationNum := cur.val - 1
		if destinationNum < 1 {
			destinationNum = numCups
		}

		for selected.val == destinationNum || selected.next.val == destinationNum || selected.next.next.val == destinationNum {
			destinationNum--
			if destinationNum < 1 {
				destinationNum = numCups
			}
		}
		destination := index[destinationNum]

		selected.next.next.next = destination.next
		destination.next = selected
		cur = cur.next
	}

	return index[1]
}
