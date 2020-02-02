class Solution(object):

    def __init__(self, radius, x_center, y_center):
        """
        :type radius: float
        :type x_center: float
        :type y_center: float
        """
        self.radius = radius
        self.x_center = x_center
        self.y_center = y_center

    def randPoint(self):
        """
        :rtype: List[float]
        """
        rad = random.random()*math.pi*2
        r = random.random()*self.radius*self.radius
        r = math.sqrt(r)
        return [self.x_center + r*math.cos(rad), 
            self.y_center + r*math.sin(rad)]
        


# Your Solution object will be instantiated and called as such:
# obj = Solution(radius, x_center, y_center)
# param_1 = obj.randPoint()