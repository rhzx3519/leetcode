# """
# This is MountainArray's API interface.
# You should not implement it, or speculate about its implementation
# """
#class MountainArray(object):
#    def get(self, index):
#        """
#        :type index: int
#        :rtype int
#        """
#
#    def length(self):
#        """
#        :rtype int
#        """

class Solution(object):
    def findInMountainArray(self, target, mountain_arr):
        """
        :type target: integer
        :type mountain_arr: MountainArray
        :rtype: integer
        """
        p = self.findPeak(mountain_arr)
        left = self.findLeft(target, mountain_arr, 0, p)
        right = self.findRight(target, mountain_arr, p, mountain_arr.length()-1)
        return left if left!=-1 else right

    def findLeft(self, target, mountain_arr, ll, rr):
        l = ll
        r = rr
        while l<=r:
            m = l + (r-l)/2
            k = mountain_arr.get(m)
            if k==target:
                return m
            elif k < target:
                l = m + 1
            else:
                r = m - 1
        return -1

    def findRight(self, target, mountain_arr, ll, rr):
        l = ll
        r = rr
        while l<=r:
            m = l + (r-l)/2
            k = mountain_arr.get(m)
            if k==target:
                return m
            elif k < target:
                r = m - 1
            else:
                l = m + 1
        return -1

        
    def findPeak(self, mountain_arr):
        l = 0
        r = mountain_arr.length()-1
        while l<r:
            m1 = l + (r-l)/2
            m2 = m1 + 1
            if mountain_arr.get(m1) < mountain_arr.get(m2):
                l = m2
            else:
                r = m1
        return l
