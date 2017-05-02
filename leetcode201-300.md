### 201. Bitwise AND of Numbers Range
	逐数相与

### 202. Happy Number
	判重复的都用hash

### 203. Remove Linked List Elements
	...

### 204. Count Primes
	素数筛法

### 205. Isomorphic Strings
	双射

### 206. Reverse Linked List
	...

### 207. Course Schedule
	拓扑排序，看看是否能删除掉所有的边

### 208. Implement Trie (Prefix Tree)
	构造字典树

### 209. Minimum Size Subarray Sum
	双指针

### 210. Course Schedule II
	和207一样

### 211. Add and Search Word - Data structure design
	字典树

### 212. Word Search II
	字典树加dfs

### 213. House Robber II
	分成搜索第一间房子和不搜索第一间房子的情况

### 214. Shortest Palindrome
	KMP算法，暂时不会，抄点网上的题解
	思路：让在前面补一些字符使得给定的字符串变成回文，观察可以发现我们需要添加多少个字符与给定字符串的前缀子串回文的长度有关．也就是说去掉其前缀的回文子串，我们只需要补充剩下的子串的逆序就行了，举个栗子：aacecaaa，其前缀的最大回文子串是aacecaa，剩下了一个a，因此我们只需要在前面加上一个a的逆序a就行了．再例如abcd，其前缀的最大回文是a，因此剩下的子串就为bcd，所以需要在前面加上bcd的逆序，就变成了dcbabcd．所以这样问题就转化为求字符串的前缀最大回文长度．

	OK，思路是这样，接下来就要看怎么实现了，一个naive的方法是先判断整个字符串是否回文，否的话再判断前n-1个子串是否回文，这样依次缩减长度，直到找到一个回文子串就是最大的前缀回文子串．这种方法简单粗暴，容易理解和实现，如果在面试中要求不是很严格的情况下说不定可以过关, 反正总比不会好点^.^．其时间复杂度是O(n!)．

	当然问题不会这么简单，接下来的才是真正有效的解法，KMP．这是一个非常高效的字符串比较算法．其原理是给定一个字符串S和Ｐ，要在Ｓ中寻找是否存在P，一般的方法是逐位比较，如果不能完全匹配，则S再回到开始位置向右移动一位，P回到０位置再开始比较．在KMP中不需要回到首部重新开始比较，借助与记录从P的开头到当前位置P中的前缀和后缀有多少位是相等的，这样当P和Ｓ比较的时候如果P[i] != S[j]了，不需要回到P[0]的位置重新比较，我们可以查看P中已经匹配过的子串中（也就是P[0, i-1]子串）前缀和后缀有多少位是相等的，然后将P的前缀和已经S[j]之前的后缀是匹配的，就可以不用回溯S了．

	所以借助与KMP记录最长前缀和后缀的方法，我们可以将原字符串翻转以后加在原字符串的后面，其最大的前缀和后缀就是前缀的最大回文长度．我们还需要在这两个字符串之间加一个冗余字符，因为形如aaaaa这种字符串如果不加一个冗余字符最大前缀和后缀会变大．

### 215. Kth Largest Element in an Array
	小顶堆

### 216. Combination Sum III
	dfs

### 217. Contains Duplicate
	hash

### 219. Contains Duplicate II
	滑动窗口

### 220. Contains Duplicate III
	set lower_bound

### 221. Maximal Square
	简单DP

### 222. Count Complete Tree Nodes
	用暴力法, recursive求会超时 O(N).   如果从某节点一直向左的高度 = 一直向右的高度, 那么以该节点为root的子树一定是complete binary tree. 而 complete binary tree的节点数,可以用公式算出 2^h - 1. 如果高度不相等, 则递归调用 return countNode(left) + countNode(right) + 1.  复杂度为O(h^2)   

### 223. Rectangle Area
	...

### 224. Basic Calculator
	使用两个栈分别记录操作数和运算符，注意运算符之间的优先关系

### 225. Implement Stack using Queues
	机智

### 226. Invert Binary Tree
	dfs

### 227. Basic Calculator II
	同224

### 228. Summary Ranges
	...

### 229. Majority Element II
	投票算法

### 230. Kth Smallest Element in a BST
	优先队列

### 231. Power of Two
	n&(n-1)

### 232. Implement Queue using Stacks
	...

### 233. Number of Digit One
	编程之美2.4
	假设N = abcdef, 如果要计算百位上的数字出现1的次数，分为如下三种情况：
	1) 如果百位上的数字为0，等于更高位数字(abc)x当前位数(100)
	2) 如果百位上的数字为1，等于更高位数字(abc)x当前位数(100)，加上等位数字(ef)+1
	3) 如果百位上的数字为2-9，等于更高位数字加1(abc+1)x当前位数(100)

### 234. Palindrome Linked List
	快慢指针，翻转链表

### 235. Lowest Common Ancestor of a Binary Search Tree
	dfs

### 236. Lowest Common Ancestor of a Binary Tree
	根据深度值求

### 237. Delete Node in a Linked List
	转化为删除下一个节点

### 238. Product of Array Except Self
	从前向后，从后向前

### 239. Sliding Window Maximum
	hash, rbegin() 返回map 的key的最大值

### 240. Search a 2D Matrix II
	选择左下或者右上角开始搜索

### 241. Different Ways to Add Parentheses
	根据运算符分为左右两部分，返回左右部分的笛卡尔积

### 242. Valid Anagram
	投票

### 257. Binary Tree Paths
	dfs

### 258. Add Digits
	神奇

### 260. Single Number III
	所有元素异或，可以得到一个非0的值。从此值中随便取一个1位，即可作为筛选的依据。所有元素会被分为两部分。两部分各自做异或操作，最后得到的就是结果。

### 263. Ugly Number
	看看能否被2、3、5整除

### 264. Ugly Number II
	设置三个指针分别指向2,3,5的倍数

### 268. Missing Number
	异或运算有如下性质: a^b^b = a，

### 273. Integer to English Words
	...

### 274. H-Index
	排序，返回a[i]>=n-i;

### 275. H-Index II
	二分

### 278. First Bad Version
	二分

### 279. Perfect Squares
	dp, 状态转移方程是:i = sqrt(n), dp[n] = min(dp[n], dp[n - i*i] + 1);

### 282. Expression Add Operators
	dfs, 需要特殊考虑的情况是乘法，需要记录前一个运算数的值。

### 283. Move Zeroes
	双指针

### 284. Peeking Iterator
	...

### 287. Find the Duplicate Number
	类似求单链表的环入口结点

### 289. Game of Life
	二次循环遍历统计neighbor的情况

### 290. Word Pattern
	双射

### 292. Nim Game
	4的倍数就赢不了

### 295. Find Median from Data Stream
	low_bound

### 297. Serialize and Deserialize Binary Tree
	dfs

### 299. Bulls and Cows
	用一个vector<vector<int>>存储出现过的未匹配的secret的数字

### 300. Longest Increasing Subsequence
	dp，dp[i]记录了长度为i的最长递增子序列的最大元素