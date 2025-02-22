package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// startTime := time.Now()

	vars := make(map[byte]byte)
	var_order := []byte{}

	last_var_name := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] == ':' { // 找出变量名
			last_var_name = input[i-1]
			vars[last_var_name] = '0'
			var_order = append(var_order, last_var_name)
		} else if input[i] == ';' {
			if input[i-2] == '=' {
				vars[last_var_name] = input[i-1]
			}
		}
	}

	// fmt.Println(vars)
	for i := 0; i < len(var_order); i++ {
		var_name := var_order[i]
		for {
			value := vars[var_name]
			if '0' <= value && value <= '9' {
				fmt.Printf("%d ", value-'0'+0)
				break
			} else {
				var_name = value
			}
		}
	}

	fmt.Println()

	// endTime := time.Now()

	// elapsedTime := endTime.Sub(startTime)
	// fmt.Println("执行时间：", elapsedTime)
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
