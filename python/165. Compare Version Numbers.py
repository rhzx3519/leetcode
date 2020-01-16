class Solution(object):
    def compareVersion(self, version1, version2):
        """
        :type version1: str
        :type version2: str
        :rtype: int
        """
        v1 = version1.split('.')
        v2 = version2.split('.')
        l1, l2 = len(v1), len(v2)
        i1 = i2 = 0
        while i1 < l1 and i2 < l2:
            a = int(v1[i1])
            b = int(v2[i2])
            if a != b:
                return a > b and 1 or -1
            i1 += 1
            i2 += 1
        
        while i1 < l1:
            if int(v1[i1]) != 0:
                return 1
            i1 += 1
            
        while i2 < l2:
            if int(v2[i2]) != 0:
                return -1
            i2 += 1
        return 0