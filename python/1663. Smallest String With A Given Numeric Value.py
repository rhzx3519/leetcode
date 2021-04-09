#!usr/bin/env python  
#-*- coding:utf-8 _*- 

class Solution(object):
    def getSmallestString(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        ans = []
        count_z = (k-n)/25
        count_x = (k-n)%25
        count_a = max(0, n - count_z - int(count_x > 0))
        print count_a, count_x, count_z
        return ''.join(['a']*count_a + [chr(ord('a') + count_x)]*int(count_x>0) + ['z']*count_z)

if __name__ == '__main__':
    n = 9
    k = 34
    n = 5
    k = 130
    n = 3
    k = 27
    n = 5
    k = 73
    su = Solution()
    print su.getSmallestString(n, k)


# 思路：字符串由x个a，y个z，还有1或者0个其他字符构成. 