func twoSum(nums []int, target int) []int {
    knownNumbers := make(map[int]int) 
    for i := 0; i < len(nums); i++ {
        secondNum := target - nums[i]
        j, ok := knownNumbers[secondNum] 
        if ok {
            return []int{j,i}
        } else { knownNumbers[nums[i]]=i}
    }
    return nil;
}
