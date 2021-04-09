class Solution(object):
    def tallestBillboard(self, rods):
        """
        :type rods: List[int]
        :rtype: int
        """
        dp = {0 : 0}
        for x in rods:
            for k, v in list(dp.items()):
                dp[k + x] = max(dp.get(k + x, 0), v + x)
                dp[k - x] = max(dp.get(k - x, 0), v)
        return dp[0]

# 思路：问题可以转化为，rods数组中的元素分别乘以-1, 0, 1时，总和等于0时，正数和的最大值。 dp[i]表示总和等于i时的最大正数和.