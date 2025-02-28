package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n, &m)

	pos_map_name := make(map[int]string)
	pos_map_direction := []int{}
	// 读取小人信息
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		input_str := strings.Fields(input)

		pos_map_name[i] = input_str[1]
		direction := -1
		if input_str[0] == "0" { // 0表示向内
			direction = 1
		}
		pos_map_direction = append(pos_map_direction, direction)
		// 后续向左改为-1，向右改为1，然后乘以方向，就确定了是顺时针还是逆时针
	}

	length := len(pos_map_direction)

	next_pos := 0
	for i := 0; i < m; i++ {
		input, _ := reader.ReadString('\n')
		input_str := strings.Fields(input)

		left_or_right := 1
		if input_str[0] == "0" {
			left_or_right = -1
		}

		step, _ := strconv.Atoi(input_str[1])

		next_pos = next_pos + (left_or_right*pos_map_direction[next_pos])*step
		if next_pos < 0 {
			next_pos = length - ((-1 * next_pos) % length)
		} else {
			next_pos = next_pos % length
		}
	}
	fmt.Println(pos_map_name[next_pos])
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
