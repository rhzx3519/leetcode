#!usr/bin/env python  
#-*- coding:utf-8 _*-  


# 数组 A 是 [0, 1, ..., N - 1] 的一种排列，N 是数组 A 的长度。全局倒置指的是 i,j 满足 0 <= i < j < N 并且 A[i] > A[j] ，局部倒置指的是 i 满足 0 <= i < N 并且 A[i] > A[i+1] 。

# 当数组 A 中全局倒置的数量等于局部倒置的数量时，返回 true 。

#  

# 示例 1:

# 输入: A = [1,0,2]
# 输出: true
# 解释: 有 1 个全局倒置，和 1 个局部倒置。
# 示例 2:

# 输入: A = [1,2,0]
# 输出: false
# 解释: 有 2 个全局倒置，和 1 个局部倒置。
# 注意:

# A 是 [0, 1, ..., A.length - 1] 的一种排列
# A 的长度在 [1, 5000]之间
# 这个问题的时间限制已经减少了。

class Solution(object):
    def isIdealPermutation(self, A):
        """
        :type A: List[int]
        :rtype: bool
        """
        for i in range(len(A)):
            if A[i] != i and abs(A[i]-i)>1:
                return False
        return True

# # 审题!数列排序后就就是0,1,2.......N,刚好和下标是对应的.如果一个数和他的下标偏移量超过了1,即是和他的有序排列偏移超过了1,那么全局偏移和局部偏移必然不相等
