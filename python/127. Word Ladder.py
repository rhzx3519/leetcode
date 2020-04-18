class Solution(object):
    def ladderLength(self, beginWord, endWord, wordList):
        """
        :type beginWord: str
        :type endWord: str
        :type wordList: List[str]
        :rtype: int
        """
        wordList = set(wordList)
        if endWord not in wordList:
            return 0
        que = [beginWord]
        lv = 1
        while que:
            sz = len(que)
            while sz:
                sz -= 1
                word = que.pop(0)
                if word==endWord: return lv
                for i in range(len(word)):
                    for j in range(26):
                        ch = chr(ord('a') + j)
                        w = word[:i] + ch + word[i+1:]
                        if w in wordList:
                            que.append(w)
                            wordList.remove(w)

            lv += 1
        return 0
# BFS，入队的条件是替换一个字母之后，单词在字典里，技巧是入队之后从字典中删除该单词
# 时间复杂度O(N), 空间复杂度O(N)