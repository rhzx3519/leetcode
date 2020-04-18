class Solution(object):
    def numDistinct(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        ls = len(s)
        lt = len(t)
        dp = [[0]*(lt+1) for _ in range(ls+1)]
        dp[0][0] = 1
        for i in range(1, ls+1):
            dp[i][0] = 1
        for i in range(1, ls+1):
            for j in range(1, lt+1):
                if i < j: continue
                if s[i-1]==t[j-1]:
                    dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
                else:
                    dp[i][j] = dp[i-1][j]
        
        return dp[ls][lt]

# 我来解释下2个问题，1： 为啥状态方程这样对？ 2：怎么想到这样的状态方程？
# 我个人习惯dp[i][j] 表示为s[0-i] 和t[0-j]均闭区间的子序列个数，但这样不能表示s和t空串的情况
# 所以声明 int[][] dp = new int[m + 1][n + 1]; 这样dp[0][x]可以表示s为空串，dp[x][0]同理。
# 先不扣初始化的细节，假设dp[i][j] 就是s[i] 和t[j] 索引的元素子序列数量
# 1：为啥状态方程是： s[i] == t[j] 时 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
# s[i] != t[j] 时 dp[i][j] = dp[i-1][j]
# 先看s[i] == t[j] 时，以s = "rara" t = "ra" 为例，当i = 3, j = 1时，s[i] == t[j]。
# 此时分为2中情况，用最后一位的a + 不用最后一位的a。
# 如果用s串最后一位的a,那么t串最后一位的a也被消耗掉，此时的子序列其实=dp[i-1][j-1]
# 如果不用s串最后一位的a，那就得看"rar"里面是否有"ra"子序列的了，就是dp[i-1][j]
# 所以 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
# 再看s[i] != t[j] 比如 s = "rarb" t = "ra" 还是当i = 3, j = 1时，s[i] != t[j]
# 此时显然最后的b想用也用不上啊。所以只能指望前面的"rar"里面是否有能匹配"ra"的
# 所以此时dp[i][j] = dp[i-1][j]
# 2: 怎么想到这样状态方程的？
# 一点个人经验，见过的很多2个串的题，大部分都是dp[i][j] 分别表示s串[0...i] 和t串[0...j]怎么怎么样 然后都是观察s[i]和t[j]分等或者不等的情况 而且方程通常就是 dp[i-1][j-1] 要么+ 要么 || dp[i-1][j]类似的
# 类似的题比如有 10：正则表达式匹配 44：通配符匹配 编辑距离 1143：最长公共子序列等等的 还有几道想不起来了