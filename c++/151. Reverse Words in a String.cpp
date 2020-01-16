class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        words = s.split()
        words = words[::-1];
        res = ''
        i = 0
        while i < len(words):
            if i == len(words) - 1:
                res += words[i]
            else:
                res += words[i] + ' '
            i += 1
        
        return res
        