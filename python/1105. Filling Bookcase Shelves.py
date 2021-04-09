class Solution(object):
    def minHeightShelves(self, books, shelf_width):
        """
        :type books: List[List[int]]
        :type shelf_width: int
        :rtype: int
        """
        n = len(books)
        dp = [float('inf')] * (n + 1)
        dp[0] = 0
        for i in range(1, n+1):
            w = h = 0
            for j in range(i, 0, -1):
                w += books[j-1][0]
                if w > shelf_width:
                    break
                h = max(h, books[j-1][1])
                dp[i] = min(dp[i], dp[j-1] + h)
        return dp[n]
