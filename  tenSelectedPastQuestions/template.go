package main

import "fmt"

/*
問題文 Product
シカのAtCoDeerくんは二つの正整数 
a,b を見つけました。 
a と 
b の積が偶数か奇数か判定してください。

制約
1 ≤ a,b ≤ 10000
a,b は整数

偶数Even
奇数odd
*/

func ABC086_A(var a, b int) string{
	var a, b int
	var str string
	
	if a*b % 2 == 0{
		str = "Even"
	}else{
		str = "Odd"
	}

	return str
}