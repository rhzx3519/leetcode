### 303. Range Sum Query - Immutable
	解法1：经典的线段树查询，但是速度还是慢
	解法2：构造一个累加数组

### 304. Range Sum Query 2D - Immutable
	前缀和

### 306. Additive Number
	字符串转int的题目，显然要用python写比较快
	本题学习了两个Python内嵌的函数，一个是itertools.combination(list, num)。可以列出list中取num个数字的所有combination而且没有重复。
	另外一个是 string1.startswith(string2, beg, end)。这个函数可以检查string1的从beg开始长度为end的sub string是不是以string2开头。

### 307. Range Sum Query - Mutable
	线段树

### 309. Best Time to Buy and Sell Stock with Cooldown
	状态的跳转依旧是时间之间的跳转，即第i天的收益情况依赖于第i-1天的收益情况。不过现在需要三个状态，即buy，sell，cooldown，我们记录第i-1天的这三个状态的收益情况是last_ buy,last_sell,last_cooldown。那么第i天的这三个收益情况的依赖关系是： 
	1. buy=max(last_buy, last_cooldown - price[i]) 
	2. sell = max(last_sell, last_buy + price[i]) 
	3. cooldown = max(cooldown, last_sell); 
	然后就是每天计算完之后将该天的这三个状态值赋值last。 

### 310. Minimum Height Trees
	依次删除度为1的叶子节点，直到最后剩下1个或者两个结点为止

### 313. Super Ugly Number
	和丑数一样的吧

### 312. Burst Balloons
	记忆化搜索+动态规划，dp[i][j]表示从下标i到下标j中能取得的最大硬币数量

### 315. Count of Smaller Numbers After Self
	凡是这种限定数据范围、下标范围，统计个数的问题，都可以用树状数组来做。必要时可以用哈希表对离散化的数据进行连续化映射。

### 316. Remove Duplicate Letters
	使用两个字典数组分别记录源字符串和结果字符串的字符数量。

### 318. Maximum Product of Word Lengths
	转化为位运算

### 319. Bulb Switcher
	数学题

### 321. Create Maximum Number
	这题包含的内容有点多，首先需要写一个从数组中取k个最大序列的算法，其次
	需要实现归并排序，排序的大小判断依据是两个数组之间的比较

### 322. Coin Change
	背包问题

### 324. Wiggle Sort II
	先排序

### 326. Power of Three
	循环呗

### 327. Count of Range Sum
	这题用树状数组，但是看不懂网上的题解!!!

### 328. Odd Even Linked List
	easy

### 329. Longest Increasing Path in a Matrix
	记忆化搜索+dp，和拼多多那一题一样

### 330. Patching Array
	太巧秒了

### 331. Verify Preorder Serialization of a Binary Tree
	记录每一个结点的出度和入度的差值diff = outdegree - indegree，每遍历下一个结点，diff减1，因为它提供了一个入度，当当前结点非空时,diff加2，因为它提供两个出度。遍历结果，如果diff==0，则为二叉树。

### 332. Reconstruct Itinerary
	拓扑排序，以字符串代替下标的结点邻接矩阵使用map<string,map<string,int>>来构造，当所有边全部遍历完的时候，排序结束

### 334. Increasing Triplet Subsequence
	使用a1,a2分别记录三元组的前两个递增元素