class Solution(object):
    def visiblePoints(self, points, angle, location):
        """
        :type points: List[List[int]]
        :type angle: int
        :type location: List[int]
        :rtype: int
        """
        same = 0
        angles = []
        for x, y in points:
            if x==location[0] and y==location[1]:
                same += 1
            else:
                angles.append(math.atan2(y-location[1], x-location[0]) * 180.0/math.pi)
        angles.sort()
        angles.extend([x + 360 for x in angles])

        # print angles, same
        max_val = l = 0
        for r in range(len(angles)):
            while (angles[r] - angles[l]) > angle:
                l += 1
            max_val = max(max_val, r-l+1)
        return max_val + same
        
# 思路：通过atan求每个点到location的极角，排序之后，利用双指针求取区间最大值