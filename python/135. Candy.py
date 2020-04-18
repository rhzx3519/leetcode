class Solution(object):
    def candy(self, ratings):
        """
        :type ratings: List[int]
        :rtype: int
        """
        n = len(ratings)
        candies = [1]*n
        for i in range(1, n):
            if ratings[i] > ratings[i-1]:
                candies[i] = candies[i-1]+1
        
        for i in range(n-1, 0, -1):
            if ratings[i-1] > ratings[i]:
                candies[i-1] = max(candies[i-1], candies[i]+1)
        return sum(candies)

# 思路： 先从左向右遍历，确定右边元素的candy， 再从右向左遍历，确定左边元素的candy        