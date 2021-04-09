class Solution(object):
    def matrixReshape(self, nums, r, c):
        """
        :type nums: List[List[int]]
        :type r: int
        :type c: int
        :rtype: List[List[int]]
        """
        m, n = len(nums), len(nums[0])
        if r*c != m*n: return nums
        now = [[0]*c for _ in range(r)]
        for i in range(m):
            for j in range(n):
                idx = i*n + j
                x = idx/c
                y = idx%c
                # print i, j, idx, x, y
                now[x][y] = nums[i][j]
        return now
