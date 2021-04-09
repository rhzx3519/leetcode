


# 为了理解本文的算法，我们需要熟悉 Kadane 算法。在这个章节，我们解释算法背后的核心逻辑。

# 对于一个给定数组 A，Kadane 算法可以用来找到 A 的最大子段和。这里，我们只考虑非空子段。

# Kadane 算法基于动态规划。令 dp[j] 为以 A[j] 结尾的最大子段和。也就是，

# \text{dp}[j] = \max\limits_i (A[i] + A[i+1] + \cdots + A[j])
# dp[j]= 
# i
# max
# ​   
#  (A[i]+A[i+1]+⋯+A[j])

# 那么，以 j+1 j结尾的子段（例如 A[i], A[i+1] + ... + A[j+1]）最大化了 A[i] + ... + A[j] 的和，当这个子段非空那么就等于 dp[j] 否则就等于 0。所以，有以下递推式：

# \text{dp}[j+1] = A[j+1] + \max(\text{dp}[j], 0)
# dp[j+1]=A[j+1]+max(dp[j],0)

# 由于一个子段一定从某个位置截止，所以 \max\limits_j dp[j] 
# j
# max
# ​   
#  dp[j] 就是需要的答案。

# 为了计算 dp 数组更快，Kadane 算法通常节约空间复杂度的形式表示。我们只维护两个变量 ans 等于 \max\limits_j dp[j] 
# j
# max
# ​   
#  dp[j] 和 cur 等于 dp[j]dp[j]。随着 j 从 00 到 A.\text{length}-1A.length−1 遍历。

# 作者：LeetCode
# 链接：https://leetcode-cn.com/problems/maximum-sum-circular-subarray/solution/huan-xing-zi-shu-zu-de-zui-da-he-by-leetcode/
# 来源：力扣（LeetCode）
# 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。