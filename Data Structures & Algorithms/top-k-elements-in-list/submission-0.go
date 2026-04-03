func getMax(dict map[int]int) int{
	max := 0
	for _,v := range dict {
		if v > max {
			max = v
		}
	}

	return max
}


func topKFrequent(nums []int, k int) []int {

	freq := make(map[int]int)
	result := []int{}
	for _,v := range nums {
		freq[v]++
	}

	for i := getMax(freq); i > 0; i-- {
		for key, value := range freq {
			if value == i {
				result = append(result, key)
				if len(result) == k {
					return result
				}
			}
		}
	}


	return result

	

}

