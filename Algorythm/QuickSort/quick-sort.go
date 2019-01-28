package quicksort

import "fmt"

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	pivot, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		fmt.Println("values", values, "pivot", pivot)
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
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}

func quickSort(values []int, left, right int) {
	// 基准
	pivot := values[left]
	p := left           // 基准的位置
	i, j := left, right // 左右指针
	for i <= j {
		for j >= p && values[j] >= pivot {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for values[i] <= pivot && i <= p {
			i++
		}
		if i <= p {
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

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort(values, 0, len(values)-1)
}
