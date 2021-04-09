class Solution(object):
    def countVowelStrings(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n==0: return 0
        dp = [1]*5
        for i in range(2, n+1):
            for j in range(5):
                for k in range(j+1, 5):
                    dp[j] += dp[k]

        return sum(dp)