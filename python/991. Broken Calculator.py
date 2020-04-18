class Solution(object):
    def brokenCalc(self, X, Y):
        """
        :type X: int
        :type Y: int
        :rtype: int
        """
        if X>=Y:
            return X-Y
        if Y%2==0:
            return self.brokenCalc(X,Y/2) + 1
        else:
            return self.brokenCalc(X,Y+1) + 1