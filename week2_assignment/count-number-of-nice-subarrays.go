package week2_assignment

func numberOfSubarrays(nums []int, k int) int {
	// 1. 生成前缀和数组
	preSum := make([]int, 0, 0)
	preSum = append(preSum, 0)
	for index,value := range nums {
		preSum = append(preSum, value%2 + preSum[index])
	}

	// 2. 开散列，记录每个前缀和出现的次数
	sumMap := make(map[int]int, 0)
	for _,value := range preSum {
		sumMap[value] ++
	}

	// 3. 遍历前缀和数组，记录符合条件的结果个数
	var ans int
	for _,value := range preSum {
		if _, ok := sumMap[value-k]; ok {
			ans += sumMap[value-k]
		}
	}

	return ans
}