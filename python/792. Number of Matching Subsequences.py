class Solution(object):
    def numMatchingSubseq(self, S, words):
        """
        :type S: str
        :type words: List[str]
        :rtype: int
        """

        def is_subseq(word):
            i = j = 0
            n = len(S)
            wn = len(word)
            if word[0] not in d: return False
            j = d[word[0]]
            if wn>n: return False
            while i < wn and j < n:
                if word[i]==S[j]:
                    i += 1
                j += 1
            return i==wn

        count = collections.defaultdict(int)
        for w in words:
            count[w] += 1

        d = {}
        for i, ch in enumerate(S):
            if ch not in d:
                d[ch] = i            

        res = 0
        for w in set(words):
            if is_subseq(w):
                res += count[w]
        return res
        # return len(filter(is_subseq, set(words)))
    