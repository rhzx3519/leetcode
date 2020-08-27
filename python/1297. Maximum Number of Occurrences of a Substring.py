import collections

class Solution(object):
    def maxFreq(self, s, maxLetters, minSize, maxSize):
        """
        :type s: str
        :type maxLetters: int
        :type minSize: int
        :type maxSize: int
        :rtype: int
        """
        l = r = 0
        n = len(s)
        mem = collections.defaultdict(int)
        while r <= n:
            t = r - l
            num = len(set(s[l:r]))
            

            # print s[l:r], l, r, len(set(s[l:r]))
            while minSize<=r-l<=maxSize:
                # print s[l:r], l, r
                if len(set(s[l:r])) <= maxLetters:
                    # print s[l:r], l, r
                    mem[s[l:r]] += 1 
                l += 1
                # t = r - l
                # num = len(set(s[l:r]))
            r += 1
        if not mem:
            return 0
        return max(mem.values())
        

if __name__ == '__main__':
    s = "aababcaab"
    maxLetters = 2
    minSize = 3
    maxSize = 4
    su = Solution()
    print su.maxFreq(s, maxLetters, minSize, maxSize)
