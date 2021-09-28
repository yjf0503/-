package week1_homework

//加一
func plusOne(digits []int) []int {
	digitsLen := len(digits)
	//进位标识符
	var jinwei bool

	for i := digitsLen - 1; i >= 0; i-- {
		//如果是最后一位或者需要进位，就需要更新
		if jinwei == true || i == digitsLen - 1 {
			//如果加一之后大于9，就需要进位，该位置为0
			if digits[i] + 1 > 9{
				jinwei = true
				digits[i] = 0

				//如果到了数组首位，需要进位，就生成新数组
				if i == 0 {
					temp := []int{1}
					temp = append(temp, digits...)
					digits = temp
					break
				}
			} else {
				//如果更新之后不需要进位了，退出循环，流程结束
				digits[i] += 1
				break
			}
		}
	}
	return digits
}
