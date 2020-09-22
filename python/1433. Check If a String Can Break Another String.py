class Solution(object):
    def checkIfCanBreak(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: bool
        """
        if len(s1) != len(s2):
            return False
        n = len(s1)
        s1, s2 = list(s1), list(s2)
        s1.sort(reverse=True)
        s2.sort(reverse=True)
        f1 = f2 = True
        for x, y in zip(s1, s2):
            if x < y:
                f1 = False
            if x > y:
                f2 = False
            if not f1 and not f2:
                break
        return f1 or f2
        