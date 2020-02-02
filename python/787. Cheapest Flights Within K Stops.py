from Queue import PriorityQueue

class Comparable(object):
    def __init__(self, u, w, k):
        self.u = u
        self.w = w
        self.k = k

    def __cmp__(self, other):
        if self.w < other.w:
            return -1
        elif self.w==other.w:
            return 0
        else:
            return 1

class Solution(object):
    def findCheapestPrice(self, n, flights, src, dst, K):
        """
        :type n: int
        :type flights: List[List[int]]
        :type src: int
        :type dst: int
        :type K: int
        :rtype: int
        """
        INF = 1<<31 - 1
        grid = [[INF]*n for _ in range(n)]
        for flight in flights:
            u = flight[0]
            v = flight[1]
            w = flight[2]
            grid[u][v] = w
        
        pq = PriorityQueue()
        pq.put(Comparable(src, 0, 0))
        while pq.qsize:
            t = pq.get()
            if dst==t.u: return t.w
            if t.k <= K:
                for i, v in enumerate(grid[t.u]):
                    if v!=INF:
                        pq.put(Comparable(i, t.w+v, t.k+1))
        return -1

       

if __name__ == '__main__':
    n = 3
    flights = [[0,1,100],[1,2,100],[0,2,500]]
    src = 0
    dst = 2
    k = 1

    su = Solution()
    print su.findCheapestPrice(n, flights, src, dst, k)    


