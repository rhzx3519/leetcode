class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        res = 0
        l, r = 0, len(height) - 1
        while l < r:
            minh = min(height[l], height[r])
            if height[l] == minh:
                l += 1
                while l < r and height[l] < minh:
                    res += minh - height[l]
                    l += 1
            else:
                r -= 1
                while l < r and height[r] < minh:
                    res += minh - height[r]
                    r -= 1
        return res