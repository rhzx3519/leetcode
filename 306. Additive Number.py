class Solution(object):
    def isAdditiveNumber(self, num):
        """
        :type num: str
        :rtype: bool
        """
        n = len(num)
        for i, j in itertools.combinations(range(1, n), 2):
            a, b = num[:i], num[i:j]
            if a != str(int(a)) or b != str(int(b)):
                continue
            while j < n:
                c = str(int(a) + int(b))
                if num.startswith(c, j):
                    j += len(c)
                    a, b = b, c
                else:
                    break
            if j == n:
                return True
        return False