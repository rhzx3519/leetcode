class Solution(object):
    def findShortestSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        count = collections.Counter(nums)
        degree = max(count.values())
        n = len(nums)
        l = 0
        ans = float('inf')
        mem = collections.defaultdict(int)
        max_val = 0
        # print degree
        for r in range(n):
            mem[nums[r]] += 1
            max_val = max(mem.values())
            # print max_val, mem
            
            while l <= r and max_val == degree:
                mem[nums[l]] -= 1
                if mem[nums[l]]==0: del mem[nums[l]]
                max_val = max(mem.values()) if mem else 0
                ans = min(ans, r-l+1)
                l += 1
        return ans

