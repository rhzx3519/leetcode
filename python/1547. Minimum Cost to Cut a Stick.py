class Solution(object):
    def minCost(self, n, cuts):
        """
        :type n: int
        :type cuts: List[int]
        :rtype: int
        """
        cuts.extend([0, n])
        cuts.sort()
        mem = {}
        def dp(i, j):
            if i+1 == j: return 0
            if (i, j) in mem:
                return mem[(i, j)]
            
            ans = float('inf')
            for k in range(i+1, j):
                ans = min(ans, dp(i, k) + dp(k, j) + cuts[j] - cuts[i])
            mem[(i, j)] = ans
            return ans
        
        return dp(0, len(cuts)-1)