class Solution(object):
    def maxLength(self, arr):
        """
        :type arr: List[str]
        :rtype: int
        """
        self.max_len = 0
        valid = 0
        def dfs(idx, bit):
            if idx==len(arr):
                return 0
            tmp = bit
            for ch in arr[idx]:
                l = ord(ch) - ord('a')
                if tmp&(1<<l):
                    return dfs(idx+1, bit)
                tmp |= (1<<l)
            return max(len(arr[idx]) + dfs(idx+1, tmp), dfs(idx+1, bit))
        
        return dfs(0, 0)
    # dfs
