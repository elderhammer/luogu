package main

import (
	"fmt"
)

func main() {
	var (
		n int
		m int
	)
	fmt.Scan(&n, &m)
	nums := []int{}
	for i := 0; i < m; i++ {
		num := 0
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	// quick_sort(&nums, 0, len(nums)-1)
	quick_sort_iter(&nums)

	fmt.Println(nums)
}

func quick_sort_iter(nums *[]int) {
	if len(*nums) == 1 {
		return
	}
	pqs := []int{0, len(*nums) - 1}
	pointer := 0
	for {
		if pointer >= len(pqs) {
			break
		}

		p, q := pqs[pointer], pqs[pointer+1]
		pivot := partition(nums, p, q)
		if p < pivot-1 {
			pqs = append(pqs, []int{p, pivot - 1}...)
		}
		if pivot+1 < q {
			pqs = append(pqs, []int{pivot + 1, q}...)
		}

		pointer += 2
	}
}

func quick_sort(nums *[]int, p int, q int) {
	if p < q {
		pivot_pos := partition(nums, p, q)
		quick_sort(nums, p, pivot_pos-1)
		quick_sort(nums, pivot_pos+1, q)
	}
}

// 默认选择第一个元素作为中枢
// 完成分区后，返回新的中枢
func partition(nums *[]int, p int, q int) int {
	pivot := (*nums)[p] // 不单独拿出来的话，会有数据丢失
	for p < q {
		// 既然选择了 p 作为中枢（p 的位置空出来了），那就从 q 开始比较
		for p < q && pivot <= (*nums)[q] {
			q--
		}
		// 以上循环退出了，说明发现了一个比中枢小的值，甚至全部都小于中枢
		(*nums)[p] = (*nums)[q] // 交换后，q 的位置空出来了
		for p < q && (*nums)[p] <= pivot {
			p++
		}
		// 以上循环退出了，说明发现了一个比中枢大的值，甚至全部都大于中枢
		(*nums)[q] = (*nums)[p] // 交换后，p 的位置空出来了
	}
	(*nums)[p] = pivot // 将中枢放到 p
	return p           // 返回新中枢
}

// 注意，[start, end] 是闭区间
func radix_sort_words(words *[]string, start int, end int, k int) {
	if start == end {
		return
	}

	/* 计数排序 */
	c := [26]int{}

	// 计数
	for i := start; i <= end; i++ {
		kth := "a"[0] // 越界了就当作"a"
		if k <= len((*words)[i])-1 {
			kth = (*words)[i][k] // 第k个关键字
		}
		idx := kth - 'a' // 字母转整数索引
		c[idx] += 1
	}

	// 前缀和
	for i := 0; i < 26-1; i++ {
		c[i+1] += c[i]
	}

	// 辅助数组
	tmp := []string{}
	for i := start; i <= end; i++ {
		tmp = append(tmp, "")
	}
	// 排序
	for i := end; i >= start; i-- { // 逆序保证稳定性
		kth := "a"[0] // 越界了就当作'a'
		if k <= len((*words)[i])-1 {
			kth = (*words)[i][k] // 第k个关键字
		}
		idx := kth - 'a'
		tmp[c[idx]-1] = (*words)[i] // 注意前缀和是从1开始算的，对应到索引要减1
		c[idx] -= 1
	}
	// 复原
	for i := start; i <= end; i++ {
		(*words)[i] = tmp[i-start]
	}

	// 接着处理第 k+1 个关键字
	count, new_start := 0, 0
	var last_alphabet byte = 0
	for i := start; i <= end; i++ {
		if k > len((*words)[i])-1 {
			continue
		} else if (*words)[i][k] != last_alphabet {
			// 递归处理
			if count > 1 { // 关键字数量只有1的跳过
				radix_sort_words(words, new_start, i-1, k+1)
			}
			// 清理状态
			count, new_start, last_alphabet = 1, i, (*words)[i][k]
		} else {
			count++
		}
	}
}

func counting_sort(nums *[]int) []int {
	c := [100]int{}
	// 计数
	b := []int{}
	for _, num := range *nums {
		c[num] += 1
		b = append(b, 0)
	}
	// 计算前缀和
	for i := 0; i < len(c)-1; i++ {
		c[i+1] += c[i]
	}
	// 前缀和表示相同元素中，排名最后（还未就绪的元素）的位置，即从大到小
	// 所以要逆序重排
	for i := len(*nums) - 1; i >= 0; i-- {
		num := (*nums)[i]
		pos := c[num] - 1
		b[pos] = (*nums)[i]
		c[(*nums)[i]] -= 1
	}
	return b
}

func insertion(nums *[]int) {
	for i := 1; i < len(*nums); i++ {
		new_card := (*nums)[i] // 新牌
		j := i - 1
		for j >= 0 && (*nums)[j] > new_card { // 直到第一张牌
			(*nums)[j+1] = (*nums)[j] // 集体后移，保证了稳定性
			j--
		}
		(*nums)[j+1] = new_card // 按牌面放置
	}
}

func bubble(nums *[]int) {
	flag := true // 交换过则为 true
	for flag {
		flag = false
		for i := 0; i < len(*nums)-1; i++ {
			if (*nums)[i] > (*nums)[i+1] {
				tmp := (*nums)[i]
				(*nums)[i] = (*nums)[i+1]
				(*nums)[i+1] = tmp
				flag = true
			}
		}
	}
}

func selection(nums *[]int) {
	len := len(*nums)
	for i := 0; i < len-1; i++ {
		ith := i
		for j := i + 1; j < len; j++ {
			if (*nums)[j] < (*nums)[ith] {
				ith = j
			}
		}
		tmp := (*nums)[i]
		(*nums)[i] = (*nums)[ith]
		(*nums)[ith] = tmp
	}
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
