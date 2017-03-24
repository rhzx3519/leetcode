### 5. Longest Palindromic Substring
	这里使用menecher方法，就是动态规划，首先在原先的字符串之间插入'#, 这样可以统一处理奇数串和偶数串, 
使用两个变量纪录状态, far_pos表示最长回文字符串的最大边界距离，ci表示最长回文字符串的中心位置, 状态数据dp[i]
表示i位置的回文串半径 j = min(far_pos - i + 1, dp[2*ci-i])


### 8. String to Integer (atoi)
	字符串前置空格先去除，然后考虑＋－号，之后考虑字符是否合法，最后考虑int越界的问题

### 9. Palindrome Number
	负数直接返回false，比较前后是否相等就行	

### 10. Regular Expression Matching 
	配置i, j分别指向两个字符串，如果s[i]==p[j], 或者p[j] == '.', 比较下一个字符，碰到star的时候要入栈

### 11. Container With Most Water
	二分查找

### 12. Integer to Roman 
	配置每一位数字对应的roman数字，然后转换就行

### 13. Roman to Integer 
	配置每个roman数字单位对应的阿拉伯数字，依次遍历字符串，如果s[i] < s[i+1]，则减去s[i], 否则加上s[i]，
注意最后一个数字之后已经没有数字，所以最后加上最后的数字

### 14. Longest Common Prefix 
	取集合中第一个字符串当作最长前缀，依次与剩下的字符串比较，每次比较完更新前缀长度size

### 15. 3Sum 
	首先排序，第一个循环按从小到大，内层循环使用二分查找，上下限分别是i+1和数组的最大索引, 注意内外层循环中碰到重复的
值都要跳过

### 16. 3Sum Closest
	和上一题一样

### 17. Letter Combinations of a Phone Number 
	dfs

### 18. 4Sum 
	i = 0, j = i + 1, 其他和求三个数和一样

### 19. Remove Nth Node From End of List 
	遍历两遍，第一次遍历时保存前指针就好

### 20. Valid Parentheses 
	...

### 21. Merge Two Sorted Lists 
	... 

### 22. Generate Parentheses 
	从生成最内层括号开始考虑，dfs

### 23. Merge k Sorted Lists 
	使用小顶堆，将各个结点依次放入堆中，每次从堆中取出一个结点插入到链表中，python的优先队列需要自己构造，
比较麻烦，以后用到堆的题目还是用c++好

### 24. Swap Nodes in Pairs 
	链表题都是一个套路

### 25. Reverse Nodes in k-Group 
	注意指针操作吧

### 26. Remove Duplicates from Sorted Array 
	k纪录不重复序列的最后位置，i位置表示新的不重复元素	，j用于跳过重复元素
	[原题](https://leetcode.com/problems/remove-element/#/description)

### 27. Remove Element
	跟上一题差不多

### 28. Implement strStr() 
	暴力搜索 O(n^2)的时间复杂度

### 29. Divide Two Integers 
	除数每次左移一位，知道比被除数大为止，结果就是除数增大的倍数

### 30. Substring with Concatenation of All Words 
	统计words中每个单词出现的次数，遍历s，截取字符串看看是否符合哈希的值

### 31. Next Permutation 
	从后往前寻找第一个比后面元素小的位置，交换，后面排序(python 这个不能部分排序啊, 真麻烦，list当入参时只能在原地址上修改，重新赋上新的地址是无效的)

### 32. Longest Valid Parentheses 
	这题参考了晚上的代码，用一个字典存储‘（’出现的位置，key是‘（’的个数，当碰到‘）’时，如果存在对应cnt的值，则删除该键值对，cnt--, 如果不存在对应cnt的值，设置对应的键值，否则更新最长值

### 33. Search in Rotated Sorted Array 
	现寻找旋转位置，再用二分查找

### 34. Search for a Range 
	二分查找

### 35. Search Insert Position 
	二分

### 36. Valid Sudoku 
	python这只能用while循环啊

### 37. Sudoku Solver 
	回溯查找，标记填过的格子，递归终止的条件是所有的空格子都填充完毕
	特别技巧：
	1. 使用^=标记某行或者某列出现过的数字,
	 row[i] ^= 1<<d-1, 
	 if row[i]&1<<d-1 == True :
	 	d数字在i行出现过
	2. ns = sqrt(n), 使用block[i/ns*ns + j/ns] 标记各个block出现过的数字
	3. pos = [], 存储空格子的位置, 坐标转换为i*n+j, 当pos为空时，递归终止

### 38. Count and Say 
	递归做呗

### 39. Combination Sum 
	dfs，递归终止条件是遍历完所有元素

### 40. Combination Sum II 
	dfs, 注意跳过相同的元素

### 41. First Missing Positive 
	将正数放到对应的位置上就行, 正常的样子应该是[1,2,3,4,5],
	遍历时发现nums[i] != i + 1时，交换nums[i]和nums[nums[i]-1]

### 42. Trapping Rain Water 
	设置左右指针l、r，如果l位置的高度较低，则判断l右侧的高度，如果高度下降，则代表能容纳水，l指针右移，
	知道遇到高度和l位置高度相等或大于时停止，此时重新判断l、r位置高度； 右侧同理
