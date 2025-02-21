package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	codex := make(map[string]int)
	codex["a"] = 1
	codex["both"] = 2
	codex["first"] = 1
	codex["second"] = 2
	codex["third"] = 3
	codex["another"] = 1
	codex["one"] = 1
	codex["two"] = 2
	codex["three"] = 3
	codex["four"] = 4
	codex["five"] = 5
	codex["six"] = 6
	codex["seven"] = 7
	codex["eight"] = 8
	codex["nine"] = 9
	codex["ten"] = 10
	codex["eleven"] = 11
	codex["twelve"] = 12
	codex["thirteen"] = 13
	codex["fourteen"] = 14
	codex["fifteen"] = 15
	codex["sixteen"] = 16
	codex["seventeen"] = 17
	codex["eighteen"] = 18
	codex["nineteen"] = 19
	codex["twenty"] = 20

	nums := []int{}
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		num := codex[words[i]]
		if num > 0 {
			nums = append(nums, num*num)
		}
	}

	if len(nums) > 0 {
		sort.Ints(nums)
		for i := 0; i < len(nums); i++ {
			if i == 0 {
				fmt.Printf("%d", nums[i])
			} else {
				fmt.Printf("%02d", nums[i])
			}
		}
		fmt.Println()
	} else {
		fmt.Println(0)
	}
}

func reverse_num(input []byte) {
	if len(input) == 1 {
		fmt.Printf("%c", input[0])
	} else {
		for i := 0; i < len(input)/2; i++ {
			input[i], input[len(input)-i-1] = input[len(input)-i-1], input[i]
		}
		meet_not_zero := false
		for i := 0; i < len(input); i++ {
			if input[i] == '0' && !meet_not_zero {
				// 不输出
			} else {
				meet_not_zero = true
				fmt.Printf("%c", input[i])
			}
		}
	}
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
