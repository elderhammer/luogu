package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var q int
	fmt.Scan(&q)

	var str string
	fmt.Scanln(&str)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		scanner.Scan()
		input := scanner.Text()
		ops := strings.Fields(input)
		switch ops[0] {
		case "1": // 后接插入
			str = str + ops[1]
			fmt.Printf("%s\n", str)
		case "2":
			start, _ := strconv.Atoi(ops[1])
			length, _ := strconv.Atoi(ops[2])
			end := start + length
			if end >= len(str) {
				str = str[start:]
			} else {
				str = str[start:end]
			}
			fmt.Printf("%s\n", str)
		case "3":
			start, _ := strconv.Atoi(ops[1])
			if start >= len(str) {
				str = str + ops[2]
			} else {
				str = str[:start] + ops[2] + str[start:]
			}
			fmt.Printf("%s\n", str)
		case "4":
			ans := -1
			target_str := ops[1]
			if len(target_str) > len(str) {
				ans = -1
			} else if len(target_str) == len(str) {
				if target_str == str {
					ans = 0
				} else {
					ans = -1
				}
			} else {
				target_len := len(target_str)
				for i := 0; i < len(str); i++ {
					if i+target_len >= len(str) {
						if str[i:] == target_str {
							ans = i
							break
						}
					} else {
						if str[i:i+target_len] == target_str {
							ans = i
							break
						}
					}
				}
			}
			fmt.Printf("%d\n", ans)
		}
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
