package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count [26]int
	height := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		input := scanner.Text()
		for j := 0; j < len(input); j++ {
			if 'A' <= input[j] && input[j] <= 'Z' {
				idx := input[j] - 'A'
				count[idx] += 1
				if count[idx] > height {
					height = count[idx]
				}
			}
		}
	}

	cube := make_cube(height, 26)
	for col := 0; col < 26; col++ {
		for row := height - 1; row >= 0 && count[col] > 0; row -= 1 {
			cube[row][col] = '*'
			count[col] -= 1
		}
	}

	for row := 0; row < height; row++ {
		for col := 0; col < 26; col++ {
			if col == 25 {
				fmt.Printf("%c", cube[row][col])
			} else {
				fmt.Printf("%c ", cube[row][col])
			}
		}
		fmt.Println()
	}
	for col := 0; col < 26; col++ {
		c := 'A' + col
		if col == 25 {
			fmt.Printf("%c", c)
		} else {
			fmt.Printf("%c ", c)
		}
	}
	fmt.Println()
}

func make_cube(height int, width int) [][]rune {
	var cube [][]rune
	for i := 0; i < height; i++ {
		var some_row []rune
		for j := 0; j < width; j++ {
			some_row = append(some_row, ' ')
		}
		cube = append(cube, some_row)
	}

	return cube
}

func n2i(n rune) int {
	return int(n) - int('0')
}

func a2i(a rune) int {
	return int(a) - int('A') + 1
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
