package hot100

func groupAnagrams(strs []string) [][]string {
    mp := map[string][]string{} //创建一个空的哈希映射
    for _,str := range strs {
        s := []byte(str) //字符串在go中只读，不能修改里面的字符，所以转换成字节切片
		//将eat底层字节复制一份，生成一个新的字节切片[101 97 116]
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j]})
		//如果i在j前面，则返回true
		// 对s进行原地升序排序
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str) 
        
    }
    
    ans := make([][]string, 0, len(mp))
	for _, v := range mp{
		ans = append(ans, v)
	}
    return ans
}