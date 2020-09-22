class Solution(object):
    def minJumps(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        import collections
        d = collections.defaultdict(list)
        n = len(arr)
        for i in range(n):
            d[arr[i]].append(i)

        vis = [0] * n
        vis[0] = 1
        que = [(0, 0)]
        while que:
            i, k = que.pop(0)
            if i==n-1:
                return k
            for j in d[arr[i]] + [i+1, i-1]:
                if j<0 or j>n or vis[j]:
                    continue
                vis[j] = 1
                que.append((j, k+1))
            d[arr[i]] = []
        return -1