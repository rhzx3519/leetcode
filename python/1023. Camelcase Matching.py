class Solution(object):
    def camelMatch(self, queries, pattern):
        """
        :type queries: List[str]
        :type pattern: str
        :rtype: List[bool]
        """
        def isMatch(word, p):
            nw = len(word)
            np = len(p)
            i = j = 0
            while i < nw and j < np:
                if word[i]==p[j]:
                    i += 1
                    j += 1
                elif ord('A')<=ord(word[i])<=ord('Z'):
                    return False
                else:
                    i += 1
            if j < np: return False
            while i < nw:
                if ord('A')<=ord(word[i])<=ord('Z'):
                    return False
                i += 1
            return True
        
        return [isMatch(w, pattern) for w in queries]
