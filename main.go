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
	var na int
	var nb int
	fmt.Scan(&n, &na, &nb)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a_input := scanner.Text()
	a_attacks := []int{}
	for _, attack := range strings.Fields(a_input) {
		a, _ := strconv.Atoi(attack)
		a_attacks = append(a_attacks, a)
	}

	scanner.Scan()
	b_input := scanner.Text()
	b_attacks := []int{}
	for _, attack := range strings.Fields(b_input) {
		a, _ := strconv.Atoi(attack)
		b_attacks = append(b_attacks, a)
	}

	// fmt.Println(a_attacks, b_attacks)

	cube := [5][5]int{} // 1表示x胜（y输），0表示平手，-1表示x输（y赢）
	cube[0][0] = 0
	cube[0][1] = -1
	cube[0][2] = 1
	cube[0][3] = 1
	cube[0][4] = -1

	cube[1][0] = 1
	cube[1][1] = 0
	cube[1][2] = -1
	cube[1][3] = 1
	cube[1][4] = -1

	cube[2][0] = -1
	cube[2][1] = 1
	cube[2][2] = 0
	cube[2][3] = -1
	cube[2][4] = 1

	cube[3][0] = -1
	cube[3][1] = -1
	cube[3][2] = 1
	cube[3][3] = 0
	cube[3][4] = 1

	cube[4][0] = 1
	cube[4][1] = 1
	cube[4][2] = -1
	cube[4][3] = -1
	cube[4][4] = 0

	a_score, b_score := 0, 0
	for i := 0; i < n; i++ {
		result := cube[a_attacks[i%len(a_attacks)]][b_attacks[i%len(b_attacks)]]
		if result == 1 {
			a_score += 1
		} else if result == -1 {
			b_score += 1
		}
	}
	fmt.Printf("%d %d\n", a_score, b_score)
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
