class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        res = []
        candidates.sort()
        a = candidates
        self.dfs(a, 0, [], target, res)
        return res
        
    def dfs(self, a, idx, nums, target, res):
        if 0 == target:
            res.append(nums[:])
            return
        
        i = idx
        while i < len(a) and target >=  a[i]:
            nums.append(a[i])
            self.dfs(a, i + 1, nums, target - a[i], res)
            nums.pop()
            while i+1 < len(a) and a[i+1] == a[i]:
                i += 1
                
            i += 1
        