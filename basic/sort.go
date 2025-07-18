package basic

import (
	"bufio"
	"fmt"
	"os"
)

func Sort() {
	const N = 100010
	reader := bufio.NewReader(os.Stdin)
	var n int
	var nums []int = make([]int, N)
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}
	//quickSort(nums, 0, n-1)
	mergeSort(nums, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", nums[i])
	}
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	i := l - 1
	j := r + 1
	x := nums[mid]
	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}

func mergeSort(nums []int, l, r int) {
	if l == r {
		return
	}

	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	i := l
	j := mid
	tmp := make([]int, r-l+1)
	k := 0
	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = nums[j]
		k++
		j++
	}
	for i, j = l, 0; i <= r; i++ {
		nums[i] = tmp[j]
		j++
	}
}

func findK() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanln(reader, &n, &k)
	const N = 100010
	nums := make([]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}
	sortQuick(nums, 0, n-1)
	fmt.Println(nums[k-1])
}

func sortQuick(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	x := nums[mid]
	i := l - 1
	j := r + 1
	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	sortQuick(nums, l, j)
	sortQuick(nums, j+1, r)
}
