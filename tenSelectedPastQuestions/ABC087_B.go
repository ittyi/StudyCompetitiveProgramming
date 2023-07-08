package tenselectedpastquestions

func ABC087_B(A, B, C, X int) int{
	count := 0
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				if (500 * i + 100 * j + 50 * k) == X {
					count++
				}
			}
		}
	}

	return count
}