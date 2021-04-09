class Solution(object):
    def baseNeg2(self, N):
        """
        :type N: int
        :rtype: str
        """
        ans = []
        while N:
            N, r = -(N//2), N%2
            ans.append(str(r))
        return ''.join(ans[::-1]) if ans else '0'