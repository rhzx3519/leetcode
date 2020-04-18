class Solution(object):
    def gcdOfStrings(self, str1, str2):
        """
        :type str1: str
        :type str2: str
        :rtype: str
        """
        n1, n2 = len(str1), len(str2)
        if n1 < n2: return self.gcdOfStrings(str2, str1)
        if str1[:n2]!=str2: return ''

        if str2:
            return self.gcdOfStrings(str2, str1[n2:])
        else:
            return str1