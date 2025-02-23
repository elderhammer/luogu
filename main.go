package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var notp [100001]int // 0表示素数，1表示合数
	notp[0], notp[1] = -1, -1

	var primes []int
	for i := 2; i < len(notp); i++ {
		if notp[i] == 0 {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= len(notp) {
				// 超过边界
				break
			}
			notp[i*p] = 1 // p的i倍是合数
			if i%p == 0 {
				// i找到了其最小素因子
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)

		if notp[num] == 0 {
			fmt.Printf("%d ", num)
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
