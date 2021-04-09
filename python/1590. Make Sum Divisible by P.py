class Solution(object):
    def minSubarray(self, nums, p):
        """
        :type nums: List[int]
        :type p: int
        :rtype: int
        """
        n = len(nums)
        s = sum(nums)
        if s%p==0:
            return 0
        k = s%p
        mem = {0:0}
        s = 0
        ans = n
        for i in range(1, n+1):
            s += nums[i-1]
            x = (s - k + p)%p
            if x in mem:
                ans = min(ans, i - mem[x])
            mem[s%p] = i
        return ans if ans != n else -1

# 思路：寻找最短的子数组，满足子数组的和s%p=k, 其中k=sum(nums)%p
# t = (a - k)%p, b%p == t. --> (a-b)%p = k