class Solution(object):
    def smallerNumbersThanCurrent(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        bucket = [0]*101
        for num in nums:
            bucket[num] += 1
        
        cnt = 0
        for i in range(101):
            t = bucket[i]
            bucket[i] = cnt
            cnt += t

        ans = []
        for i in nums:
            ans.append(bucket[i])
        return ans