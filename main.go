package main

import (
	"fmt"
	"math"
)

func main() {
	var p int
	fmt.Scan(&p)

	// 求位数，利用指数对数的知识
	fmt.Println(int(math.Log10(2)*float64(p)) + 1)

	ans := []int{2}
	times := []int{1}

	for p > 0 {
		if len(ans) > 500 {
			ans = ans[:500]
		}
		if len(times) > 500 {
			times = times[:500]
		}
		if p == 1 {
			p -= 1
		} else if p%2 == 1 {
			times = apb(&times, &ans)
			p -= 1
		} else {
			p /= 2
			ans = apb(&ans, &ans)
		}
	}
	ans = apb(&ans, &times)

	ans = amb(&ans, &[]int{1})

	j := 0
	for i := 499; i >= 0; i-- {
		fmt.Print(ans[i])
		j++
		if j%50 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func amb(a *[]int, b *[]int) []int {
	// 长度补齐
	a_len := len(*a)
	b_len := len(*b)
	if b_len > a_len {
		for i := 0; i < b_len-a_len; i++ {
			(*a) = append((*a), 0)
		}
	}

	// 比较最高位
	a_highest := len(*a) - 1
	for (*a)[a_highest] == 0 && a_highest > 0 {
		a_highest--
	}
	b_highest := len(*b) - 1
	for (*b)[b_highest] == 0 && b_highest > 0 {
		b_highest--
	}

	highest_symbol := 1
	if a_highest < b_highest || (a_highest == b_highest && (*a)[a_highest] < (*b)[b_highest]) {
		tmp := a
		a = b
		b = tmp
		highest_symbol = -1
	}

	// 逐位减去
	for i := 0; i < len(*b); i++ {
		if (*a)[i] >= (*b)[i] {
			(*a)[i] -= (*b)[i]
		} else {
			if i+1 >= len(*a) { // 没法借位了
				(*a)[i] -= (*b)[i]
			} else {
				(*a)[i+1] -= 1
				(*a)[i] += 10 - (*b)[i]
			}
		}
	}

	// 补位
	for i := 0; i < len(*a)-1; i++ {
		if (*a)[i] < 0 {
			(*a)[i] += 10
			(*a)[i+1] -= 1
		}
	}

	// 去掉前面的0
	len := len(*a) - 1
	for (*a)[len] == 0 && len > 0 {
		len--
	}

	// 校正符号
	(*a)[len] *= highest_symbol

	return (*a)
}

func apb(a *[]int, b *[]int) []int {
	c := [1001]int{}

	// 逐位相乘并累加
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len(*b); j++ {
			c[i+j] += (*a)[i] * (*b)[j]
		}
	}

	// 进位
	for i := 0; i < len(c); i++ {
		if c[i] > 9 {
			c[i+1] += c[i] / 10
		}
		c[i] = c[i] % 10
	}

	// 去掉后面的0
	// len := len(c) - 1
	// for c[len] == 0 && len > 0 {
	// 	len--
	// }

	return c[:500]
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
