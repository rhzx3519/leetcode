class Solution(object):
    def isNumber(self, s):
        """
        :type s: str
        :rtype: bool
        """
        n = len(s)
        if n==0:
            return False
        self.i = 0
        def scanInteger():
            if self.i>=n:
                return False
            if s[self.i] in ('+', '-'):
                self.i += 1
            return scanUnsignedInteger()
        
        def scanUnsignedInteger():
            start = self.i
            while self.i < n and s[self.i].isdigit():
                self.i += 1
            return self.i > start
 
        while self.i < n and s[self.i] == ' ':
            self.i += 1

        numeric = scanInteger()
        if self.i>=n:
            return numeric

        if s[self.i] == '.':
            self.i += 1
            numeric = scanUnsignedInteger() or numeric

        if self.i>=n:
            return numeric

        if s[self.i] in ('e', 'E'):
            self.i += 1
            numeric = scanInteger() and numeric
        
        while self.i < n and s[self.i] == ' ':
            self.i += 1
        return numeric and self.i==n
            

if __name__ == '__main__':
    s = '.9 '
    su = Solution()
    print su.isNumber(s)


