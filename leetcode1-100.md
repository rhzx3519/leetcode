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

### 44. Wildcard Matching 
	1) 如果p[j]=='*', 则依次将p[j+1, length]和s[i, length]、s[i+1, length]比较，如果全都失败了，则将i和j都回溯到上一个*号的位置
	2) 如果p[j]=='?'或者p[j]==p[i]，则i++,j++
	3) 否则说明当前不匹配，回溯到上一个＊号的位置

### 45. Jump Game II 
	一次遍历

### 46. Permutations 

### 47. Permutations II 
	每次交换之前对pos之后的元素进行排列，遍历的时候跳过重复的元素

### 48. Rotate Image 
	取转置矩阵，然后每行翻转

### 49. Group Anagrams 
	字符串排序，然后用map存储相同的字符串

### 50. Pow(x, n) 
	递归做，求pow(x, n/2), 注意判断n是否溢出

### 51.N-Queens
	使用一维数组A[N]存储row行上的皇后位置col,即A[row]=col，数组元素的初始值是INT_MIN。按行搜索每一列上可能摆放的位置,如果A[i] == col || |A[i]-col| == |i-row| || A[row]!=INT_MIN, 则表明位置row,col所在的行、列或者斜线上有皇后，不可摆放，否则设置A[row]=i，继续搜索下一行。

### 52. N-Queens II
	将上面的构造结果的部分改为统计数量就行

### 53. Maximum Subarray
	遍历一遍

### 54. Spiral Matrix
	设置4个指针li,lj,hi,hj代表当次循环的4个边界，依次遍历上右下左，注意遍历下左的时候要判断li!=hi和lj!=hj，去除重复遍历的情况。

### 55. Jump Game
	注意保存状态，new_pos保存了当前能走的最大距离需要和i+nums[i]取一个最大值.

### 56. Merge Intervals
	首先需要按照interval的start大小从小到大排序数组，之后遍历数组，分别比较前一个interval和后一个interval，
	1）如果前一个interval完全覆盖后一个interval，则删除后一个interval，数组大小减1
	2）如果前一个interval和后一个interval有重叠位置，则更新前一个interval的end，删除后一个interval，数组大小减1
	3）否则前一个interval和后一个interval没有交集，继续遍历下一个。


### 57. Insert Interval
	一次遍历，考虑三种情况
	1) 首先如果newInterval的end比intervals[i]的start小，说明后的intervals都不会有和newInterval有交集了，跳出循环
	2）如果newInterval的start比intervals[i]的end大，跳过
	3）剩下的情况就是newInterval和当前intervals[i]有交集，更新newInterval的start和end，重叠数量加1
	最后删除重叠的interval，插入newInterval

### 58. Length of Last Word
	...

### 59. Spiral Matrix II
	和54题一样，遍历的同时插入数字到二维数组当中

### 60. Permutation Sequence
	给定一个数n，以1开头的排列共有(n-1)!种，同理以2、3...9开头的排列各有(n-1)种，如果k的值小于等于
	(n-1)!，那个头一个数字就是1，如果k的值大于(n-1)!，那么头个数字就是(k-1)/(n-1)! + 1。
	求完第一个数字之后，k%=(n-1)!, 第二个数字的值就是(k-1)/(n-2)! + 1,
	重复上述步骤，依次求得个位数字
	注意:
	1) 构造排列数从最高位开始，当选出一个数字后，就应当把这个数字erase掉，防止后面又出现

### 61. Rotate List
	链表题注意指针

### 62. Unique Paths
	dp题，状态转移方程是:dp[i][j]表示到坐标(i,j)的路径数目, 它的状态由它左方和上方格子的路径数目决定
	dp[i][j] = dp[i-1][j] + dp[i][j-1]
	注意初始状态设置，首行与首列的所有格子路径都是惟一的。

### 63. Unique Paths II
	和上题一样，不同的是初始化状态时，碰到格子有障碍就停止，计算dp{i][j]时要判断(i,j)位置是否有障碍

### 64. Minimum Path Sum
	dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])

### 65. Valid Number
	遍历字符串，注意多种情况吧

### 66. Plus One
	...
	
### 67. Add Binary
	O(m+n)

### 68. Text Justification
	遇到长度超过就生成字符串，注意最后一行不是平均分隔，而是空一格

### 69. Sqrt(x)
	二分查找

### 70. Climbing Stairs
	第n级台阶的走法可以通过第n-1级台阶走1步或者第n-2级台阶走两步得到，d(n) = d(n-1) + d(n-2)

### 71. Simplify Path
	使用栈

### 72. Edit Distance
	dp[i][j]表示word1的前i个字符改为word2的前j个字符所需要的最少操作数
	if (c1 == c2)
		dp[i][j] = dp[i-1][j-1];
	else
		dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1;	

