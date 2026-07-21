package hot100
/*
按顺序枚举集合中的每个数x，应用跳过规则（检测x-1是否存在
如果不在就不跳过，x成为连续序列的第一个数，然后进入内层循环，不断查找，直到x不存在，记录长度
如果x-1在集合中，就跳过，因为x-1已经存在，以x开头的序列肯定比x-1短，不进入内层循环
*/
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	
	longestStreak :=  0
	for num := range numSet {
		if !numSet[num-1]{
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++

			}
			if longestStreak < currentStreak{
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}