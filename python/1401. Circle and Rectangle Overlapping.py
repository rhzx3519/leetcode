class Solution(object):
    def checkOverlap(self, radius, x_center, y_center, x1, y1, x2, y2):
        """
        :type radius: int
        :type x_center: int
        :type y_center: int
        :type x1: int
        :type y1: int
        :type x2: int
        :type y2: int
        :rtype: bool
        """
        x3 = (x1 + x2)*0.5
        y3 = (y1 + y2)*0.5
        p = (x_center, y_center)
        c = (x3, y3)
        v = (abs(x_center - x3), abs(y_center - y3))
        h = (x2 - x3, y2 - y3)
        t = (max(0, v[0] - h[0]), max(0, v[1] - h[1]))
        return t[0]**2 + t[1]**2 <= radius**2


        
