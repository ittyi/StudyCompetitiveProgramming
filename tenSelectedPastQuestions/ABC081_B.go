package tenselectedpastquestions

import "fmt"

/*
問題文 Shift Only
黒板に 
N 個の正の整数 
A 
1
​
 ,...,A 
N
​
  が書かれています．

すぬけ君は，黒板に書かれている整数がすべて偶数であるとき，次の操作を行うことができます．

黒板に書かれている整数すべてを，
2 で割ったものに置き換える．
すぬけ君は最大で何回操作を行うことができるかを求めてください．

*/

func ABC081_B() int{

	var n int
    fmt.Scan(&n)

	var a []int = make([]int, n, 200)
	
	for i := 0; i < n; i++{
		fmt.Scan(&a[i])
	}

	count := 0
	for {
		for i := 0; i < n; i++ {
			if a[i] % 2 != 0 {
				return count
			}

			if a[i] % 2 == 0 {
				a[i] = a[i] / 2
			}
		}
		count++
	}


	// count := 0
	// index := 0
	// for  {
	// 	// 無限にループにしたので、このまま行くと
	// 	if index == n {
	// 		index = 0
	// 		count++
	// 	}

	// 	if a[index] % 2 != 0 {
	// 		return count
	// 	}

	// 	a[index] = a[index] / 2
	// 	index++
	// }

	return count
}