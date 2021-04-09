class Solution(object):
    def checkPossibility(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        n = len(nums)
        if n <= 1: return True
        cnt = 0
        for i in range(n-1):
            x, y = nums[i], nums[i+1]
            if x <= y: continue
            cnt += 1
            if cnt > 1: return False
            if i>=1 and y < nums[i-1]:
                nums[i+1] = x # nums[i]==x >= nums[i-1]

        return True