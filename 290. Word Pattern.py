class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        um1, um2 = {}, {}
        words = str.split()
        if len(words) != len(pattern):
            return False;
        i = 0
        n = len(words)
        while i < n:
            if um1.get(pattern[i]) == None and um2.get(words[i]) == None:
                um1[pattern[i]] = words[i]
                um2[words[i]] = pattern[i]
            elif um1.get(pattern[i]) == None or um2.get(words[i]) == None:
                return False
            elif um1[pattern[i]] != words[i] or um2[words[i]] != pattern[i]:
                return False
            i += 1
        
        return True
        
s = Solution()
print(s.wordPattern("abc", "b c a"))