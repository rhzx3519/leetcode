#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def longestMountain(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        res = 0
        n = len(A)
        left = [0]*n
        right = [0]*n
        for i in range(1, n):
            if A[i] > A[i-1]:
                left[i] = left[i-1] + 1
            else:
                left[i] = 0
        for i in range(n-2, -1, -1):
            if A[i] > A[i+1]:
                right[i] = right[i+1] + 1
            else:
                right[i] = 0
        print left, right
        for i in range(1, n-1):
            l, r = left[i], right[i]
            if l>=1 and r>=1:
                res = max(res, l + r + 1)
        return res

if __name__ == '__main__':
	A = [2,1,4,7,3,2,5]
	su = Solution()
	print su.longestMountain(A)