### 101. Symmetric Tree
	递归
### 102. Binary Tree Level Order Traversal
	首先在根节点之后插入一个空指针当做标记，每次从队列中取到空节点，就知道这一层已经遍历完了，将结果保存到返回集合汇总，重新插入一个空指针当做标记。

### 103. Binary Tree Zigzag Level Order Traversal
	跟上一题一样，加入一个bool值判断是否需要反转

### 104. Maximum Depth of Binary Tree
	递归

### 105. Construct Binary Tree from Preorder and Inorder Traversal
	递归

### 106. Construct Binary Tree from Inorder and Postorder Traversal
	递归

### 107. Binary Tree Level Order Traversal II
	层序遍历

### 108. Convert Sorted Array to Binary Search Tree
	每次选中间数当根结点

### 109. Convert Sorted List to Binary Search Tree
	递归

### 110. Balanced Binary Tree
	递归

### 111. Minimum Depth of Binary Tree
	递归

### 112. Path Sum
	使用逻辑或判断返回结果

### 113. Path Sum II
	递归

### 114. Flatten Binary Tree to Linked List
	分析：本题采用递归的方法解决，关键是要知道由左子树转化的链表的头和尾，以及由右子树转化的链表的头和尾。头一定是二叉树的根节点，尾是右子树的尾（如果右子树不空）或者左子树的尾（如果右子树空，左子树不空）或者根（如果左右子树都是空）。

### 115. Distinct Subsequences
	分析：一般来说，如果题目里面给出两个字符串，基本是两种思路，一种就是递归判断，一种就是动态规划，这里我们可以用f(i,j)表示S中前i个字符串中，T的前j个字符出现的次数，不管S[i]和T[j]相不相等，首先f(i,j)=f(i-1,j)，其次要是S[i]==T[j]的话，f(i,j) = f(i-1,j)+f(i-1,j-1),可以看到，i的状态只与i-1有关，于是可以用滚动数组来进行优化。代码类似01背包

### 116. Populating Next Right Pointers in Each Node
	和下一题一样

### 117. Populating Next Right Pointers in Each Node II
	十分巧妙

### 118. Pascal's Triangle
	dp

### 119. Pascal's Triangle II
	跟上题一样，不同的是，只用两个一维数组保存最近的结果

### 120. Triangle
	从底向顶更新路径和保存在三角形中。

### 121. Best Time to Buy and Sell Stock
	一遍遍历

### 122. Best Time to Buy and Sell Stock II
	一遍遍历

### 123. Best Time to Buy and Sell Stock III
	动态规划法。以第i天为分界线，计算第i天之前进行一次交易的最大收益preProfit[i]，和第i天之后进行一次交易的最大收益postProfit[i]，最后遍历一遍，max{preProfit[i] + postProfit[i]} (0≤i≤n-1)就是最大收益。第i天之前和第i天之后进行一次的最大收益求法同Best Time to Buy and Sell Stock I。

### 124. Binary Tree Maximum Path Sum
	递归

### 125. Valid Palindrome
	... 

### 127. Word Ladder
	bfs，使用26个字母依次替换队列头的每一个字母，看看list是否存在，如果存在则压入队列，从list删除该字母

### 128. Longest Consecutive Sequence
	我们也可以采用哈希表来做，刚开始哈希表为空，然后遍历所有数字，如果该数字不在哈希表中，那么我们分别看其左右两个数字是否在哈希表中，如果在，则返回其哈希表中映射值，若不在，则返回0，然后我们将left+right+1作为当前数字的映射，并更新res结果，然后更新d-left和d-right的映射值.

### 129. Sum Root to Leaf Numbers
	dfs

### 131. Palindrome Partitioning
	dfs

### 133. Clone Graph
	bfs

### 134. Gas Station
	类似求数组最大子串和

### 135. Candy
	贪心两遍遍历

### 136. Single Number
	异或

### 137. Single Number II
	位操作

### 138. Copy List with Random Pointer
	hash

### 139. Word Break
	解题思路：
	s = lintcode 可以分为 l和intcode，若是l是可分割单词，并且intcode单词在dict中，则表示s也是可分割单词。
	若不符合，s = lintcode 可以分为 li和ntcode，若是li是可分割单词，并且ntcode单词在dict中，则表示s也是可分割单词。
	...
	同理得出BWord[ n ]表示字符串S[0,n]是否是可分割单词,是则为true，否为false。
	BWord[ n ] = BWord[ i ] &&( S[ i+1 ,n ]在dict中 )

### 141. Linked List Cycle
	快慢指针

### 142. Linked List Cycle II
	hash

### 143. Reorder List
	先拆后合并

