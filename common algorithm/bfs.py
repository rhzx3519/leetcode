#!usr/bin/env python  
#-*- coding:utf-8 _*- 


class Solution(object):
    '''
        grid是一个m*n的矩阵，矩阵由0,1组成，0表示可以通过，1表示无法通过
        求start 到 end的最短路径，并输出, 你可以往上下左右四个方向行走
        example:
        grid = [[0, 1, 0, 0, 0],
                [0, 1, 0, 1, 0],
                [0, 0, 0, 1, 0],
                [0, 1, 1, 0, 0]]
        start = (0, 0) 左上角点
        end = (3, 4) 右下角店
        start -> end的最短路径为 
        [(0, 0), (1, 0), (2, 0), (2, 1), (2, 2), (1, 2), (0, 2), (0, 3), (0, 4), (1, 4), (2, 4), (3, 4)]
    '''
    def find(self, grid, start, end):
        '''
            graid: list[list[int]]
            start: list[int]
            end: list[int]
        '''
        offset = ((1, 0), (-1, 0), (0, 1), (0, -1))
        m, n = len(grid), len(grid[0])
        vis = [[0]*n for _ in range(m)]
        vis[start[0]][start[1]] = 1
        step = 0
        que = [start]
        while que:
            sz = len(que)
            while sz:
                sz -= 1
                p = que.pop(0)
                if p==end:
                    # print p, prev 
                    # print vis
                    break
                for dx, dy in offset:
                    nx = dx + p[0]
                    ny = dy + p[1]
                    if 0<=nx<m and 0<=ny<n and grid[nx][ny]!=1 and not vis[nx][ny]:
                        vis[nx][ny] = p
                        que.append((nx, ny))
            step += 1

        path = []
        if not vis[end[0]][end[1]]:
            return path
        p = end
        while p != start:
            path.append(p)
            p = vis[p[0]][p[1]]
        path.append(start)
        return path[::-1]
        

if __name__ == '__main__':
    grid = [[0, 1, 0, 0, 0],
            [0, 1, 0, 1, 0],
            [0, 0, 0, 1, 0],
            [0, 1, 1, 0, 0]]

    start = (0, 0)
    end = (3, 4)
    su = Solution()
    print su.find(grid, start, end)

