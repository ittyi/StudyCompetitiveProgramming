package tenselectedpastquestions

/*
 1 以上
N 以下の整数のうち、
10 進法での各桁の和が
A 以上
B 以下であるものの総和を求めてください。

制約
*/

func calcDigitSum(n int) int {
	var sum int
	for n > 0 {
		sum = sum + (n % 10)
		n /= 10
	}
	return sum
}

func ABC083_B(n, a, b int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		/* 各桁の合計を求める */
		num := calcDigitSum(i)
		if num >= a && num <= b {
			sum += i
		}
		/* a以上b以下 */
	}

	return sum
}
