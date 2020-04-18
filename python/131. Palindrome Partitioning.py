class Solution(object):
    def partition(self, s):
        """
        :type s: str
        :rtype: List[List[str]]
        """
        def is_palindrome(a):
            l, r = 0, len(a)-1
            while l<r:
                if a[l]!=a[r]: return False
                l += 1
                r -= 1
            return True
        
        res = []
        def dfs(l, r, a):
            if l >= r:
                res.append(a[:])
                return
            
            for i in range(l+1, r+1):
                if is_palindrome(s[l:i]):
                    a.append(s[l:i])
                    dfs(i, r, a)
                    a.pop()
        
        n = len(s)
        dfs(0, n, [])
        return res

if __name__ == '__main__':
    s = "aab"
    su = Solution()
    print su.partition(s)