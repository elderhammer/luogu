package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var str string
	fmt.Scan(&str)

	i, times, count := 0, 1, 0
	for i < n-1 {
		step := 1
		if str[i] == 'V' && str[i+1] == 'K' {
			count += 1
			step = 2
		} else if str[i] == 'K' && str[i+1] == 'K' {
			if times > 0 {
				// str[i] = 'V'
				times -= 1
				count += 1
			}
			step = 2
		} else if str[i] == 'V' && str[i+1] == 'V' {
			if i+2 >= n { // 越界了
				if times > 0 {
					// str[i+1] = 'K'
					times -= 1
					count += 1
				}
				step = 2
			} else { // 后面还有元素
				if str[i+2] == 'K' {
					count += 1
					step = 3
				} else {
					if times > 0 {
						// str[i+1] = 'K'
						times -= 1
						count += 1
					}
					step = 2
				}
			}
		}
		i += step
	}

	fmt.Println(count)
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
