class Solution(object):
    def numFactoredBinaryTrees(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        mod = 10**9+7
        hash = collections.defaultdict(int)
        dp = collections.defaultdict(int)
        for i in A:
            hash[i] += 1
            dp[i] = 1
        n = len(A)
        res = 0
        A.sort()
        for k in range(n):
            cur = A[k]
            for j in range(k):
                if cur%A[j]==0 and cur/A[j] in hash:
                    dp[cur] = (dp[cur] + dp[A[j]] * dp[cur/A[j]]) % mod
            res = (res+dp[cur])%mod
        return res
