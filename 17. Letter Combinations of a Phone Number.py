class Solution(object):
    def __init__(self):
        self.visit = []
        self.mp = ['', '', 'abc', 'def', 'ghi', 'jkl', 'mno', 'pqrs', 'tuv', 'wxyz']
        
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        res = []
        self.visit = [False for x in range(len(digits))]
        self.dfs(digits, '', 0, res)
        return res
        
    def dfs(self, digits, num, depth, res):
        if depth == len(digits) and num:
            res.append(num)
            return
        i = len(num)
        while i < len(digits):
            if not self.visit[i]:
                self.visit[i] = True
                for x in self.mp[int(digits[i])]:
                    self.dfs(digits, num+x, depth+1, res)
                self.visit[i] = False
            i += 1
                    
s = Solution()
print s.letterCombinations('23')        