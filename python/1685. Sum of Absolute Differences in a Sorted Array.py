class Solution(object):
    def getSumAbsoluteDifferences(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        ans = []
        tmp = 0
        for i in range(1, n):
            tmp += abs(nums[0] - nums[i])
        ans.append(tmp)

        for i in range(1, n):
            x = nums[i] - nums[i-1]
            t = ans[-1] + x*i - x*(n-i)
            ans.append(t)
        return ans