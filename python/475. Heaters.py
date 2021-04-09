#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def findRadius(self, houses, heaters):
        """
        :type houses: List[int]
        :type heaters: List[int]
        :rtype: int
        """
        # 取前面的索引, 也就是最小返回0
        def bi(a, target):
            l, r = 0, len(a)-1
            while l<=r:
                mid = (l+r)/2
                val = a[mid]
                if val == target:
                    return mid
                elif val > target:
                    r = mid - 1
                else:
                    l = mid + 1
            return max(0, r)
        
        ans = 0
        heaters.sort()
        for h in houses:
            i = bi(heaters, h)
            min_val = abs(h - heaters[i])
            if i + 1 < len(heaters):
                min_val = min(min_val, abs(heaters[i+1] - h))
            ans = max(ans, min_val)
        return ans


            
# 对于每个房屋，要么用前面的暖气，要么用后面的，二者取近的，得到距离；
# 对于所有的房屋，选择最大的上述距离。