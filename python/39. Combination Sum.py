class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        sorted(candidates)
        
        res = []
        def dfs(idx, cur, a):
            if idx >= len(candidates):
                if cur==target:
                    res.append(a)      
                return
          
            k = 0
            while cur+k*candidates[idx] <= target:
                dfs(idx+1, cur+k*candidates[idx], a + k*[candidates[idx]])
                k += 1
                
        a = []
        dfs(0, 0, a)
        return res

        
        
s = Solution()
print s.combinationSum([2,3,6,7], 7)