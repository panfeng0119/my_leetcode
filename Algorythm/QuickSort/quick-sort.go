package quicksort

import (
	"fmt"
	"math/rand"
)
// 递归思想
func QuickSortDG(values []int) {
	if len(values) <= 1 {
		return
	}
	pivot, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		fmt.Println("array:", values, ", pivot:", pivot)
		if values[i] > pivot {
			fmt.Println(values[i], pivot)
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = pivot
	QuickSortDG(values[:head])
	QuickSortDG(values[head+1:])
}

func quickSort(values []int, left, right int) {
	// 基准
	pivot := values[left]
	p := left           // 基准的位置
	i, j := left, right // 左右指针
	//fmt.Println("pivot:", pivot, ", left(i):", left, ", right(j):", right, ", arr:", values)
	for i <= j {
		//fmt.Println("    1: i:", i, ", j:", j, ", p:", p)
		for j >= p && values[j] >= pivot {
			//fmt.Printf("\t\tj--, j(%v) >= p(%v),  values[j](%v) >= pivot(%v) \n", j, p, values[j], pivot)
			j--
		}
		//fmt.Println("    2: i:", i, ", j:", j, ", p:", p)
		if j >= p {
			//fmt.Println("    values[p]", values[p], " => ", values[j], ", p:", p, " => ", j)
			values[p] = values[j]
			p = j
		}
		for i <= p && values[i] <= pivot {
			//fmt.Printf("\t\ti++, i(%v) <= p(%v),  values[i](%v) <= pivot(%v) \n", i, p, values[i], pivot)
			i++
		}
		//fmt.Println("    3: i:", i, ", j:", j, ", p:", p)
		if i <= p {
			//fmt.Println("    values[p]", values[p], " => ", values[i], ", p:", p, " => ", i)
			values[p] = values[i]
			p = i
		}
	}
	values[p] = pivot
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

// 分治思想
func QuickSortFZ(values []int) {
	if len(values) < 2 {
		return
	}
	quickSort(values, 0, len(values)-1)
}

/* ---------------------------------------------------------------------
**/
// 默写
func qs(arr []int) {
	if len(arr) < 2 {
		return
	}
	qst(arr, 0, len(arr)-1)
}

func qst(arr []int, left, right int) {
	if left >= right {
		return
	}
	// 基准
	p := left + rand.Intn(right-left+1) // 当前空位，将匹配的值存到这个位置

	pivot := arr[p] // 把这个位置上的基准提取出来
	i, j := left, right
	// 每一次调用相当于 把当前基准放到适当的位置上
	for i <= j {
		// 从右开始
		for p <= j && arr[j] >= pivot {
			j--
		}
		if j > p {
			arr[p] = arr[j] // 把位置交出来，赋值给需要移动的值
			p = j           // 这个位置就没用了
		}
		for p >= i && arr[i] <= pivot {
			i++
		}
		if i < p {
			arr[p] = arr[i]
			p = i
		}
	}

	arr[p] = pivot      // 放回基准
	qst(arr, left, p-1) // p 已经找到指定位置，无需再计算
	qst(arr, p+1, right)
}

// 重复元素排序
func qs2(arr []int) {
	if len(arr) < 2 {
		return
	}
	qSort(arr, 0, len(arr)-1)
}

func qSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	// 随机基准
	p := left + rand.Intn(right-left+1)
	arr[p], arr[left] = arr[left], arr[p]

	//fmt.Println(p, arr[p], arr[left])
	//fmt.Println(arr)
	pivot := arr[left] // 把这个位置上的基准提取出来
	i, j := left, right
	// 每一次调用相当于 把当前基准放到适当的位置上
	for i <= j {
		//fmt.Println("i,j:", i, j)
		for i <= j && arr[j] > pivot { // 大的都在右边
			//fmt.Println(j)
			j--
		}

		for i <= j && arr[i] <= pivot { // 小的都在左边
			//fmt.Println(i)
			i++
		}
		// 交换
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[j] = pivot        // 放回基准
	qSort(arr, left, p-1) // p 已经找到指定位置，无需再计算
	qSort(arr, p+1, right)
}
