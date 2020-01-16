import math

class Solution(object):
    def __init__(self):
        self.row = []
        self.col = []
        self.block = []
        self.ns = 0
        self.n = 0
        self.pos = []
        
    def mark(self, i, j, d):
        self.row[i] ^= 1 << d - 1 # mark row[i]
        self.col[j] ^= 1 << d - 1 # mark col[j]
        self.block[i/self.ns*self.ns + j/self.ns] ^= 1 << d - 1 # mark block[i/self.ns*self.ns + j/self.ns]
        
    def solveSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        self.n = len(board)
        self.ns = int(math.sqrt(self.n))
        self.row = [0 for x in range(self.n)]
        self.col = [0 for x in range(self.n)]
        self.block = [0 for x in range(self.n)]
        
        i = 0
        while i < self.n:
            j = 0
            while j < self.n:
                t = board[i][j]
                if t == '.':
                    self.pos.append(i*self.n + j)
                else:
                    self.mark(i, j, ord(t) - ord('0'))
                j += 1
            i += 1
        
        self.dfs(board)
        
    
    def dfs(self, board):
        if not self.pos:
            return True
        idx = self.pos[-1]
        x = idx / self.n
        y = idx % self.n
        for bt in range(1, self.n + 1):
            if self.row[x] & 1<<bt-1:
                continue
            if self.col[y] & 1<<bt-1:
                continue
            if self.block[x/self.ns*self.ns + y/self.ns] & 1<<bt-1:
                continue
            self.pos.pop()
            self.mark(x, y, bt)
            board[x][y] = chr(bt + ord('0'))
            if self.dfs(board):
                return True
            else:
                board[x][y] = '.'
                self.pos.append(idx)
                self.mark(x, y, bt)
            
        return False
        
        