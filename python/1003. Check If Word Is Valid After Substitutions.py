class Solution(object):
    def isValid(self, S):
        """
        :type S: str
        :rtype: bool
        """
        while S != '' and 'abc' in S:
            S = S.replace('abc', '')
        return not S