### 144. Binary Tree Preorder Traversal
	非递归遍历很慢

### 145. Binary Tree Postorder Traversal
	非递归

### 146. LRU Cache
	思路: 这题的大致思路就是用一个hash表来保存已经存在的key, 然后用另外一个线性容器来存储其key-value值, 我们可以选择链表list, 因为需要调整结点的位置, 而链表可以在O(1)时间移动结点的位置, 数组则需要O(n). 

	如果新来一个set请求, 我们先去查hash表

	1. 如果已经存在了这个key, 那么我们需要更新其value, 然后将其在list的结点位置移动到链表首部.

	2. 如果不存在这个key, 那么我们需要在hash表和链表中都添加这个值, 并且如果添加之后链表长度超过最大长度, 我们需要将链表尾部的节点删除, 并且删除其在hash表中的记录


	如果来了一个get请求, 我们仍然先去查hash表, 如果key存在hash表中, 那么需要将这个结点在链表的中的位置移动到链表首部.否则返回-1.

	另外一个非常关键的降低时间复杂度的方法是在hash中保存那个key在链表中对应的指针, 我们知道链表要查找一个结点的时间复杂度是O(n), 所以当我们需要移动一个结点到链表首部的时候, 如果直接在链表中查询那个key所对于的结点, 然后再移动, 这样时间复杂度将会是O(n), 而一个很好的改进方法是在hash表中存储那个key在链表中结点的指针, 这样就可以在O(1)的时间内移动结点到链表首部.

	c++的stl库提供非常丰富的函数, 掌握这些东西将会让代码长度大大减小.

### 147. Insertion Sort List
	插入

### 148. Sort List
	快速排序的单边扫描算法：
	通过partition找到切分位置，p指针从begin位置开始，一直向后遍历到end，找到第一个不连续的小于key的位置q，交换p和q的值，遍历完成之后，交换初始位置和和q的值

### 149. Max Points on a Line
	二层循环遍历，用hash统计斜率相同的点，注意hash的key是pair，注意key的归一化处理

### 150. Evaluate Reverse Polish Notation
	栈

### 151. Reverse Words in a String
	...

### 152. Maximum Product Subarray
	分析：访问到每个点的时候，以该点为子序列的末尾的乘积，要么是该点本身，要么是该点乘以以前一点为末尾的序列，注意乘积负负得正，故需要记录前面的最大最小值。

### 153. Find Minimum in Rotated Sorted Array
	比较巧妙的二分，

### 154. Find Minimum in Rotated Sorted Array II
	还是二分

### 155. Min Stack
	单独用一个栈记录最小值

### 162. Find Peak Element
	二分

### 164. Maximum Gap
	简单来

### 165. Compare Version Numbers
	坑爹

### 166. Fraction to Recurring Decimal
	对于可以整除的，和对于不可以整除但是算出来的小数不是循环小数的，都可以直接计算结果转化为字符串。
	比较麻烦的是循环节的小数。为了找到循环节，可以用hashtable保存每一步余数，如果当前余数和之前某个余数相同，则重复之间的这段发生循环。

### 167. Two Sum II - Input array is sorted
	二分

### 168. Excel Sheet Column Title
	十进制转26进制

### 169. Majority Element
	一遍遍历

### 171. Excel Sheet Column Number
	26进制转10进制

### 172. Factorial Trailing Zeroes
	编程之美

### 173. Binary Search Tree Iterator
	二叉树非递归遍历

### 174. Dungeon Game
	逆向

### 179. Largest Number
	排序

### 187. Repeated DNA Sequences
	字符串中只包含4个字符：A, C, G, T。它们的ASCII码的二进制形式是：
	A : 0100 0001
	C : 0100 0011
	G : 0100 0111
	T : 0101 0100
	这4个字符的末3位是不同的，因此，我们可以 3个bits 来表示其中的一个字符。
	又因为每个子串的长度为10，因此总的位数是：10 x 3 = 30，一个int就足够存放它。

### 189. Rotate Array
	和反转链表一样

### 190. Reverse Bits
	位运算

### 191. Number of 1 Bits
	n&=(n-1), 判断一个数是否是2的方幂n>0&&n&(n-1)

### 198. House Robber
	分析：又是一个动态规划问题，我们设dp[i]表示到i为止且包括i为最后一个盗窃的房子得到的最大值，很明显取决于dp[i-2]和dp[i-3],和dp[i-4]无关了，因为取dp[i-4]的情况可以包含在dp[i-2]的情况下，注意使用滚动数组

### 199. Binary Tree Right Side View
	dfs, 巧妙

### 200. Number of Islands
	dfs