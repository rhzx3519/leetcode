class Solution(object):
    def minIncrementForUnique(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        A.sort()
        move = 0
        for i in range(1, len(A)):
            if A[i]<=A[i-1]:
                pre = A[i]
                A[i] = A[i-1] + 1
                move += A[i] - pre
        return move
        

if __name__ == '__main__':
    A = [3,2,1,2,1,7]
    su = Solution()
    print su.minIncrementForUnique(A)            