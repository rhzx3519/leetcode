class Solution(object):
    def longestSubsequence(self, arr, difference):
        """
        :type arr: List[int]
        :type difference: int
        :rtype: int
        解法1: dfs超时
        """
        def dfs(idx, a):
            if idx==len(arr):
                return 0
            if a and arr[idx] - a[-1] != difference:
                return dfs(idx+1, a)
            return max(1 + dfs(idx+1, a+[arr[idx]]), dfs(idx+1, a))
        
        return dfs(0, [])

    
            