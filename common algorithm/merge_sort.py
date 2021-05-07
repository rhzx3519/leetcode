#!usr/bin/env python  
#-*- coding:utf-8 _*- 

# 归并排序可以求逆序数

class Solution(object):

    def mergeSort(self, a):
        n = len(a)
        self.aux = [0]*n
        l, r = 0, n-1
        self.inverseNum = 0
        self.sort(a, l, r)

    def sort(self, a, l, r):
        if l >= r:
            return
        mid = l + (r-l)/2
        self.sort(a, l, mid)
        self.sort(a, mid+1, r)
        self.merge(a, l, mid, r)

    def merge(self, a, l, mid, r):
        i, j = l, mid + 1
        for k in range(l, r+1):
            self.aux[k] = a[k]
        for k in range(l, r+1):
            if i > mid:
                a[k] = self.aux[j]
                j += 1
            elif j > r:
                a[k] = self.aux[i]
                i += 1
            elif self.aux[i] <= self.aux[j]:
                a[k] = self.aux[i]
                i += 1
            else:
                a[k] = self.aux[j]
                j += 1
                self.inverseNum += mid - i + 1 #  aux[i:mid+1]之间的数都比aux[j]要大

if __name__ == '__main__':
    a = [3,4,2,5,1]
    su = Solution()
    su.mergeSort(a)
    print a, su.inverseNum