from collections import Counter, defaultdict
class Solution(object):
    def findSubstring(self, s, words):
        """
        :type s: str
        :type words: List[str]
        :rtype: List[int]
        """
        res = []
        if not words: return res
        mem = Counter(words)
        len_words = len(words)
        len_w = len(words[0])
        n = len(s)

        i = 0
        while i <= n - len_w*len_words:
            t = defaultdict(int)
            j = i
            while j < i + len_w*len_words:
                word = s[j:j+len_w]
                if word not in mem: break
                t[word] += 1
                j += len_w

            if t==mem:
                res.append(i)
            i += 1
        
        return res

s = "barfoofoobarthefoobarman"
words = ["bar","foo","the"]
su = Solution()
print su.findSubstring(s, words)