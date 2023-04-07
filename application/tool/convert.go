package tool

type TYPE interface {
	~string | ~int | ~int64
}

// ArrayDeDuplicate 数组去重,稳定
func ArrayDeDuplicate[T TYPE](arr *[]T) {
	sort := make([]T, 0, len(*arr))
	set := make(map[T]bool)
	for _, item := range *arr {
		set[item] = false
		sort = append(sort, item)
	}

	t := make([]T, 0, len(set))
	for _, item := range sort {
		if set[item] {
			continue
		}

		t = append(t, item)
		set[item] = true
	}
	*arr = t
}

// Array2Set 数组转集合
func Array2Set[T TYPE](arr []T) map[T]struct{} {
	res := make(map[T]struct{})
	for _, item := range arr {
		res[item] = struct{}{}
	}
	return res
}
