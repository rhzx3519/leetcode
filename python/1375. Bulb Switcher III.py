#!usr/bin/env python  
#-*- coding:utf-8 _*-  


class Solution(object):
    def numTimesAllBlue(self, light):
        """
        :type light: List[int]
        :rtype: int
        """
        n = len(light)
        bubble = [0] * n
        blue = [0] * n
        cnt = ans = 0
        left = 0 # 从左往右数第一个不亮的灯泡
        for num, idx in enumerate(light):
            i = idx - 1
            bubble[i] = 1
            if i <= left:
                j = i
                while j < n and bubble[j]:
                    if blue[j] == 0:
                        cnt += 1
                        blue[j] = 1
                    j += 1
                    left = j
                    
            if num+1 == cnt:
                ans += 1
        return ans

            
if __name__ == '__main__':
    # light = [2,1,3,5,4]
    # light = [3,2,4,1,5]
    # light = [4,1,2,3]
    # light = [2,1,4,3,6,5]
    light = [1,2,3,4,5,6]
    su = Solution()
    su.numTimesAllBlue(light)