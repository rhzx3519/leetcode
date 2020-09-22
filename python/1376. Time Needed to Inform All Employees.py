class Solution(object):
    def numOfMinutes(self, n, headID, manager, informTime):
        """
        :type n: int
        :type headID: int
        :type manager: List[int]
        :type informTime: List[int]
        :rtype: int
        """
        que = []
        edges = [[] for _ in range(n)]
        for i, head in enumerate(manager):
            if head==-1:
                que.append((i, informTime[i]))
                continue
            edges[head].append(i)
        ans = 0
        while que:
            i, t = que.pop(0)
            ans = max(ans, t)
            for sub in edges[i]:
                que.append((sub, informTime[sub] + t))
        return ans