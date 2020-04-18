class Solution(object):
    def maxEnvelopes(self, envelopes):
        """
        :type envelopes: List[List[int]]
        :rtype: int
        """
        envelopes.sort(key=lambda x: (x[0], -x[1]))
        nums = [x[1] for x in envelopes]

        def LIS(nums):
            n = len(nums)
            dp = [1]*(n+1)
            dp[0] = 0
            for i in range(n+1):
                for j in range(1, i):
                    if nums[i-1] > nums[j-1]:
                        dp[i] = max(dp[i], dp[j]+1)
            return max(dp)
        return LIS(nums)

# 思路： 
# 先对宽度 w 进行升序排序，如果遇到 w 相同的情况，则按照高度 h 降序排序。之后把所有的 h 作为一个数组，在这个数组上计算 LIS 的长度就是答案。
# 这个解法的关键在于，对于宽度 w 相同的数对，要对其高度 h 进行降序排序。因为两个宽度相同的信封不能相互包含的，逆序排序保证在 w 相同的数对中最多只选取一个。