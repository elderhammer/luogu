package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	edge_len := _2_index(n)

	// 差分矩阵
	diff_cube := make_cube(edge_len, edge_len, 0)

	// 切分
	carve(0, 0, edge_len, &diff_cube)

	// 计算前缀和
	for row := 0; row < edge_len; row++ {
		for col := 0; col < edge_len; col++ {
			if row == 0 && col == 0 {
				// diff_cube[row][col] = diff_cube[row][col]
			} else if row == 0 {
				diff_cube[row][col] = diff_cube[row][col] + diff_cube[row][col-1]
			} else if col == 0 {
				diff_cube[row][col] = diff_cube[row][col] + diff_cube[row-1][col]
			} else {
				diff_cube[row][col] = diff_cube[row][col] + diff_cube[row][col-1] + diff_cube[row-1][col] - diff_cube[row-1][col-1]
			}
			// fmt.Printf("%d ", diff_cube[row][col])
		}
		// fmt.Println()
	}

	// 投影到原矩阵
	ori_cube := make_cube(edge_len, edge_len, 1)
	for row := 0; row < edge_len; row++ {
		for col := 0; col < edge_len; col++ {
			ori_cube[row][col] += diff_cube[row][col]
			fmt.Printf("%d ", ori_cube[row][col])
		}
		fmt.Println()
	}
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

func _2_index(times int) int {
	product := 1
	for i := 0; i < times; i++ {
		product *= 2
	}

	return product
}

func make_cube(height int, width int, default_value int) [][]int {
	var cube [][]int
	for i := 0; i < height; i++ {
		var some_row []int
		for j := 0; j < width; j++ {
			some_row = append(some_row, default_value)
		}
		cube = append(cube, some_row)
	}

	return cube
}

func mul(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * mul(n-1)
	}
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
