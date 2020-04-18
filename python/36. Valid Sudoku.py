class Solution(object):
    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """
        n = len(board)
        for i in range(n):
            r = []
            c = []
            for j in range(n):
                if board[i][j]!='.':
                    if board[i][j] in r:
                        return False
                    r.append(board[i][j])

                if board[j][i]!='.':
                    if board[j][i] in c:
                        print board[j][i], c                    
                        return False
                    c.append(board[j][i])

        for i in range(3):
            for j in range(3):
                t = []
                for m in range(i*3, (i+1)*3, 1):
                    for n in range(j*3, (j+1)*3, 1):
                        if board[m][n]=='.': continue
                        if board[m][n] in t:
                            return False
                        t.append(board[m][n])
        return True

        
board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
s = Solution()
print(s.isValidSudoku(board))