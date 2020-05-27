class Solution(object):
    def isHappy(self, n):
        """
        :type n: int
        :rtype: bool
        """
        s = set([n])
        while n!=1:
            nums = list(str(n))
            t = 0
            for num in nums:
                t += int(num)*int(num)
            if t in s: return False
            n = t
        return True

if __name__ == '__main__':
    su = Solution()
    su.isHappy(2)        