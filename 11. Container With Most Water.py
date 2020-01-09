class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        res = 0
        l, h = 0, len(height) - 1
        while l < h:
            area = 0
            if height[l] < height[h]:
                area = (h-l)*height[l]
                l += 1
            else:
                area = (h-l)*height[h]
                h -= 1
            res = max(area, res)
        
        return res
        