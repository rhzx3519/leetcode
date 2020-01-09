class Solution(object):
    def optimalDivision(self, nums):
        """
        :type nums: List[int]
        :rtype: str
        """
        n = len(nums)
        if n==1:
            return str(nums[0])
        if n==2:
            return '/'.join([str(i) for i in nums])
        return '{}/({})'.format(nums[0], '/'.join([str(i) for i in nums[1:]]))
