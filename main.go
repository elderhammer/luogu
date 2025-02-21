package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	codex := make(map[rune]int, 0)
	codex[' '] = 1
	codex['a'] = 1
	codex['b'] = 2
	codex['c'] = 3
	codex['d'] = 1
	codex['e'] = 2
	codex['f'] = 3
	codex['g'] = 1
	codex['h'] = 2
	codex['i'] = 3
	codex['j'] = 1
	codex['k'] = 2
	codex['l'] = 3
	codex['m'] = 1
	codex['n'] = 2
	codex['o'] = 3
	codex['p'] = 1
	codex['q'] = 2
	codex['r'] = 3
	codex['s'] = 4
	codex['t'] = 1
	codex['u'] = 2
	codex['v'] = 3
	codex['w'] = 1
	codex['x'] = 2
	codex['y'] = 3
	codex['z'] = 4

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	count := 0
	for _, c := range message {
		count += codex[c]
	}
	fmt.Println(count)
}

func num_len(num int) int {
	len := 0
	for num > 0 {
		num /= 10
		len += 1
	}

	return 1
}

func make_cube(n int) [][]rune {
	var cube [][]rune
	for i := 0; i < n; i++ {
		var some_row []rune
		for j := 0; j < n; j++ {
			some_row = append(some_row, 0)
		}
		cube = append(cube, some_row)
	}

	return cube
}
