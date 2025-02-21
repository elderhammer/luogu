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
	fmt.Scan(&n)

	var str string
	fmt.Scan(&str)

	var answers []string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			input := scanner.Text()
			params := strings.Fields(input)
			switch params[0] {
			case "1":
				str = str + params[1]
				answers = append(answers, str)
			case "2":
				start, _ := strconv.Atoi(params[1])
				length, _ := strconv.Atoi(params[2])
				end := start + length

				if end >= len(str) {
					str = str[start:]
				} else {
					str = str[start:end]
				}
				answers = append(answers, str)
			case "3":
				start, _ := strconv.Atoi(params[1])
				if start >= len(str) {
					str = str + params[2]
				} else {
					str = str[0:start] + params[2] + str[start:]
				}
				answers = append(answers, str)
			case "4":
				find := -1
				target_str := params[1]
				target_len := len(target_str)
				str_len := len(str)

				if str_len < target_len {
					find = -1
				} else {
					for i := 0; len(str)-i >= target_len; i++ {
						end := i + target_len
						if end >= str_len {
							if str[i:] == target_str {
								find = i
								break
							}
						} else {
							if str[i:end] == target_str {
								find = i
								break
							}
						}
					}
				}

				answers = append(answers, strconv.Itoa(find))
			}
		}
	}

	for _, str := range answers {
		fmt.Println(str)
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
