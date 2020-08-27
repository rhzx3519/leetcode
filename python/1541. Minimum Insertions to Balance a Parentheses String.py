class Solution(object):
    def minInsertions(self, s):
        """
        :type s: str
        :rtype: int
        """
        ans = 0
        left = 0    # 记录左括号的数量
        n = len(s)
        i = 0
        while i < n:
            if s[i] == '(': 
                left += 1
            elif s[i] == ')':
                if i+1 < n and s[i+1]==')':
                    i += 1
                else:
                    ans += 1
                if left > 0:
                    left -= 1
                else:
                    ans += 1
            i += 1
        ans += left*2   # 每个左括号匹配两个右括号
        return ans