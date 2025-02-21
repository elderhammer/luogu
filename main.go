package main

import "fmt"

func main() {
	clue := make(map[rune][]rune)
	clue['b'] = []rune{'o', 'b'}
	clue['o'] = []rune{'y', 'b'}
	clue['y'] = []rune{' ', 'b'}
	clue['g'] = []rune{'i', 'g'}
	clue['i'] = []rune{'r', 'g'}
	clue['r'] = []rune{'l', 'g'}
	clue['l'] = []rune{' ', 'g'}

	var str string
	fmt.Scan(&str)

	var str_in_rune []rune
	for _, c := range str {
		str_in_rune = append(str_in_rune, c)
	}
	str_in_rune = append(str_in_rune, '.')

	boy_count, girl_count := 0, 0
	for i := 0; i < len(str); i++ {
		c, next := str_in_rune[i], str_in_rune[i+1]
		if c != '.' {
			info := clue[c]
			expect, kind := info[0], info[1]
			if next != expect {
				// 统计
				if kind == 'b' {
					boy_count += 1
				} else {
					girl_count += 1
				}
			}
		}
	}

	fmt.Println(boy_count)
	fmt.Println(girl_count)
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
