class Solution(object):
    def ambiguousCoordinates(self, S):
        """
        :type S: str
        :rtype: List[str]
        """
        S = S[1:-1]
        n = len(S)

        def dfs(x):
            lx = len(x)
            if not x or (lx>1 and x[0]=='0' and x[-1]=='0'):
                return []
            if x=='0': return [x[0]]
            if x[0]=='0' and lx > 1:
                return [x[0] + '.' + x[1:]]
            if x[-1]=='0' and lx > 1:
                return [x]
            return [x] + [x[:i] + '.' + x[i:] for i in range(1, lx)]
        
        res = []
        for i in range(1, n):
            x = S[:i]
            y = S[i:]
            ax = dfs(x)
            ay = dfs(y)
            for i in ax:
                for j in ay:
                    if i and j:
                        res.append('({}, {})'.format(i, j))
        return res
