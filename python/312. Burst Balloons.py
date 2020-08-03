class Solution(object):
    def maxCoins(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        mem = {}

        def dp(i, j, ll, rr):
            if (i, j) in mem:
                return mem[(i, j)]
            if i==j:
                mem[(i, j)] = ll * nums[i] * rr
                return mem[(i, j)]
            ans = 0
            for k in range(i, j+1):
                left = dp(i, k-1, ll, nums[k]) if i<k else 0
                right = dp(k+1, j, nums[k], rr) if j>k else 0
                ans = max(ans, left + right + ll*nums[k]*rr)
            mem[(i, j)] = ans
            return ans

        return dp(0, n-1, 1, 1)

# 思路：dp(i, j, ll, rr)表示 [i, j]下标i->j可以获得的最大硬币数量, ll/rr表示左右两边的最大硬币数量
# time: O(N^3), space: O(N)

