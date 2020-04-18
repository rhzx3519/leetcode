class Solution(object):

    def findLadders(self, beginWord, endWord, wordList):
        """
        :type beginWord: str
        :type endWord: str
        :type wordList: List[str]
        :rtype: List[List[str]]
        """
        res = []
        wordList = set(wordList)
        if endWord not in wordList: return res
        que = [[beginWord]]
        lv = 1
        min_lv = 1<<31 - 1
        last_lv_words = set([])
        while que:
            for w in last_lv_words:
                wordList.remove(w)
            last_lv_words = set([])
            # print lv, last_lv_words
            if lv > min_lv: break

            sz = len(que)
            while sz:
                sz -= 1
                path = que.pop(0)

                word = path[-1]
                if word==endWord:
                    if lv < min_lv:
                        min_lv = lv
                        res = [path]
                    elif lv==min_lv:
                        res.append(path)
  
                tmp = 1<<31-1
                for k in range(len(word)):
                    for i in range(26):     
                        ch = chr(ord('a') + i)
                        if ch == word[k]: continue
                        t = word[:k] + ch + word[k+1:]
                        if t in wordList and t not in path:
                            last_lv_words.add(t)
                            que.append(path + [t])

            lv += 1 
        return res



if __name__ == '__main__':
    beginWord = "hit"
    endWord = "cog"
    wordList = ["hot","dot","dog","lot","log","cog"]
    su = Solution()
    print su.findLadders(beginWord, endWord, wordList)
