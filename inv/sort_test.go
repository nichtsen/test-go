package inv

func kFrequenct(n []int, k int) []int {
	// number -> count
	count := make(map[int]int)
	max := 0
	for _, num := range n {
		count[num]++
		if count[num] > max {
			max = count[num]
		}
	}
	// count -> []num
	buckets := make([][]int, max+1)
	for num, cnt := range count {
		buckets[cnt] = append(buckets[cnt], num)
	}
	res := make([]int, 0)
	for i:=max;len(res) < k&&i>0;i-- {
		res = append(res, buckets[i]...)
	}
	return res
}
		
		
