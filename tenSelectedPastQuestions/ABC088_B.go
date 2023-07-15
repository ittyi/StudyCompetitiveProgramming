package tenselectedpastquestions

import (
	"sort"
)

/*
    // ソートする配列
    list := []int{12, 1, 9, 2, 8, 3, 22, 4, 8}

    // 配列をクイックソートする(昇順)
    sort.Slice(list, func(i, j int) bool {
        return list[i] < list[j]
    })
s := []int{4, -1, 12, -26, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
*/

func ABC088_B(n int, a []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	alice := 0
	bob := 0
	for i, v := range a {
		if i%2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}

	return alice - bob
}
