func twoSum(nums []int, target int) []int {

	result := []int{}

	var i,j = 0,0
	for i < len(nums)-1 {
		 j = i+1
		for j <= len(nums)-1 {

			if (nums[i] + nums[j]) == target {
			result = append(result,i)
			result = append(result,j)
			return result
		}
		j++


		}

		i++
	}

	return result
    
}
