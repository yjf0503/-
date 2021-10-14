package week2_assignment

func subarraySum(nums []int, k int) int {
	var preSum, ans int
	sumMap := make(map[int]int)
	sumMap[0] = 1
	for _,value := range nums {
		preSum += value
		if _,ok := sumMap[preSum - k]; ok {
			ans += sumMap[preSum - k]
		}
		sumMap[preSum]++
	}

	return ans
}