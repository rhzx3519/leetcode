class Solution(object):
    def maxJumps(self, arr, d):
        """
        :type arr: List[int]
        :type d: int
        :rtype: int
        """
        n = len(arr)
        a = [i for i in range(n)]
        a.sort(key=lambda x: arr[x])
        dp = [1] * n
        for i in range(n):
            idx = a[i]
            dp[idx] = 1
            # 向左找
            for j in range(idx-1, idx-d-1, -1):
                if j < 0 or arr[j] >= arr[idx]:
                    break
                dp[idx] = max(dp[idx], dp[j] + 1)
            # 向右找
            for j in range(idx+1, idx+d+1, 1):
                if j >= n or arr[j] >= arr[idx]:
                    break
                dp[idx] = max(dp[idx], dp[j] + 1)
        return max(dp) 
                

# dp，从高度最小的下标开始迭代, 向左向右找比i高度小的下标
# dp[i] = max(dp[i], dp[j]), dp[j] = dp[i-d] -> dp[i+d]
# time: O(N^2), space: O(N)