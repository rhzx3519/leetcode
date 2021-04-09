class Solution(object):
    def minimumLength(self, s):
        """
        :type s: str
        :rtype: int
        """
        n = len(s)
        l, r = 0, n-1
        while l < r:
            if s[l] != s[r]: break
            
            i = l + 1
            while i < r and s[i]==s[l]:
                i += 1
            l = i

            j = r - 1
            while j > l and s[j]==s[r]:
                j -= 1
            r = j

        print s[l:r+1]
        return len(s[l:r+1])

if __name__ == '__main__':
    s = 'ca'
    s = "cabaabac"
    s = "aabccabba"
    su = Solution()
    su.minimumLength(s)