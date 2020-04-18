class Solution(object):
    def maxProduct(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        def get_mask(word):
            i = 0
            for w in word:
                i |= 1<<(ord(w)-ord('a'))
            return i
        
        res = 0
        n = len(words)
        for i in range(n):
            for j in range(i+1, n):
                if get_mask(words[i]) & get_mask(words[j]) != 0:
                    continue
                res = max(res, len(words[i]) * len(words[j]))  

        return res       

if __name__ == '__main__':
    words = ["abcw","baz","foo","bar","xtfn","abcdef"]
    su = Solution()
    print su.maxProduct(words)        

0b010000000000000000000111 
0b100010000010000000100000
