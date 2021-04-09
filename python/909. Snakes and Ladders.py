class Solution(object):
    def snakesAndLadders(self, board):
        """
        :type board: List[List[int]]
        :rtype: int
        """
        def coordinate(pos, N):
            pos -= 1
            x = pos / N
            if x&1:
                y = N - pos%N - 1
            else:
                y = pos%N

            return N-1-x, y


        N = len(board)
        step = 0
        dis = [float('inf')]*(N*N + 1)
        dis[1] = 0
        que = [1]
        while que:
            sz = len(que)
            while sz:
                sz -= 1
                pos = que.pop(0)
                if pos >= N*N:
                    return dis[pos]
                i = 1
                while i<=6 and pos+i<=N*N:
                    x, y = coordinate(pos+i, N)
                    npos = pos+i if board[x][y]==-1 else board[x][y]
                    if dis[pos] + 1 < dis[npos]:
                        dis[npos] = dis[pos] + 1
                        que.append(npos)
                    i += 1


        return -1
                    




if __name__ == '__main__':
    # N = 6
    # pos = 33
    # print coordinate(pos, N)
    board = [[-1,-1,-1,46,47,-1,-1,-1],
            [51,-1,-1,63,-1,31,21,-1],
            [-1,-1,26,-1,-1,38,-1,-1],
            [-1,-1,11,-1,14,23,56,57],
            [11,-1,-1,-1,49,36,-1,48],
            [-1,-1,-1,33,56,-1,57,21],
            [-1,-1,-1,-1,-1,-1,2,-1],
            [-1,-1,-1,8,3,-1,6,56]]

    su = Solution()
    print su.snakesAndLadders(board)










