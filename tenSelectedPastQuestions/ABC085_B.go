/*
数字受け取る
ソートする
forで数字の個数分繰り返し
i番目とi-1番目が同じだったらカウントしない
同じじゃなかったらカウント
*/
/*
連想配列使う
なんかmap使って数字いれたら、良いはず。
*/

package tenselectedpastquestions

func ABC085_B(n int, d []int) int {
	noDuplication := map[int]int{}
	for i := 0; i < n; i++ {
		noDuplication[d[i]] = 0
	}

	return len(noDuplication)
}
