class Solution(object):
    def summaryRanges(self, nums):
        """
        :type nums: List[int]
        :rtype: List[str]
        """
        if not nums: return []
        n = len(nums)
        last = 0
        ans = []
        for i in range(1, n):
            if nums[i] != nums[i-1] + 1:
                if nums[last] == nums[i-1]:
                    ans.append(str(nums[last]))
                else:
                    ans.append(str(nums[last]) + "->" + str(nums[i-1]))
                last = i

        if nums[last] == nums[-1]:
            ans.append(str(nums[last]))
        else:
            ans.append(str(nums[last]) + "->" + str(nums[-1]))
        return ans
        