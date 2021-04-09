#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class ShortestPathFasterAlgorithm(object):
    '''
        从图中的某个顶点出发到达另外一个顶点的所经过的边的权重和最小的一条路径，称为最短路径
        SPFA算法是求解单源最短路径问题的一种算法
    '''
    def spfa(self, n, s, t, edges):
        '''
            edges: (u, w, c), 表示起点u，终端w，费用c
        '''
        prev = [-1]*n
        adj = [[] for _ in range(n)]
        for u, w, c in edges:
            adj[u].append((w, c))
        cost = [float('inf')]*n
        cost[s] = 0
        que = [s]

        while que:
            u = que.pop(0)
            for w, c in adj[u]:
                if cost[u] + c < cost[w]:
                    prev[w] = u
                    cost[w] = cost[u] + c
                    if w not in que:
                        que.append(w)
            # print cost
        print cost
        print prev
        for i in range(n):
            st = [i]
            t = prev[i]
            while t != -1:
                st.append(t)
                t = prev[t]
            print st[::-1]

    def test(self):
        n = 6
        s = 0
        t = 5
        edges = [[0,4,30], [0,5,100], [0,2,10], [1,2,5], [2,3,50],
             [4,3,20], [4,5,60], [3,5,10]]
        self.spfa(n, s, t, edges)

class Edge(object):
    def __init__(self, frm, to, capacity, flow, cost):
        self.frm = frm
        self.to = to
        self.capacity = capacity
        self.flow = flow
        self.cost = cost

class FlowNetwork(object):

    def __init__(self, n):
        self.n = n
        self.adj = [[] for _ in range(n)]

    def addEdge(frm, to, capacity, flow, cost):
        edge = Edge(frm, to, capacity, flow, cost)
        self.adj[frm].append(edge)
        self.adj[to].append(edge)


class MinCostMaxFlowAlgorithm(object):


    def __init__(self):
        pass

    def 



                            
if __name__ == '__main__':
    su = ShortestPathFasterAlgorithm()
    su.test()


# https://blog.csdn.net/u014800748/article/details/44059993





