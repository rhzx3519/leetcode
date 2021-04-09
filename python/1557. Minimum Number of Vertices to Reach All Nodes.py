class Solution(object):
    def findSmallestSetOfVertices(self, n, edges):
        """
        :type n: int
        :type edges: List[List[int]]
        :rtype: List[int]
        """
        ans = []
        ind = [0]*n
        for a, b in edges:
            ind[b] += 1
        for i, t in enumerate(ind):
            if t==0:
                ans.append(i)
        return ans

            