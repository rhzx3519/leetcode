class Solution(object):
    def complexNumberMultiply(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        a = a.split('+')
        b = b.split('+')
        a0 = int(a[0])
        a1 = int(a[1][:-1])

        b0 = int(b[0])
        b1 = int(b[1][:-1])

        c0 = a0*b0 - a1*b1
        c1 = a0*b1 + b0*a1
        c = '{}+{}i'.format(c0, c1)
        return c