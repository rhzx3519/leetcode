#!usr/bin/env python  
#-*- coding:utf-8 _*- 

class Solution(object):

    def search(self, A, target):
        l, r = 1, len(A)
        while l < r:
            mid = l + (r - l)/2
            if A[mid] <= target:
                l = mid + 1
            else:
                r = mid
        return l - 1

    def lower_bound(self, A, target):
        sz = len(A)
        l = 0
        while sz > 0:
            half = sz>>1
            mid = l + half
            if A[mid] < target:
                l = mid + 1
                sz = sz - half - 1
            else:
                sz = half
        return l


    def upper_bound(self, A, target):
        sz = len(A) - 1
        l = 0
        while sz > 0:
            half = sz>>1
            mid = l + half
            if A[mid] > target:
                sz = half
            else:
                l = mid + 1
                sz = sz - half - 1
        return l



so=Solution()
# s1=[1,3,5,6]
# s2=[1,3,5,5,5,6]
# s3 = [2]
s4 = [0,2,3]
target=-1
print(so.lower_bound(s4, target))
