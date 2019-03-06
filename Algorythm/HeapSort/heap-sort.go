package heapsort

import "fmt"

func HeapSort(arr []int) {
	// 1. 初始化: 构建（初始状态不一定是大（小）顶堆，）
	fmt.Println(arr)
	for i := len(arr)/2 - 1; i >= 0; i-- {
		//从第一个非叶子结点从下至上，从右至左调整结构
		adjustHeap(arr, i, len(arr))
		fmt.Println("in:",arr)
	}
	fmt.Println(arr)

	// 2. 调整堆顶并重新构建
	for j := len(arr) - 1; j >= 0; j-- {
		// 交换
		arr[0], arr[j] = arr[j], arr[0]
		// 调整
		adjustHeap(arr, 0, j)
	}

}

// 调整结构
// @param a 需要调整的数组左边界
// @param b 需要调整的数组右边界
func adjustHeap(arr []int, a, b int) {
	tmp := arr[a]
	// 从i的左孩子节点开始遍历，每次比较节点和左右孩子的大小
	for k := a*2 + 1; k < b; k = k*2 + 1 {
		// 左右孩子比较，如果右孩子大，则转向右孩子
		if k+1 < b && arr[k] < arr[k+1] {
			k++
		}
		if arr[k] > tmp {
			arr[a] = arr[k]
			a = k
		} else {
			break
		}
	}
	arr[a] = tmp
}
