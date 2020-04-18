class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        candidates.sort()
        res = []

        def dfs(idx, cur, a):
            if cur==target: 
                res.append(a[:])            

            if idx>=len(candidates):
                return

            i = idx
            while i < len(candidates) and cur<=target:
                a.append(candidates[i])
                dfs(i+1, cur+candidates[i], a)
                a.pop()
                while i+1 < len(candidates) and candidates[i+1]==candidates[i]:
                    i += 1
                i += 1
        a = []
        dfs(0, 0, a)
        return res
        
if __name__ == '__main__':
    su = Solution()
    print su.combinationSum2([10,1,2,7,6,1,5], 8)                    
        