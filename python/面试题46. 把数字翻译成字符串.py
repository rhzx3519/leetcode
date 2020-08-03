class Solution(object):
    def translateNum(self, num):
        """
        :type num: int
        :rtype: int
        """
        if num<10:
            return 1
        ba = num%100
        if ba<10 or ba>25:
            return self.translateNum(num/10)
        else:
            return self.translateNum(num/10) + self.translateNum(num/100)