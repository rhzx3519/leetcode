class Solution(object):
    def dieSimulator(self, n, rollMax):
        """
        :type n: int
        :type rollMax: List[int]
        :rtype: int
        """
        MOD = 1e9 + 7
        maxK = max(rollMax)
        dp = [[[0]*(maxK + 1) for _ in range(6)] for _ in range(n)]
        for j in range(6):
            dp[0][j][1] = 1
        
        for i in range(1, n):
            for j in range(6):
                for k in range(1, rollMax[j]+1):
                    if k==1:
                        dp[i][j][k] = sum([sum(dp[i-1][l])%MOD for l in range(6) if l!=j])
                    else:
                        dp[i][j][k] = dp[i-1][j][k-1]
        
        # print dp
        return int(sum(sum(col) for col in dp[-1])%MOD)

# 作者：niu-meng
# 链接：https://leetcode-cn.com/problems/dice-roll-simulation/solution/dong-tai-gui-hua-si-lu-qing-xi-by-niu-me-m27g/
# 来源：力扣（LeetCode）
# 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。