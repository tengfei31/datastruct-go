package innersort

//选择排序

//SelectSort 简单选择排序
func (lst *List) SelectSort() {
	var small, i, j int
	for i = 0; i < lst.Size; i++ {
		small = i
		for j = i + 1; j < lst.Size; j++ {
			if lst.Elements[j].GetKey() < lst.Elements[small].GetKey() {
				small = j
			}
		}
		lst.Elements[i], lst.Elements[small] = lst.Elements[small], lst.Elements[i]
	}
}
