class Solution(object):
    def numOfSubarrays(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        Mod = 10**9+7
        n = len(arr)

        ans = 0
        pre = odd = 0
        even = 1
        for num in arr:
            pre += num
            if pre&1: # 当前前缀和为奇数，新增的和为奇数的子数组个数=之前前缀和等于偶数的子数组数目, sum[j, i] = sum[i] - sum[j]
                ans += even
                odd += 1
            else:   # 当前前缀和为偶数，新增的和为奇数的子数组个数=之前前缀和等于奇数的子数组数目
                ans += odd
                even += 1
        return ans%Mod
# 思路：dp[i]表示i和为 奇数 的子数组数目
