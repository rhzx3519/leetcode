class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
        	return False
        
        a = []
        while x:
        	a.append(x%10)
        	x /= 10
        l, h = 0, len(a)-1
        while l < h:
        	if a[l] != a[h]:
        		return False
        	l += 1
        	h -= 1

        return True

s = Solution()
print s.isPalindrome(-2147447412)    