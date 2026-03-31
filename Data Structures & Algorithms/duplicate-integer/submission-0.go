func hasDuplicate(nums []int) bool {

    no_dup := make(map[int]int)

    for _,v := range nums {
        no_dup[v] ++
    }

    return len(nums) != len(no_dup)
    
}
