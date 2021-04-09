class Solution(object):
    def expressiveWords(self, S, words):
        """
        :type S: str
        :type words: List[str]
        :rtype: int
        """
        def check(W):
            ns, nw = len(S), len(W)
            i = j = 0
            l = 1
            while i < nw and j < ns:
                if W[i] != S[j]:
                    return False
                i += 1
                j += 1
                stepw = steps = 0
                while i < nw and W[i-1]==W[i]:
                    stepw += 1
                    i += 1
                while j < ns and S[j-1]==S[j]:
                    steps += 1
                    j += 1
                if stepw > steps:
                    return False
                if steps > stepw and steps < 2:
                    return False
            return i==nw and j==ns

        return len(filter(check, words))
        
if __name__ == '__main__':
    S = "heeellooo"
    words = ["hello", "hl", "helo"]
    su = Solution()
    print su.expressiveWords(S, words)