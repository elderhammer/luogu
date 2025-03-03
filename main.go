package main

import (
	"fmt"
	"strings"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	order := []int{} // 安排顺序
	for i := 0; i < m*n; i++ {
		workpiece := 0 // 工件号
		fmt.Scan(&workpiece)
		order = append(order, workpiece-1)
	}

	workpiece_machines := [][]int{} // 每个工件每个工序对应的机器号
	for i := 0; i < n; i++ {
		step_machines := []int{} // 每个工序对应的机器号
		for j := 0; j < m; j++ {
			machine := 0 // 机器号
			fmt.Scan(&machine)
			step_machines = append(step_machines, machine-1)
		}
		workpiece_machines = append(workpiece_machines, step_machines)
	}

	max_time := 0         // 最大时间
	op_times := [][]int{} // 每个工件每个工序的耗时
	for i := 0; i < n; i++ {
		step_times := []int{} // 每个工序所需的时间
		for j := 0; j < m; j++ {
			op_time := 0 // 工序耗时
			fmt.Scan(&op_time)
			max_time += op_time
			step_times = append(step_times, op_time)
		}
		op_times = append(op_times, step_times)
	}

	// fmt.Println(m, n, order, op_times)

	machine_timelines := [][]string{} // 每台机器的时间线
	for i := 0; i < m; i++ {
		machine_timeline := []string{}
		for j := 0; j < max_time; j++ {
			machine_timeline = append(machine_timeline, "") // "" 表示还未安排，"n-m" 表示安排给了第几个工件的第几道工序
		}
		machine_timelines = append(machine_timelines, machine_timeline)
	}

	workpiece_next_moments := []int{} // 每个工件的下一个操作的开始时刻
	for i := 0; i < n; i++ {
		workpiece_next_moments = append(workpiece_next_moments, 0)
	}

	workpiece_finish_steps := []int{} // 每个工件已经完成的工序
	for i := 0; i < n; i++ {
		workpiece_finish_steps = append(workpiece_finish_steps, -1)
	}

	// 开始安排
	actual_max_time := 0
	for _, workpiece := range order {
		// 已经完成的工序
		last_step := workpiece_finish_steps[workpiece]
		// 当前要安排的工序
		current_step := last_step + 1
		// 当前工序的耗时
		current_step_time := op_times[workpiece][current_step]
		// 当前工序的机器号
		current_step_machine := workpiece_machines[workpiece][current_step]
		// 对应的机器时间线
		machine_timeline := machine_timelines[current_step_machine]

		// 当前工序的开始时刻
		current_step_start_time := workpiece_next_moments[workpiece]

		// 找空档
		for i := current_step_start_time; i < len(machine_timeline); i++ { // 暂时不用考虑边界？
			workpiece_next_monent := i + current_step_time
			if strings.Join(machine_timeline[i:workpiece_next_monent], "") == "" {
				for j := i; j < workpiece_next_monent; j++ {
					mark := fmt.Sprintf("%d-%d", workpiece+1, current_step+1)
					machine_timeline[j] = mark
				}

				// 找到并且已安排了
				workpiece_finish_steps[workpiece] = current_step
				workpiece_next_moments[workpiece] = workpiece_next_monent
				if workpiece_next_monent > actual_max_time {
					actual_max_time = workpiece_next_monent
				}
				break
			}
		}
	}

	fmt.Println(actual_max_time)

	// for _, timeline := range machine_timelines {
	// 	for _, moment := range timeline {
	// 		if moment == "" {
	// 			fmt.Print("无 ")
	// 		} else {
	// 			fmt.Printf("%s ", moment)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
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
