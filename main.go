package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n, &m)

	cube := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		some_row := []int{}
		scanner.Scan()
		input := scanner.Text()
		for j := 0; j < len(input); j++ {
			if input[j] == '*' { // 1表示地雷
				some_row = append(some_row, 1)
			} else {
				some_row = append(some_row, 0)
			}
		}
		cube = append(cube, some_row)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 是否是地雷格
			if cube[i][j] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(count_mine(&cube, n, m, i, j))
			}
		}
		fmt.Println()
	}
}

func count_mine(cube *[][]int, height int, width int, x int, y int) int {
	count := 0

	// 左上
	if is_mine(cube, height, width, x-1, y-1) {
		count += 1
	}

	// 上
	if is_mine(cube, height, width, x-1, y) {
		count += 1
	}

	// 右上
	if is_mine(cube, height, width, x-1, y+1) {
		count += 1
	}

	// 右
	if is_mine(cube, height, width, x, y+1) {
		count += 1
	}

	// 右下
	if is_mine(cube, height, width, x+1, y+1) {
		count += 1
	}

	// 下
	if is_mine(cube, height, width, x+1, y) {
		count += 1
	}

	// 左下
	if is_mine(cube, height, width, x+1, y-1) {
		count += 1
	}

	// 左
	if is_mine(cube, height, width, x, y-1) {
		count += 1
	}

	return count
}

func is_mine(cube *[][]int, height int, width int, x int, y int) bool {
	if x < 0 || x >= height {
		return false
	}

	if y < 0 || y >= width {
		return false
	}

	return (*cube)[x][y] == 1
}

func read() string {
	var input_str string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		for i := 0; i < len(input); i++ {
			// 判断是否结束输入
			if input[i] == 'E' {
				input_str += input[:i]
				return input_str
			}
		}
		input_str += input
	}

	return input_str
}

func stat(input_str *string, game_point int) {
	win, lose := 0, 0
	for i := 0; i < len(*input_str); i++ {
		// 记分
		if (*input_str)[i] == 'W' {
			win += 1
		} else {
			lose += 1
		}
		// 其中一方到达赛点
		if win >= game_point || lose >= game_point {
			diff := win - lose
			if diff <= -2 || 2 <= diff {
				fmt.Printf("%d:%d\n", win, lose)
				win, lose = 0, 0 // 重置记分
			}
		}
	}
	// 0比0都要输出
	fmt.Printf("%d:%d\n", win, lose)
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
