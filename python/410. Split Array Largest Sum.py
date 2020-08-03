class Solution(object):
    def splitArray(self, nums, m):
        """
        :type nums: List[int]
        :type m: int
        :rtype: int
        """
        mem = {} 
        suffixSum = [0]*len(nums)
        suffixSum[-1] = nums[-1]     # 后缀和
        for i in range(len(nums)-2, -1, -1):
            suffixSum[i] = suffixSum[i+1] + nums[i]

        def dp(idx, m): # 从下标i开始，分解成m个数组的最大值最小
            if m==1:
                return suffixSum[idx] # 返回nums[idx:]的和
            if (idx, m) in mem:
                return mem[(idx, m)]
            min_val = float('inf')
            s = 0
            for i in range(idx, len(nums)-m+1):
                s += nums[i]
                min_val = min(min_val, max(s, dp(i+1, m-1)))
            mem[(idx, m)] = min_val
            return min_val
            
        return dp(0, m)
# 思路：自顶向下的dp