class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        star = backup = None
        i = j = 0
        while i < len(s):
        	print i, j
        	if j < len(p) and p[j] == '*':
        		backup = i
        		j += 1
        		star = j
        	elif j < len(p) and (p[j] == '?' or s[i] == p[j]):
        		i += 1
        		j += 1
        	elif star:
        		j = star
        		backup += 1
        		i = backup
        	else:
        		break

        while j < len(p) and p[j] == '*':
        	j += 1

        return i == len(s) and j == len(p)

s = Solution()
print s.isMatch('aab', 'aa*')    

