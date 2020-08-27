class Solution(object):
    def palindromePairs(self, words):
        """
        :type words: List[str]
        :rtype: List[List[int]]
        """
        res = []
        mem = {w[::-1]: i for i, w in enumerate(words)}

        def isPalindrome(s, l, r):
            sub = s[l:r+1]
            return sub[::-1]==sub
        
        def find(s, l, r):
            return mem.get(s[l:r+1], -1)
        
        for i, w in enumerate(words):
            m = len(w)
            for j in range(m+1):
                if isPalindrome(w, j,  m-1):
                    leftId = find(w, 0, j-1)
                    if leftId != -1 and i != leftId:
                        res.append([i, leftId])
                if j > 0 and isPalindrome(w, 0, j-1):
                    rightId = find(w, j, m-1)
                    if rightId != -1 and rightId != i:
                        res.append([rightId, i])
        return res

# 四种情况：

# 数组里有空字符串，并且数组里还有自己就是回文的字符串，每出现一个可与空字符串组成两对。
# 如果自己的翻转后的字符串也在数组里，可以组成一对，注意翻转后不能是自己。
# 如果某个字符串能找到一个分割点，分割点前的部分是回文，后半部分翻转后也在数组里，可组成一对。
# 把3反过来，如果后部分是回文，前半部分翻转后在数组里，可组成一对。

