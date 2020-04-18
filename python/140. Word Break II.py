class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: List[str]
        """
        mem = {}

        def dfs(s):
            if mem.has_key(s): return mem[s]
            if len(s)==0:
                return ['']

            res = []
            for word in wordDict:
                if len(word)>len(s): continue
                if s[:len(word)]==word:
                    sub_res = dfs(s[len(word):])
                    for sub in sub_res:
                        if not sub:
                            res.append(word)
                        else:
                            res.append(word + ' ' + sub)
            mem[s] = res
            return res

        return dfs(s)

if __name__ == '__main__':
    s = "catsanddog"
    wordDict = ["cat","cats","and","sand","dog"]
    su = Solution()
    print su.wordBreak(s, wordDict)