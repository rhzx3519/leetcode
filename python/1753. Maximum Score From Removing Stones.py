class Solution(object):
    def maximumScore(self, a, b, c):
        """
        :type a: int
        :type b: int
        :type c: int
        :rtype: int
        """
        stones = [a,b,c]
        stones.sort()
        if stones[0] + stones[1] <= stones[2]:
            return stones[0] + stones[1]
        return self.maximumScore(stones[0]-1, stones[1]-1, stones[2]) + 1