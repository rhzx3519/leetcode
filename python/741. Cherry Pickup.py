# -*- coding: utf-8 -*-

class Solution(object):
    def cherryPickup(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        n = len(grid)
        dp = [[[-1]*(n+1) for _ in range(n+1)] for _ in range(n+1)]
        dp[1][1][1] = grid[0][0]
        for i in range(1, n+1):
        	for j in range(1, n+1):
        		for k in range(1, min(i+j, n)+1):
        			l = i + j - k
        			if l>n or l<1:
        				continue
        			if grid[i-1][j-1]==-1 or grid[k-1][l-1]==-1:
        				continue

        			ans = max(dp[i-1][j][k-1], dp[i-1][j][k], dp[i][j-1][k], dp[i][j-1][k-1])
        			if ans < 0:
        				continue
        			ans += grid[i-1][j-1]
        			if i!=k or j!=l:
        				ans += grid[k-1][l-1]
        			dp[i][j][k] = ans
        return dp[n][n][n] if dp[n][n][n]>0 else 0

# 可以这样想：从（1，1）到（n，n）一定会经过一个点（i，j） 从（n，n）回到（1,1）一定会经过一个点（k，l）
# 而进入（i，j）必然是从他上面或者左边进入的 从（k，l）回去一定是从他右边或者下边进去这个点的
# 所以定义dp(i,j,k,l)表示从（1,1）进入（i,j）再从（i，j）进入（n，n），再从（n，n）回到（k，l）再回到（1,1）能摘到的最大的樱桃数
# 可以这样转移：
# dp(i,j,k,l)=max{dp(i-1,j,k-1,l),dp(i-1,j,k,l-1),dp(i,j-1,k-1,l),dp(i,j-1,k,l-1)}+grid(i,j)+grid(k,l)
# 但是这样做是O（n^4）超时
# 可以注意点从（1,1）到（i，j）和从（k，l）到（1,1）的步数是相同的 所以i+j=k+l 这样就不用枚举l了 复杂度降到O（n^3）

# 作者：JiaQiaoY
# 链接：https://leetcode-cn.com/problems/cherry-pickup/solution/dong-tai-gui-hua-by-jiaqiaoy/
# 来源：力扣（LeetCode）
# 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。        


if __name__ == '__main__':
	# grid = [[1,1,-1],[1,-1,1],[-1,1,1]]
	grid = [[1,1,1,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,0,0,1],[1,0,0,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,1,1,1]]
	su = Solution()
	print su.cherryPickup(grid)
