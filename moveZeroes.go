package hot100
/*
用双指针，左指针指向下一个放非零数的位置，其左边全都是非零数
右指针负责从做往右扫描非零数，右指针发现了非零数就把它扔给左指针，然后左指针往右挪一格
我趣了这么男的题是简单档
*/
func moveZeros(nums []int) {
	left, right,n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left],nums[right] = nums[right],nums[left]
			left++
		}
		right++
	}
}