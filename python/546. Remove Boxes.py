class Solution(object):
    def removeBoxes(self, boxes):
        """
        :type boxes: List[int]
        :rtype: int
        """
        n = len(boxes)
        mem = {}
        def dp(l, r, k):
            if l > r:
                return 0
            if (l, r, k) in mem:
                return mem[(l, r, k)]
            while r > l and boxes[r]==boxes[r-1]:
                r -= 1
                k += 1
            ans = dp(l, r-1, 0) + (k+1)*(k+1)
            for i in range(l, r):
                if boxes[i]==boxes[r]:
                    ans = max(ans, dp(l, i, k+1) + dp(i+1, r-1, 0))
            mem[(l, r, k)] = ans
            return ans
        return dp(0, n-1, 0)
            