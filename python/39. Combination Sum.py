class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        res = []
        a = candidates
        self.dfs(a, 0, target, [], 0, res)
        return res
        
    def dfs(self, a, idx, target, nums, sum, res):
        if idx == len(a):
            if sum == target:
                t = nums[:]
                res.append(t)
            return
        
        i = 0
        while sum + a[idx]*i <= target:
            j = 0
            while j < i:
                nums.append(a[idx])
                j += 1
            self.dfs(a, idx + 1, target, nums, sum + a[idx]*i, res)
            
            j = 0
            while j < i:
                nums.pop()
                j += 1
            
            i += 1
        
        
s = Solution()
print s.combinationSum([2,3,6,7], 7)