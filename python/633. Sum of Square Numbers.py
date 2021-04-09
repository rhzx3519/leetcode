class Solution(object):
    def judgeSquareSum(self, c):
        """
        :type c: int
        :rtype: bool
        """
        i = 0
        j = int(math.sqrt(c))
        while i<=j:
            total = i*i + j*j
            if total==c:
                return True
            elif total<c:
                i += 1
            else:
                j -= 1
        return False