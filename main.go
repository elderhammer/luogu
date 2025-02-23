package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	fmt.Scan(&x, &y)

	var leap_years []int
	for i := x; i <= y; i++ {
		if (i%4 == 0 && i%100 != 0) || i%400 == 0 {
			leap_years = append(leap_years, i)
		}
	}

	fmt.Println(len(leap_years))
	for i := 0; i < len(leap_years); i++ {
		fmt.Printf("%d ", leap_years[i])
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
