class Solution(object):
    def reconstructMatrix(self, upper, lower, colsum):
        """
        :type upper: int
        :type lower: int
        :type colsum: List[int]
        :rtype: List[List[int]]
        """
        n = len(colsum)
        grid = [[0]*n for _ in range(2)]
        self.ans = []
        
        def dfs(idx, grid, upper, lower):
            if self.ans:
                return
            if idx == n:
                if upper==0 and lower==0:
                    if not self.ans:
                        for row in grid:  
                            self.ans.append(row[:])
                return
            if upper < 0 or lower < 0:
                return

            if colsum[idx]==0:
                dfs(idx+1, grid, upper,  lower)
            elif colsum[idx]==1:    
                
                if upper > 0:
                    grid[0][idx] = 1
                    dfs(idx+1, grid,  upper-1, lower)
                    grid[0][idx] = 0
                
                if lower > 0:
                    grid[1][idx] = 1
                    dfs(idx+1, grid, upper, lower-1)
                    grid[1][idx] = 0
            elif colsum[idx]==2:
                if upper > 0 and lower > 0:
                    grid[0][idx] = 1
                    grid[1][idx] = 1
                    dfs(idx+1, grid, upper-1, lower-1)

        dfs(0, grid, upper, lower)
        return self.ans


if __name__ == '__main__':
    upper = 4
    lower = 7
    colsum = [2,1,2,2,1,1,1]
    su = Solution()
    print su.reconstructMatrix(upper, lower, colsum)