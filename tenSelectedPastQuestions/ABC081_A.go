package tenselectedpastquestions

/*
問題文
すぬけ君は 1,2,3 の番号がついた3 つのマスからなるマス目を持っています。 
各マスには 0 か 1 が書かれており、マスi には si が書かれています。
すぬけ君は 1 が書かれたマスにビー玉を置きます。 ビー玉が置かれるマスがいくつあるか求めてください。

制約
s1, s2 ,s3 は 1 あるいは 0
*/

func ABC081_A(str string) int{
	var count int

	/* 受け取った引数をごにょごにょする */
	/* str[0] */

	for i := 0; i < 3; i++{
		if str[i] == '1'{
			count++
		}
	}
	return count
}