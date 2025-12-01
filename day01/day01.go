package main

import (
	"fmt"
	"strconv"

	"github.com/Olegas/goaocd"
)

func main() {
	var input = goaocd.Lines(1)
	var ans = 0
	var pos = 50
	for _, l := range input {
		fmt.Println(l)
		var dir, num = l[0], l[1:]
		var step, _ = strconv.Atoi(num)
		var delta = 1
		if dir == 'L' {
			delta = -1
		}
		for range step {
			pos += delta
			pos = (pos + 100) % 100
			if pos == 0 {
				ans += 1
			}
		}
		fmt.Print(pos, " ", ans, "\n")
	}
	fmt.Println(ans)
	goaocd.Submit(2, ans, 1)
}
