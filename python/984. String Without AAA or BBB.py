class Solution(object):
    def strWithout3a3b(self, A, B):
        """
        :type A: int
        :type B: int
        :rtype: str
        """
        a, b = 'a', 'b'
        arr = []        
        if A < B:
            A, B = B, A
            a, b = b, a


        while A > 0 or B > 0:
            if A > 0:
                arr.append(a)
                A -= 1
            if A > B:
                arr.append(a)
                A -= 1
            if B > 0:
                arr.append(b)
                B -= 1
        return ''.join(arr)