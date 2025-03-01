package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cube := [10][10]rune{}

	scanner := bufio.NewScanner(os.Stdin)
	f_x, f_y := 0, 0 // 农夫的初始位置
	c_x, c_y := 0, 0 // 牛的初始位置
	for i := 0; i < 10; i++ {
		scanner.Scan()
		input := scanner.Text()
		for j, c := range input {
			if c == 'F' {
				f_x, f_y = i, j
				c = '.'
			} else if c == 'C' {
				c_x, c_y = i, j
				c = '.'
			}
			cube[i][j] = c
		}
	}

	f_d, c_d := 0, 0 // 农夫和牛的初始方向，0上、1右、2下、3左，不停循环，直到两者同时在一个格子中

	zt := [160005]int{}

	step := 0
	for {
		flag := f_x + f_y*10 + c_x*100 + c_y*1000 + f_d*10000 + c_d*40000
		if zt[flag] == 1 {
			fmt.Println(0)
			return
		} else {
			zt[flag] = 1
		}

		// 农夫行动
		if front_is_edge(f_x, f_y, f_d) {
			f_d = change_direction(f_d)
		} else {
			front := check_front(&cube, f_x, f_y, f_d)
			switch front {
			case '*': // 障碍物，调转方向
				f_d = change_direction(f_d)
			default: // 空地，前进
				// cube[f_x][f_y] = '.'
				f_x, f_y = move(f_x, f_y, f_d)
				// cube[f_x][f_y] = 'F'
			}
		}

		// 牛行动
		if front_is_edge(c_x, c_y, c_d) {
			c_d = change_direction(c_d)
		} else {
			front := check_front(&cube, c_x, c_y, c_d)
			switch front {
			case '*': // 障碍物，调转方向
				c_d = change_direction(c_d)
			default: // 空地，前进
				// cube[c_x][c_y] = '.'
				c_x, c_y = move(c_x, c_y, c_d)
				// cube[c_x][c_y] = 'C'
			}
		}

		step += 1

		// 判断状态
		if f_x == c_x && f_y == c_y {
			break
		}

		// 打印
		// for i := 0; i < 10; i++ {
		// 	for j := 0; j < 10; j++ {
		// 		fmt.Printf("%c", cube[i][j])
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()

		// time.Sleep(time.Millisecond)
	}

	fmt.Println(step)
}

func change_direction(d int) int {
	return (d + 1) % 4
}

func front_is_edge(x int, y int, d int) bool {
	x, y = move(x, y, d)
	if x < 0 || y < 0 || x >= 10 || y >= 10 {
		return true
	}

	return false
}

func check_front(cube *[10][10]rune, x int, y int, d int) rune {
	x, y = move(x, y, d)

	return (*cube)[x][y]
}

func move(x int, y int, d int) (int, int) {
	switch d {
	case 0: //向上
		x = x - 1
	case 1: //向右
		y = y + 1
	case 2: //向下
		x = x + 1
	case 3: //向左
		y = y - 1
	}
	return x, y
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
