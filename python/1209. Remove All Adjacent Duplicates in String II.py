class Solution(object):
    def removeDuplicates(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """
        n = len(s)
        l = 1
        for i in range(1, n):
            if s[i]==s[i-1]:
                l += 1
            else:
                l = 1
            if l==k:
                return self.removeDuplicates(s[:i-k+1] + s[i+1:], k)
        return s        
        
if __name__ == '__main__':
    s = 'deeedbbcccbdaa'
    k = 3
    s = "pbbcggttciiippooaais"
    k = 2
    su = Solution()
    print su.removeDuplicates(s, k)