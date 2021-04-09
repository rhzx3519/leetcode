class Solution(object):
    def numSimilarGroups(self, strs):
        """
        :type strs: List[str]
        :rtype: int
        """
        n = len(strs)
        parent = [i for i in range(n)]

        def find(x):
            px = parent[x]
            if px != x:
                parent[x] = find(px)
            return parent[x]
        
        def union(x, y):
            px = find(x)
            py = find(y)
            if py != px:
                parent[py] = px
        
        def check(s1, s2):
            n = len(s1)
            cnt = 0
            for i in range(n):
                if s1[i] != s2[i]:
                    cnt += 1
                    if cnt > 2:
                        return False
            return True

        for i in range(n):
            for j in range(i+1, n):
                pi, pj = find(i), find(j)
                if pi == pj: continue
                if check(strs[i], strs[j]):
                    union(i, j)
        
        ans = 0
        for i in range(n):
            if parent[i] == i:
                ans += 1
        # print check('rats', 'arts')
        return ans




        