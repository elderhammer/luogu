package main

import (
	"fmt"
)

func main() {
	var a string
	var b string
	fmt.Scan(&a, &b)

	a_len := len(a)
	b_len := len(b)
	num_len := a_len
	if b_len > a_len {
		num_len = b_len
	}

	c := []int{}
	for i := 0; i <= num_len; i++ {
		c = append(c, 0)
	}

	for i := num_len; i >= 0; i-- {
		sum := 0

		offset_to_tail := i - num_len

		a_idx := a_len - 1 + offset_to_tail
		b_idx := b_len - 1 + offset_to_tail
		if a_idx >= 0 && b_idx >= 0 {
			sum = int((a[a_idx] - '0') + (b[b_idx] - '0'))
		} else if a_idx >= 0 {
			sum = int((a[a_idx] - '0'))
		} else if b_idx >= 0 {
			sum = int((b[b_idx] - '0'))
		}
		sum = sum + c[i]
		c[i] = sum % 10
		if sum >= 10 {
			c[i-1] = 1
		}
	}

	meet_nzero := false
	for i := 0; i <= num_len; i++ {
		if meet_nzero {
			fmt.Print(c[i])
		} else {
			if c[i] > 0 {
				meet_nzero = true
				fmt.Print(c[i])
			}
		}
	}
	if !meet_nzero {
		fmt.Print(0)
	}
	fmt.Println()
}

func compose(m uint64, n uint64) uint64 {
	return factorial(n) / (factorial(n-m) * factorial(m))
}

func factorial(n uint64) uint64 {
	var product uint64 = 1
	for i := n; i > 0; i-- {
		product *= i
	}
	return product
}

func cal_index(base int, times int) int {
	product := 1
	for i := 0; i < times; i++ {
		product *= base
	}

	return product
}

func carve(o_x int, o_y int, edge_len int, diff_cube *[][]int) {
	// 递归基
	if edge_len == 1 {
		return
	}

	half_edge_len := edge_len / 2

	// 左上小正方形
	x1 := o_x
	y1 := o_y
	x2 := o_x + half_edge_len - 1
	y2 := o_y + half_edge_len - 1
	// 二维差分 -1
	(*diff_cube)[x1][y1] += (-1)
	(*diff_cube)[x1][y2+1] -= (-1)
	(*diff_cube)[x2+1][y1] -= (-1)
	(*diff_cube)[x2+1][y2+1] += (-1)

	// 右上小正方形
	carve(o_x, o_y+half_edge_len, half_edge_len, diff_cube)

	// 左下小正方形
	carve(o_x+half_edge_len, o_y, half_edge_len, diff_cube)

	// 右下小正方形
	carve(o_x+half_edge_len, o_y+half_edge_len, half_edge_len, diff_cube)
}

func make_cube(height int, width int, default_value rune) [][]rune {
	var cube [][]rune
	for i := 0; i < height; i++ {
		var some_row []rune
		for j := 0; j < width; j++ {
			some_row = append(some_row, default_value)
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
