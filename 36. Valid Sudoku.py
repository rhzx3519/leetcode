class Solution(object):
    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """
        mp = {}
        r, c = len(board), len(board[0])
        
        i = 0
        while i < r:
            j = 0
            mp.clear()
            while j < c:
                t = board[i][j]
                if t == '.':
                    j += 1
                    continue
                if mp.get(t) == None:
                    mp[t] = 1
                else:
                    return False
                j += 1
            i += 1
                    
        
        j = 0
        while j < c:
            i = 0
            mp.clear()
            while i < r:
                t = board[i][j]
                if t == '.':
                    i += 1
                    continue
                if mp.get(t) == None:
                    mp[t] = 1
                else:
                    return False
                i += 1
            j += 1
            
        i = 0
        while i*3 < r:
            j = 0
            while j*3 < c: #
                m = i*3
                mp.clear()
                while m < i*3 + 3:
                    n = j*3
                    while n < j*3 + 3:
                        t = board[m][n]
                        if t == '.':
                            n += 1
                            continue
                        if mp.get(t) == None:
                            mp[t] = 1
                        else:
                            return False
                        n += 1
                    m += 1
                j += 1
            i += 1
            
        return True
        
board = [".87654321","2........","3........","4........","5........","6........","7........","8........","9........"]
s = Solution()
print s.isValidSudoku(board)        