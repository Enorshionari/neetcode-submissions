func twoSum(nums []int, target int) []int {
    for k := 0; k < len(nums); k++ {
        for v := k+1; v < len(nums); v++ {
            if nums[k]+nums[v]==target { return []int{k,v}}
        }
    }
    return nil;
}
