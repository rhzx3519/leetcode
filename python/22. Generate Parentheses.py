class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        res = []
        self.dfs(n, n, '', res)
        return res
        
    def dfs(self, l, r, s, res):
        if l > r:
            return
        if l == r and l == 0:
            res.append(s)
        if l > 0:
            self.dfs(l-1, r, s + '(', res)
        if r > 0:
            self.dfs(l, r-1, s + ')', res)