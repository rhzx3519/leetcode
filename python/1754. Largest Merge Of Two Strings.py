class Solution(object):
    def largestMerge(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        w = []
        w1, w2 = word1, word2
        n1, n2 = len(w1), len(w2)
        i = j = 0
        while i < n1 or j < n2:
            sub1 = w1[i:]
            sub2 = w2[j:]
            if sub1 >= sub2:
                w.append(sub1[0])
                i += 1
            else:
                w.append(sub2[0])
                j += 1
        return ''.join(w)
