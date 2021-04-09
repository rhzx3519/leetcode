class Solution(object):
    def findAndReplacePattern(self, words, pattern):
        """
        :type words: List[str]
        :type pattern: str
        :rtype: List[str]
        """
        def check(w):
            if len(w) != len(pattern):
                return False
            w2p = {}
            p2w = {}
            for i in range(len(w)):
                if w[i] not in w2p and pattern[i] not in p2w:
                    w2p[w[i]] = pattern[i]
                    p2w[pattern[i]] = w[i]
                elif w[i] in w2p and pattern[i] in p2w:
                    if w2p[w[i]] != pattern[i] or p2w[pattern[i]] != w[i]:
                        return False
                else:
                    return False
            return True
        
        return filter(check, words)
            