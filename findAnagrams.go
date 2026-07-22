package hot100
/*
因为异味词和本身字符串长度相同，所以窗口长度是本身字符串长度
1. 边界检查
2. 初始化频率数组
3. 检测初始窗口
4. 滑动窗口
*/

func findAnagrams(s string, p string) (ans []int) {
	sLen, pLen := len(s),len(p)
	if sLen < pLen {
		return
	}
	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i,ch := range s[:sLen-pLen]{
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}