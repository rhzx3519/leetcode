class Solution(object):
    def calcEquation(self, equations, values, queries):
        """
        :type equations: List[List[str]]
        :type values: List[float]
        :type queries: List[List[str]]
        :rtype: List[float]
        """   
        def query(a, b, val, vis):
            if a not in edges or b not in edges:
                return -1.0
            if a==b: return val
            for nx, v in edges[a]:
                if nx not in vis:
                    vis.add(nx)
                    ans = query(nx, b, val*v, vis)
                    if ans != -1.0:
                        return ans
                    vis.remove(nx)
            return -1.0

        edges = collections.defaultdict(list)
        n = len(equations)
        for i in range(n):
            a, b = equations[i]
            val = values[i]
            edges[a].append((b, val))
            edges[b].append((a, 1.0/val))
        
        ans = []
        for a, b in queries:
            vis = set([a])
            ans.append(query(a, b, 1.0, vis))
        
        return ans


        
            