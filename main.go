package main

import (
	"fmt"
)

type record struct {
	num         int
	study_score int
	pe_score    int
	sum_up      float64 // x10
}

func (r record) total_score() int {
	return r.pe_score + r.study_score
}

func main() {
	var n int
	fmt.Scan(&n)

	records := []record{}
	for i := 0; i < n; i++ {
		var (
			num         int
			study_score int
			pe_score    int
		)

		fmt.Scan(&num, &study_score, &pe_score)

		sum_up := study_score*7 + pe_score*3
		records = append(records, record{num, study_score, pe_score, float64(sum_up)})
	}

	for i := 0; i < n; i++ {
		if comment(&records[i]) {
			fmt.Println("Excellent")
		} else {
			fmt.Println("Not excellent")
		}
	}
}

func comment(r *record) bool {
	return r.total_score() > 140 && r.sum_up >= 800.0
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
