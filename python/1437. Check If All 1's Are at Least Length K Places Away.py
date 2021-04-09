class Solution(object):
    def kLengthApart(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        start = -1
        for i, t in enumerate(nums):
            if t==1:
                # print start, i
                if start != -1 and i-start-1 < k:
                    return False
                start = i
        return True
            