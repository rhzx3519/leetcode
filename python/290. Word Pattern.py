class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        a = str.split(' ')
        n = len(a)
        if n != len(pattern):
            return False
        for i in range(n):
            if pattern.index(pattern[i]) != a.index(a[i]):
                return False
        return True
        
s = Solution()
print(s.wordPattern("abc", "b c a"))