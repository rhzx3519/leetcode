class Solution(object):
    def findMinFibonacciNumbers(self, k):
        """
        :type k: int
        :rtype: int
        """
        dp = [1, 1]
        i = 2
        while i <= k:
            dp.append(dp[i-1] + dp[i-2])
            if dp[i] > k: break        
            i += 1

        ans = 0
        for i in range(len(dp)-1, -1, -1):
            if k >= dp[i]:
                k -= dp[i]
                ans += 1
        return ans

if __name__ == '__main__':
    k = 645157245
    su = Solution()
    print su.findMinFibonacciNumbers(k)