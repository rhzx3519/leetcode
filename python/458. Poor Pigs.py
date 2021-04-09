class Solution(object):
    def poorPigs(self, buckets, minutesToDie, minutesToTest):
        """
        :type buckets: int
        :type minutesToDie: int
        :type minutesToTest: int
        :rtype: int
        """
        base = minutesToTest//minutesToDie + 1
        pigs = 0
        while base**pigs < buckets:
            pigs += 1
        return pigs

# 对于minutesToTest=60, minutesToDie=15
# 一只猪一个小时内可以喝4桶水，可以检测5个水桶