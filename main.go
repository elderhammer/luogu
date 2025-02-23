package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x1 float64
		y1 float64
		x2 float64
		y2 float64
		x3 float64
		y3 float64
	)

	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	d_1 := dis(x1, y1, x2, y2)
	d_2 := dis(x1, y1, x3, y3)
	d_3 := dis(x2, y2, x3, y3)

	fmt.Printf("%.2f\n", d_1+d_2+d_3)
}

func dis(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	d_x := x2 - x1
	d_y := y2 - y1

	return math.Sqrt(d_x*d_x + d_y*d_y)
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
