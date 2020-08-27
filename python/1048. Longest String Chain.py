class Solution(object):
    def longestStrChain(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        words.sort(key=len)
        mem = {}
        for w in words:
            mem[w] = max(mem.get(w[:i] + w[i+1:], 0) + 1 for i in range(len(w)))
        return max(mem.values())

                