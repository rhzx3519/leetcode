class Solution(object):
    def angleClock(self, hour, minutes):
        """
        :type hour: int
        :type minutes: int
        :rtype: float
        """
        T = 360
        UNIT = 30
        h = hour/12.0*T
        m = minutes/60.0*T
        h += minutes/60.0 * UNIT
        # print h, m
        degree = abs(h - m)
        if degree > 180:
            degree = T - degree
        return degree