class Solution(object):
    def updateBoard(self, board, click):
        """
        :type board: List[List[str]]
        :type click: List[int]
        :rtype: List[List[str]]
        """
        m, n = len(board), len(board[0])
        vis = [[0]*n for _ in range(m)]
        offset = ((1, 0), (-1, 0), (0, 1), (0, -1),
                    (1, 1), (1, -1), (-1, 1), (-1, -1))
        over = [False]
        def dfs(x, y):
            if over[0]:
                return
            if board[x][y] not in ('E', 'M'):
                return
            if board[x][y] == 'M':
                board[x][y] = 'X'
                over[0] = True
                return
            count = 0
            for dx, dy in offset:
                nx = x + dx
                ny = y + dy
                if 0<=nx<m and 0<=ny<n and board[nx][ny]=='M':
                    count += 1
            if count > 0:
                board[x][y] = str(count)
                return
            board[x][y] = 'B'
            for dx, dy in offset:
                nx = x + dx
                ny = y + dy
                if 0<=nx<m and 0<=ny<n and board[nx][ny]=='E':
                    dfs(nx, ny)
        
        x, y = click
        dfs(x, y)
        return board
