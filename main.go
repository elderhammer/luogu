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

	cube := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, i*n+j+1)
		}
		cube = append(cube, row)
	}

	// 不要每次都创建临时数组
	temp_cube := make_cube(n, n, 0)

	reader := bufio.NewReader(os.Stdin)
	x, y, r, z := 0, 0, 0, 0
	for i := 0; i < m; i++ {
		input, _ := reader.ReadString('\n')
		input_nums := strings.Fields(input)
		x, _ = strconv.Atoi(input_nums[0])
		y, _ = strconv.Atoi(input_nums[1])
		r, _ = strconv.Atoi(input_nums[2])
		z, _ = strconv.Atoi(input_nums[3])

		e := 2*r + 1
		x1 := x - 1 - e/2
		y1 := y - 1 - e/2

		for row := 0; row < e; row++ {
			for col := 0; col < e; col++ {
				x_ := e - col - 1
				y_ := row
				if z == 1 {
					x_ = col
					y_ = e - row - 1
				}
				x_ += x1
				y_ += y1
				temp_cube[row][col] = cube[x_][y_]
			}
		}

		for row := 0; row < e; row++ {
			for col := 0; col < e; col++ {
				cube[x1+row][y1+col] = temp_cube[row][col]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", cube[i][j])
		}
		fmt.Println()
	}
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

func a_plus_b(a *[]int, b *[]int) []int {
	c := [100001]int{}

	a_len := len(*a)
	b_len := len(*b)

	// 计算，不用考虑进位
	for i := a_len - 1; i >= 0; i-- {
		for j := b_len - 1; j >= 0; j-- {
			// 错位相加
			c[(a_len-i-1)+(b_len-j-1)] += (*a)[i] * (*b)[j]
		}
	}

	// 进位
	for i := 0; i < a_len+b_len; i++ {
		if c[i] > 9 {
			c[i+1] += c[i] / 10
			c[i] = c[i] % 10
		}
	}

	// 去掉开头的0
	len := a_len + b_len
	for c[len] == 0 && len > 0 {
		len -= 1
	}

	// 翻转顺序
	r_c := []int{}
	for i := len; i >= 0; i-- {
		r_c = append(r_c, c[i])
	}

	return r_c
}

func pow_10_str(times int) string {
	product := ""
	for i := 0; i < times; i++ {
		product += "0"
	}

	return product
}

func a_add_b(a *string, b *string) {
	// 对齐
	if len(*a) > len(*b) {
		for i := len(*a) - len(*b); i > 0; i-- {
			*b = "0" + *b
		}
	} else {
		for i := len(*b) - len(*a); i > 0; i-- {
			*a = "0" + *a
		}
	}

	c := []int{}
	for i := len(*a) - 1; i >= 0; i-- {
		c = append(c, int((*a)[i]-'0')+int((*b)[i]-'0'))
	}

	// 进位
	for i := 0; i < len(c); i++ {
		if c[i] > 9 {
			c[i+1] += c[i] / 10
			c[i] = c[i] % 10
		}
	}

	// 去掉开头的0
	pos := len(c) - 1
	for c[pos] == 0 && pos > 0 {
		pos -= 1
	}

	// 打印
	for i := pos; i >= 0; i-- {
		fmt.Print(c[i])
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
