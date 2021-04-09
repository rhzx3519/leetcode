class Solution(object):
    def findMaxAverage(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: float
        """
        s = sum(nums[:k])
        max_val = s/float(k)
        for i in range(k, len(nums)):
            s -= nums[i-k]
            s += nums[i]
            max_val = max(s/float(k), max_val)
        return max_val