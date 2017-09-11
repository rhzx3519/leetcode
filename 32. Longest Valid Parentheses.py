class Solution(object):
    def longestValidParentheses(self, s):
        """
        :type s: str
        :rtype: int
        """
        res = 0
        
        i, ls = 0, len(s)
        mp = {}
        mp[0] = -1
        cnt = 0
        while i < ls:
            if s[i] == '(':
                cnt += 1
                if mp.get(cnt) == None:
                    mp[cnt] = i 
            else:
                if mp.get(cnt) != None:
                    mp.pop(cnt)
                cnt -= 1
                if mp.get(cnt) == None:
                    mp[cnt] = i
                else:
                    res = max(res, i - mp[cnt])
            print(i, mp)
            i += 1

        return res

s = Solution()
print(s.longestValidParentheses('()'))         
