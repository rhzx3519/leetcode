class Solution(object):
    def maxUniqueSplit(self, s):
        """
        :type s: str
        :rtype: int
        """
        self.ans = 0
        def dfs(idx, a):
            if idx==len(s):
                self.ans = max(self.ans, len(a))
                # print a
                return
            for i in range(idx+1, len(s)+1):
                sub = s[idx:i]
                if sub in a:
                    continue
                a.append(sub)
                dfs(i, a)
                a.pop()
        dfs(0, [])
        return self.ans
