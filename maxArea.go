package hot100
/*
双指针法
左右指针分别指向数组两端，当前容量=最小高度x距离
每次移动对应高度较小的指针，因为如果移动较大指针，容量只会减小
所有计算容量中的最大值为最终结果

*/
func maxArea(height []int) int {
	result := 0
	left, right := 0,len(height)-1
	for left < right {
		area := (right - left) * min(height[left],height[right])
		if height[left] < height[right]{
			left ++
		}else{
			right --
		}
		if result<area{
			result=area
		}
	}
	
	
	return result
}