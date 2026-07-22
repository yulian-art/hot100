package hot100

/*
方法一：枚举法
方法二：前缀和 + 哈希表优化

*/
func subarraySum(nums []int, k int) int {
    result := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += nums[j]
            if sum == k {
                result++
            }
        }
    }
    return result
}