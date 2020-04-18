class Solution(object):
    def subarraySum(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        d = collections.defaultdict(int)
        d[0] = 1
        s = 0
        res = 0
        for i in range(len(nums)):
            s += nums[i]
            if d.has_key(s-k):
                res += d[s-k]
            d[s] += 1
        return res