class Solution(object):
    def intervalIntersection(self, A, B):
        """
        :type A: List[List[int]]
        :type B: List[List[int]]
        :rtype: List[List[int]]
        """
        res = []
        na, nb = len(A), len(B)
        i = j = 0
        while i<na and j<nb:
            a = A[i]
            b = B[j]
            if b[1]>=a[0] and a[1]>=b[0]:
                res.append([max(a[0], b[0]), min(a[1], b[1])])
            if a[1]<b[1]: i += 1
            else: j += 1
        return res

        
if __name__ == '__main__':
    A = [[0,2],[5,10],[13,23],[24,25]]
    B = [[1,5],[8,12],[15,24],[25,26]]
    su = Solution()
    print su.intervalIntersection(A, B)        
