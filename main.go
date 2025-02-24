package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	notp := [10001]int{} // 1表示不是素数，0表示素数
	notp[0], notp[1] = 1, 1
	primes := []int{}
	for i := 2; i <= 10000; i++ {
		if notp[i] == 0 {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > 10000 { // 越界
				break
			}
			notp[i*p] = 1
			if i%p == 0 {
				break
			}
		}
	}

	for i := 1; i <= (n-2)/2; i++ {
		c := 2*i + 2
		for j := 2; j <= 10000 && c-j > 0; j++ {
			if notp[j] == 0 && notp[c-j] == 0 {
				fmt.Printf("%d=%d+%d\n", c, j, c-j)
				break
			}
		}
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
