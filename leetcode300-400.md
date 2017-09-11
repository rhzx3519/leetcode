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
	1. 凡是这种限定数据范围、下标范围，统计个数的问题，都可以用树状数组来做。必要时可以用哈希表对离散化的数据进行连续化映射。
	2. 归并排序求逆序数

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

### 337. House Robber III
	dfs，每个结点都有两种状态，一种是rob当前结点，另一种是不rob当前结点。

### 338. Counting Bits
	没想到这题也是个dp题，可以推出以下递归公式
	dp[i] = dp[i>>1] + (i&1)

### 341. Flatten Nested List Iterator
	使用栈，从后向前遍历将所有元素入栈

### 342. Power of Four
	...

### 343. Integer Break
	dp, dp[n]表示分成至少两个整数的乘积的最大值。有递推公式dp[n] = i*max(n-i, dp[n-i])  2<=i <=n;

### 344. Reverse String
	...

### 345. Reverse Vowels of a String
	...

### 347. Top K Frequent Elements
	priority_queue<pair<int, int>> pq 默认按照pair的first排列的大顶堆

### 349. Intersection of Two Arrays
	这题有点奇怪

### 350. Intersection of Two Arrays II
	...

### 355. Design Twitter
	主要是数据结构的设计

### 357. Count Numbers with Unique Digits
	数学题

### 365. Water and Jug Problem
	看看x、y的最大公约数是否能被z整除

### 367. Valid Perfect Square
	r = (r + x/r)/2, 可以通过其他方法求开方

### 368. Largest Divisible Subset
	这道题给了我们一个数组，让我们求这样一个子集合，集合中的任意两个数相互取余均为0，而且提示中说明了要使用DP来解。那么我们考虑，较小数对较大数取余一定为0，那么问题就变成了看较大数能不能整除这个较小数。那么如果数组是无序的，处理起来就比较麻烦，所以我们首先可以先给数组排序，这样我们每次就只要看后面的数字能否整除前面的数字。定义一个动态数组dp，其中dp[i]表示到数字nums[i]位置最大可整除的子集合的长度，还需要一个一维数组parent，来保存上一个能整除的数字的位置，两个整型变量mx和mx_idx分别表示最大子集合的长度和起始数字的位置，我们可以从后往前遍历数组，对于某个数字再遍历到末尾，在这个过程中，如果nums[j]能整除nums[i], 且dp[i] < dp[j] + 1的话，更新dp[i]和parent[i]，如果dp[i]大于mx了，还要更新mx和mx_idx，最后循环结束后，我们来填res数字，根据parent数组来找到每一个数字

### 371. Sum of Two Integers
	a^b求和，(a&b)>>1求进位

### 372. Super Pow
	数学题
	a^1234567 % k = (a^1234560 % k) * (a^7 % k) % k = (a^123456 % k)^10 % k * (a^7 % k) % k
	f(a,1234567) = f(a, 1234560) * f(a, 7) % k = f(f(a, 123456),10) * f(a,7)%k;

### 373. Find K Pairs with Smallest Sums
	无脑优先队列，注意构造的时候需要传入一个重载了小括号运算符的结构体作为参数。

### 374. Guess Number Higher or Lower
	二分

### 375. Guess Number Higher or Lower II
	记忆化搜索+dp
	For each number x in range[i~j]
	we do: result_when_pick_x = x + max{DP([i~x-1]), DP([x+1, j])}
	--> // the max means whenever you choose a number, the feedback is always bad and therefore leads you to a worse branch.
	then we get DP([i~j]) = min{xi, ... ,xj}
	--> // this min makes sure that you are minimizing your cost.

### 376. Wiggle Subsequence
	贪心

### 377. Combination Sum IV
	dp+记忆化搜索 dp[i]表示和为i的组合数目

### 378. Kth Smallest Element in a Sorted Matrix
	二分

### 380. Insert Delete GetRandom O(1)
	vector + unordered_map

### 381. Insert Delete GetRandom O(1) - Duplicates allowed
	同样用vector<pair<int, int>> + unordered_map<int, vector<int>>，此题中vector存储val和val在map中的位置, map存储val和val在vector中的位置

### 382. Linked List Random Node
	这里用到了著名了水塘抽样的思路，由于限定了head一定存在，所以我们先让返回值res等于head的节点值，然后让cur指向head的下一个节点，定义一个变量i，初始化为2，若cur不为空我们开始循环，我们在[0, i - 1]中取一个随机数，如果取出来0，那么我们更新res为当前的cur的节点值，然后此时i自增一，cur指向其下一个位置，这里其实相当于我们维护了一个大小为1的水塘，然后我们随机数生成为0的话，我们交换水塘中的值和当前遍历到底值，这样可以保证每个数字的概率相等
	水塘算法伪代码如下：
	//stream代表数据流  
	//reservoir代表返回长度为k的池塘  
	  
	//从stream中取前k个放入reservoir；  
	for ( int i = 1; i < k; i++)  
	    reservoir[i] = stream[i];  
	for (i = k; stream != null; i++) {  
	    p = random(0, i);  
	    if (p < k) reservoir[p] = stream[i];  
	return reservoir;  

### 383. Ransom Note
	...

### 384. Shuffle an Array
	like insert sort

### 385. Mini Parser
	top-down parser

### 386. Lexicographical Numbers
	有点怪

### 387. First Unique Character in a String
	...

### 388. Longest Absolute File Path
	...

### 389. Find the Difference
	...

### 390. Elimination Game
	数学题