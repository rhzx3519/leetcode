class Solution(object):
    def equationsPossible(self, equations):
        """
        :type equations: List[str]
        :rtype: bool
        """

        p = {}
        for i in range(26):
            ch = chr(ord('a') + i)
            p[ch] = ch

        def find(x):
            while x != p[x]:
                x = p[x]
            return p[x]
        
        def union(x, y):
            px = find(x)
            py = find(y)
            if px != py:
                p[py] = px
            return px

        equations.sort(key=lambda x: x[1:-1], reverse=True)

        for e in equations:
            a = e[0]
            b = e[-1]
            op = e[1:-1]
            if op=='==':
                union(a, b)
            else:
                pa = find(a)
                pb = find(b)
                if pa==pb:
                    return False
        return True

# 思路：并查集        