### 73. Set Matrix Zeroes
	时间复杂度O(mn),空间复杂度O(m+n)

### 74. Search a 2D Matrix
	二分查找

### 75. Sort Colors
	类似快排的单次遍历

### 76. Minimum Window Substring
	做法就是双指针的贪心做法，找到以i结尾的最小子串满足条件。min_len和left，right分别记录最小子串的最小值和左右位置。data存T串的信息，now存当前i到j的信息，保存的是元素的个数。
	num记录增加的元素个数是否达到条件。

### 77. Combinations
	dfs

### 78. Subsets
	dfs

### 79. Word Search
	回溯查找，在每一个位置dfs

### 80. Remove Duplicates from Sorted Array II
	跟26题差不多

### 81. Search in Rotated Sorted Array II
	先查找分割点，然后用二分查找

### 82. Remove Duplicates from Sorted List II
	略复杂，注意指针操作吧

### 83. Remove Duplicates from Sorted List
	设置前置指针，遍历链表

### 84. Largest Rectangle in Histogram
	记录当前位置向左向右延伸的最大距离

### 85. Maximal Rectangle
	将每一行看成直方图的底，上面1的数量看成直方图的高，就可以转化为84题的求法了。

### 86. Partition List
	分成两个链表之后再合并

### 87. Scramble String
	递归判断, 不用搞成二叉树

### 88. Merge Sorted Array
	从后往前插入

### 89. Gray Code
	第一步：产生 0, 1 两个字符串。
	第二步：在第一步的基础上，每一个字符串都加上0和1，但是每次只能加一个，所以得做两次。这样就变成了 00,01,11,10 （注意对称）。
	第三步：在第二步的基础上，再给每个字符串都加上0和1，同样，每次只能加一个，这样就变成了 000,001,011,010,110,111,101,100。
	好了，这样就把3位元格雷码生成好了。
	如果要生成4位元格雷码，我们只需要在3位元格雷码上再加一层0,1就可以了： 0000,0001,0011,0010,0110,0111,0101,0100,1100,1101,1110,1010,0111,1001,1000.
	也就是说，n位元格雷码是基于n-1位元格雷码产生的。

### 90. Subsets II
	可以发现从S=[1,2]变化到S=[1,2,2]时，多出来的有两个子集[2,2]和[1,2,2]，这两个子集，其实就是 [2], [1,2]末尾都加上2 而产生。而[2], [1,2] 这两个子集实际上是 S=[1,2]的解到 S=[1]的解 新添加的部分。
	因此，若S中有重复元素，可以先排序；遍历过程中如果发现当前元素S[i] 和 S[i-1] 相同，那么不同于原有思路中“将当前res中所有自己拷贝一份再在末尾添加S[i]”的做法，我们只将res中上一次添加进来的子集拷贝一份，末尾添加S[i]。

### 91. Decode Ways
	带条件的斐波那契数列,dp[i]表示前i个字符的编码数量，如果s[i]在1到26之间，dp[i] += dp[i-1], 如果s[i]和s[i-1]组成的数字在1到26之间，那么dp[i] += dp[i-2]

### 92. Reverse Linked List II
	将要反转的结点存到另一个链表中，最后合并

### 93. Restore IP Addresses
	dfs

### 94. Binary Tree Inorder Traversal
	...

### 95. Unique Binary Search Trees II
	dfs

### 96. Unique Binary Search Trees
	这道题要求可行的二叉查找树的数量，其实二叉查找树可以任意取根，只要满足中序遍历有序的要求就可以。从处理子问题的角度来看，选取一个结点为根，就把结点 切成左右子树，以这个结点为根的可行二叉树数量就是左右子树可行二叉树数量的乘积，所以总的数量是将以所有结点为根的可行结果累加起来。
	dp[i] = sigma{dp[j-1]*dp[i-j]}. 1<=j<=i
	dp[i]表示节点数为i的二叉排序树有多少种不同的生成方式。

### 97. Interleaving String
	res[i][j] 代表s1中前i个字符与s2中前j个字母是否可以交叉匹配成为s3中前i+j个字符。
	公式为：
	if s3[i + j] = s1[i] = s2[j]. res[i][j] = res[i - 1][j] || res[i][j - 1]
	if s3[i + j] = s1[i] res[i][j] = res[i - 1][j]
	if s3[i + j] = s2[j] res[i][j] = res[i][j - 1]
	else s3[i + j] = false;
	初始条件是res[0][0] = true. 即s1前0个字符加s2前0个字符，可以拼凑成s3前0个字符。

### 98. Validate Binary Search Tree
	dfs

### 99. Recover Binary Search Tree
	中序遍历，可以用递归和非递归的方法

### 100. Same Tree
	dfs