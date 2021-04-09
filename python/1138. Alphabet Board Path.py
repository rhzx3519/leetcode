class Solution(object):
    def alphabetBoardPath(self, target):
        """
        :type target: str
        :rtype: str
        """
        board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"]
        m, n = len(board), len(board[0])
        vis = [[-1]*len(board[0]) for _ in range(len(board))]
        dirt = 'RLUD'
        d = {(-1, 0): 'U', (1, 0): 'D', (0, -1): 'L', (0, 1): 'R'}

        def bfs(r, c, t):
            que = [(r, c)]
            while que:
                x, y = que.pop(0)
                # print x, y
                if board[x][y]==t:
                    return (x, y)
                for i, (dx, dy) in enumerate(((1, 0), (-1, 0), (0, 1), (0, -1))):
                    nx = x + dx
                    ny = y + dy
                    if nx<0 or nx>=m or ny<0 or ny>=len(board[nx]) or vis[nx][ny] != -1:
                        continue
                    vis[nx][ny] = (x, y)
                    que.append((nx, ny))
            return (-1, -1)

        def find(start, end):
            prev = [end]
            while end != start:
                end = vis[end[0]][end[1]]
                prev.append(end)
            # print prev
            cmd = ['!']
            for i in range(1, len(prev)):
                k = (prev[i-1][0] - prev[i][0], prev[i-1][1] - prev[i][1])
                cmd.append(d[k])
            # print cmd
            return ''.join(cmd[::-1])
            
        
        ans = []
        r = c = 0
        for t in target:
            vis = [[-1]*n for _ in range(m)]
            end = bfs(r, c, t)
            # print vis
            path = find((r, c), end)
            r, c = end
            # print (r, c), t
            ans.append(path)

        # print ans
        return ''.join(ans)
        

if __name__ == '__main__':
    target = 'leet'
    su = Solution()
    su.alphabetBoardPath(target)