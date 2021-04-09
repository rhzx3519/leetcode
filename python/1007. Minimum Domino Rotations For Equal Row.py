class Solution(object):
    def minDominoRotations(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        up = [0] * 7
        down = [0] * 7
        N = len(A)
        count = [N] * 7
        for i in range(N):
            if A[i]==B[i]:
                count[A[i]] -= 1
            else:
                up[A[i]] += 1
                down[B[i]] += 1
        # print up
        # print down
        # print count
        ans = float('inf')
        for i in range(1, 7):
            if up[i] + down[i] >= count[i]:
                ans = min(up[i], down[i], ans)
        return -1 if ans==float('inf') else ans
        