#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class FenwickTree(object):
    """树状数组的下标为[1, n]"""
    def __init__(self, n):
        super(FenwickTree, self).__init__()
        self.sums = [0]*(n+1)

    def update(self, i, k):
        """更新i位置的值
        """
        while i < len(self.sums):
            self.sums[i] += k
            i += self.lowbit(i)

    def getsum(self, i):
        """求1到i的和
        """
        res = 0
        while i>0:
            res += self.sums[i]    
            i -= self.lowbit(i)
        return res

    def lowbit(self, x):
        return x&(-x)

class NumArray(object):

    def __init__(self, nums):
        """
        :type nums: List[int]
        """
        self.nums = nums
        self.tree = FenwickTree(len(nums))
        for i in range(len(nums)):
            self.tree.update(i+1, nums[i])


    def update(self, i, val):
        """
        :type i: int
        :type val: int
        :rtype: None
        """
        self.tree.update(i+1, val-self.nums[i])
        self.nums[i] = val


    def sumRange(self, i, j):
        """
        :type i: int
        :type j: int
        :rtype: int
        """
        return self.tree.getsum(j+1) - self.tree.getsum(i)


if __name__ == '__main__':
    nums = [1, 3, 5]
    obj = NumArray(nums)
    print obj.sumRange(0, 2)   
    obj.update(1, 2)            
    print obj.sumRange(0, 2)

