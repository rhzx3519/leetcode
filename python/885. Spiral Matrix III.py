class Solution(object):
    def spiralMatrixIII(self, R, C, r0, c0):
        """
        :type R: int
        :type C: int
        :type r0: int
        :type c0: int
        :rtype: List[List[int]]
        """
        ans = [[r0, c0]]
        steps = [1, 0, 1, 0]
        offset = ((0, 1), (1, 0), (0, -1), (-1, 0))
        K = 4
        step = 0
        d = 0
        pos = [r0, c0]
        while len(ans) < R*C:
            d %= K
            step += steps[d]
            for i in range(step):
               pos[0] += offset[d][0]
               pos[1] += offset[d][1]
               if 0<=pos[0]<R and 0<=pos[1]<C:
                   ans.append(pos[:])
            d += 1
        return ans
        