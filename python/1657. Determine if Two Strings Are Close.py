class Solution(object):
    def closeStrings(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: bool
        """
        n1, n2 = len(word1), len(word2)
        if n1 != n2: return False
        c1 = collections.Counter(word1)
        c2 = collections.Counter(word2)

        s1, s2 = set(word1), set(word2)
        return s1==s2 and sorted(c1.values())==sorted(c2.values())
