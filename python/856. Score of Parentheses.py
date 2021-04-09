class Solution(object):
    def scoreOfParentheses(self, S):
        """
        :type S: str
        :rtype: int
        """
        ans = 0
        depth = 0
        for i, ch in enumerate(S):
            if ch=='(':
                depth += 1
            else:
                depth -= 1
            
            if S[i]==')' and S[i-1]=='(':
                ans += 1<<depth
        return ans