class Solution(object):
    def canReach(self, arr, start):
        """
        :type arr: List[int]
        :type start: int
        :rtype: bool
        """
        n = len(arr)
        vis = [0]*n
        def dfs(idx):
            vis[idx] = 1
            if arr[idx]==0:
                return True
            for t in (-1, 1):
                nx = idx + arr[idx]*t
                if nx>=0 and nx<len(arr) and vis[nx]==0 and dfs(nx):
                    return True
            return False
        return dfs(start)
