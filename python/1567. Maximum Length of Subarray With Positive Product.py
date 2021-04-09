class Solution(object):
    def getMaxLen(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        ans = 0
        neg = pos = 0
        first_neg_idx = -1
        for i, num in enumerate(nums):
            if num==0:
                neg = pos = 0
                first_neg_idx = -1
                continue
            if num > 0:
                pos += 1
            else:
                if first_neg_idx==-1:
                    first_neg_idx = i 
                neg += 1
            if neg&1==0:
                ans = max(neg + pos, ans)
            else:
                ans = max(ans, i - first_neg_idx)
        return ans

# 统计子数组中正数和负数的数目，碰到0就重置，记录第一个负数出现的位置.