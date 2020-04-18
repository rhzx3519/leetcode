class Solution(object):
    def rangeBitwiseAnd(self, m, n):
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        while m<n:
        	n &= (n-1)

        return n

     

if __name__ == '__main__':
	su = Solution()        	        
	print su.rangeBitwiseAnd(1,10)