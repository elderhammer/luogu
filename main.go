package main

import (
	"fmt"
)

type record struct {
	name          string
	chinese_score int
	math_score    int
	english_score int
}

func main() {
	var n int
	fmt.Scan(&n)

	var records []record
	for i := 0; i < n; i++ {
		var (
			name          string
			chinese_score int
			math_score    int
			english_score int
		)
		fmt.Scan(&name, &chinese_score, &math_score, &english_score)

		new_record := record{name, chinese_score, math_score, english_score}
		records = append(records, new_record)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if cmp(&records[i], &records[j]) {
				fmt.Printf("%s %s\n", records[i].name, records[j].name)
			}
		}
	}
}

func cmp(a_record *record, b_record *record) bool {
	if !in_range(-5, 5, a_record.chinese_score-b_record.chinese_score) {
		return false
	}

	if !in_range(-5, 5, a_record.math_score-b_record.math_score) {
		return false
	}

	if !in_range(-5, 5, a_record.english_score-b_record.english_score) {
		return false
	}

	if !in_range(-10, 10, a_record.chinese_score+a_record.english_score+a_record.math_score-(b_record.chinese_score+b_record.english_score+b_record.math_score)) {
		return false
	}

	return true
}

func in_range(min_num int, max_num int, num int) bool {
	return min_num <= num && num <= max_num
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
