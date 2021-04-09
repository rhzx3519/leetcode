class Solution(object):
    def carFleet(self, target, position, speed):
        """
        :type target: int
        :type position: List[int]
        :type speed: List[int]
        :rtype: int
        """
        cars = [ (p, float(target - p)/s) for p, s in zip(position, speed)]
        n = len(cars)
        cars.sort(key=lambda x: x[0], reverse=True)
        ans = 0
        t = float('-inf')
        for i in range(n):
            if cars[i][1] > t:
                ans += 1
                t = cars[i][1]
        return ans
# 思路，按到终点的距离降序排序，距离近且耗时长的车会组成一个车队
# time: O(nlogn), space: O(1)