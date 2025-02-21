package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	// 判断数据类型
	typ := 0 // 0表示整数，1表示小数，2表示分数，3表示百分数
	var v_input []byte
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '/':
			typ = 1
		case '.':
			typ = 2
		case '%':
			typ = 3
		}
		v_input = append(v_input, input[i])
	}

	switch typ {
	case 0: // 整数
		reverse_num(v_input)
	case 1: // 分数
		if v_input[0] == '0' {
			fmt.Print(0)
		} else {
			var numerator []byte
			var denominator []byte
			is_numerator := true
			for i := 0; i < len(v_input); i++ {
				if v_input[i] == '/' {
					is_numerator = false
				} else {
					if is_numerator {
						numerator = append(numerator, v_input[i])
					} else {
						denominator = append(denominator, v_input[i])
					}
				}
			}
			reverse_num(numerator)
			fmt.Printf("%c", '/')
			reverse_num(denominator)
		}
	case 2: // 小数
		var integer []byte
		var decimal []byte
		is_integer := true
		only_zero := true
		for i := 0; i < len(v_input); i++ {
			if v_input[i] == '.' {
				is_integer = false
			} else {
				if is_integer {
					integer = append(integer, v_input[i])
				} else {
					if v_input[i] != '0' {
						only_zero = false
					}
					decimal = append(decimal, v_input[i])
				}
			}
		}
		first_num := 0
		for i := 0; i < len(decimal); i++ {
			if decimal[i] != '0' {
				first_num = i
				break
			}
		}
		reverse_num(integer)
		fmt.Printf("%c", '.')
		if only_zero {
			fmt.Printf("0")
		} else {
			reverse_num(decimal[first_num:])
		}
	case 3: // 百分数
		reverse_num(v_input[:len(v_input)-1])
		fmt.Printf("%c", '%')
	}
	fmt.Println()
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

func make_cube(n int) [][]rune {
	var cube [][]rune
	for i := 0; i < n; i++ {
		var some_row []rune
		for j := 0; j < n; j++ {
			some_row = append(some_row, 0)
		}
		cube = append(cube, some_row)
	}

	return cube
}
