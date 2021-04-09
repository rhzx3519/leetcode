class Solution(object):
    def waysToMakeFair(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums)&1: nums.append(0)
        n = len(nums)
        ans = 0
        s = t = 0
        for i in range(0, n-1, 2):
            s += nums[i+1] - nums[i]
        for i in range(0, n-1, 2):
            s -= nums[i+1] - nums[i]
            if t == s + nums[i+1]:
                ans += 1
            if i < n-2 and t == s + nums[i]:
                ans += 1
            t += nums[i+1] - nums[i]
        return ans


if __name__ == '__main__':
    nums = [2,1,6,4]
    su = Solution()
    su.waysToMakeFair(nums)