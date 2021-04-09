class Solution(object):
    def minReorder(self, n, connections):
        """
        :type n: int
        :type connections: List[List[int]]
        :rtype: int
        """
        ans = 0
        que = set([0])
        edges = connections
        while edges:
            tmp = []
            for a, b in edges:
                if b in que: # node which can reach 0
                    que.add(a)
                elif a in que: # node which can be reached by 0
                    que.add(b)
                    ans += 1
                else:
                    tmp.append((a, b))
            # print edges, tmp, que
            edges = tmp
        return ans