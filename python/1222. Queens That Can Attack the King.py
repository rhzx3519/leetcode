class Solution(object):
    def queensAttacktheKing(self, queens, king):
        """
        :type queens: List[List[int]]
        :type king: List[int]
        :rtype: List[List[int]]
        """
        offset = ((1, 0), (-1, 0), (0, 1), (0, -1),
                (-1, -1), (-1, 1), (1, -1), (1, 1))
        N = 8
        grid = [[0]*N for _ in range(N)]
        for x, y in queens:
            grid[x][y] = 1
        
        ans = []
        for i in range(N):
            tx, ty = offset[i]
            x, y = king
            while 0<=x<N and 0<=y<N:    # king walk to 8 direcitons until find queen
                if grid[x][y]:
                    ans.append([x, y])
                    break
                x += tx
                y += ty
        return ans