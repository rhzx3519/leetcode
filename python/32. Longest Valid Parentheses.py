class Solution(object):
    def longestValidParentheses(self, s):
        """
        :type s: str
        :rtype: int
        """
        n = len(s)
        dp = [0]*n
        st = []
        for i, t in enumerate(s):
            if t=='(':
                st.append(i)
            elif st:
                dp[st.pop()], dp[i] = 1, 1
        
        print dp
        max_val = c = 0
        for i in dp:
            if i:
                c += 1
            else:
                max_val = max(max_val, c)
                c = 0
        return max(max_val, c)

s = Solution()
print(s.longestValidParentheses('()'))         
