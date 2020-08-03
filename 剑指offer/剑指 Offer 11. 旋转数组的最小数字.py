#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def minArray(self, numbers):
        """
        :type numbers: List[int]
        :rtype: int
        """
        l, r = 0, len(numbers)-1
        while l < r:
            m = (l+r)>>1
            if numbers[m] > numbers[r]:
                l = m+1
            elif numbers[m] < numbers[r]:
                r = m
            else:
                r -= 1
        return numbers[l]

# 和 I 的做法类似, 都是二分法, 每次进入无序的那部分找出最小值
# 但是由于有重复值的情况, 需要加入 mid 元素等于 hi 元素的情况
# 此时应该将 hi 减 1 防止重复数字是最小元素        

if __name__ == '__main__':
    numbers = [2,2,2,0,1]
    su = Solution()
    su.minArray(numbers)        