package innersort

//合并排序

//Merge 合并函数
func Merge(lst *List, temp []T, i1, j1, i2, j2 int, k *int) {
	var (
		i = i1
		j = i2
	)
	//若两个子序列都不空，则循环
	for i <= j1 && j <= j2 {
		if lst.Elements[i].GetKey() <= lst.Elements[j].GetKey() {
			//将较小元素存入temp[*k]
			temp[*k] = lst.Elements[i]
			*k++
			i++
		} else {
			temp[*k] = lst.Elements[j]
			*k++
			j++
		}
	}
	//将子序列1中剩余的元素存入temp
	for i <= j1 {
		temp[*k] = lst.Elements[i]
		*k++
		i++
	}
	//将自学列2中剩余的元素存入temp
	for j <= j2 {
		temp[*k] = lst.Elements[j]
		*k++
		j++
	}
}

//MergeSort 合并排序
func MergeSort(lst *List) {
	var temp []T = make([]T, lst.Size)
	//i1, j1和i2, j2分别是两个子序列的上、下界
	var i1, j1, i2, j2, i, k int
	var size = 1
	for size < lst.Size {
		i1 = 0
		k = 0
		//若i1+size<n，则说明存在两个子序列，需要两两合并
		for i1+size < lst.Size {
			//确定子序列2的下界和子序列1的上界
			i2 = i1 + size
			j1 = i2 - 1
			//设置子序列2的上界
			if i2+size-1 > lst.Size-1 {
				j2 = lst.Size - 1
			} else {
				j2 = i2 + size - 1
			}
			//合并相邻两个子序列
			Merge(lst, temp, i1, j1, i2, j2, &k)
			//确定下一次合并第一个子序列的下界
			i1 = j2 + 1
		}
		for i = 0; i < k; i++ {
			lst.Elements[i] = temp[i]
		}
		//子序列长度扩大一倍
		size *= 2
	}
}
