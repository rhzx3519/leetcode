class Solution(object):
    def solveSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: None Do not return anything, modify board in-place instead.
        """
        if not board: return res
        N = 9

        def valid(row, col, num):
            for i in range(N):
                if board[row][i]==num or board[i][col]==num:
                    return False
            for i in range(3):
                for j in range(3):
                    if board[row//3*3+i][col//3*3+j]==num:
                        return False
            return True
      

        def backtrace(idx):
            if idx==N**2:
                return
            x, y = idx//N, idx%N
            if board[x][y]!='.':
                return backtrace(idx+1)
            for num in range(1, 10):
                if valid(x, y, str(num)):
                    board[x][y] = str(num)
                    backtrace(idx+1)
                    board[x][y] = '.'
                    
        backtrace(0)

if __name__ == '__main__':
    board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
    su = Solution()
    su.solveSudoku(board)     
    print board
            
            
        
        