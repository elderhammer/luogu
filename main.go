package main

import (
	"fmt"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n, &m)

	ans := 0.0
	for i := 0; i < n; i++ {
		sum, min_score, max_score := 0.0, -1.0, 0.0
		for j := 0; j < m; j++ {
			score := 0.0
			fmt.Scan(&score)

			if min_score == -1.0 || min_score > score {
				min_score = score
			}
			if max_score < score {
				max_score = score
			}
			sum += score
		}
		sum -= min_score
		sum -= max_score
		sum /= float64(m - 2)
		if sum > ans {
			ans = sum
		}
	}
	fmt.Printf("%.2f\n", ans)
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
