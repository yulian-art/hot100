package hot100
/*
排序+双指针
排序保证三元组按非递减顺序出现，避免重复排列，省去去重步骤
也方便后续双指针移动，排序后固定a，随着b增大，c必然减小

然后遍历每个元素，作为三元组的第一个数
注意去重，跳过与上一个相同的元素，避免重复三元组

确定第三指针指向最右端，第二个指针开始从左往右开始扫描


*/
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)

	for first := 0 ; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1]{
			continue
		}
		third := n-1
		target := -1 * nums[first]
		//固定a和c，计算b+c=-a,b是递增的，所以要保证从b+c<-a开始遍历扫描，如果太大了就要右移c减小和
		for second := first + 1; second<n; second++ {
			// 需要和上一次枚举的数不同
			if second>first+1 &&  nums[second] == nums[second-1]{
				continue
			}
			for second < third && nums[second] + nums[third] > target {
				third --
			}
			// 如果指针重合，随着b的增加，就不会有满足a+b+c=0的c了，直接退出循环
			if second == third{
				break
			}
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}
	}
	return ans
}