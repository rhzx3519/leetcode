class Solution(object):
    def maxAbsoluteSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        p_val = n_val = 0
        s = 0
        n = len(nums)
        for r in range(n):
            s += nums[r]
            if s < 0:
                s = 0
            p_val = max(p_val, s)

        s = 0
        for num in nums:
            s += num
            if s > 0:
                s = 0
            n_val = min(n_val, s)   

        # print n_val
        return max(p_val, -n_val)        


if __name__ == '__main__':
    nums = [-3,-5,-3,-2,-6,3,10,-10,-8,-3,0,10,3,-5,8,7,-9,-9,5,-8]
    su = Solution()
    su.maxAbsoluteSum(nums)