class Solution(object):
    def permutation(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        if not s:
            return []
        res = []
        def dfs(s, tmp):
            if not s:
                res.append(tmp)
                return
            mem = set()
            for i in range(len(s)):
                if s[i] in mem:
                    continue
                mem.add(s[i])
                dfs(s[:i] + s[i+1:], tmp + s[i])
        dfs(s, '')
        return res