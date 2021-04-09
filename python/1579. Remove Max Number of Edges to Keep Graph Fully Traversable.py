class Uset(object):
    def __init__(self, n):
        self.n = n
        self.look_up = [i for i in range(n+1)]

    def find(self, x):
        px = self.look_up[x]
        if px!=x:
            self.look_up[x] = self.find(px)
        return self.look_up[x]
    
    def union(self, x, y):
        px = self.find(x)
        py = self.find(y)
        if px!=py:
            self.look_up[py] = px

    def clone(self):
        tmp = Uset(self.n)
        tmp.look_up = self.look_up[:]
        return tmp

    def connected(self):
        ans = 0
        for i in range(1, self.n+1):
            if i == self.look_up[i]:
                ans += 1
        return ans
    

class Solution(object):
    def maxNumEdgesToRemove(self, n, edges):
        """
        :type n: int
        :type edges: List[List[int]]
        :rtype: int
        """
        ans = 0

        common_set = Uset(n)
        
        for t, u, v in edges:
            if t!=3: continue
            pu = common_set.find(u)
            pv = common_set.find(v)
            if pu==pv:
                ans += 1
            else:
                common_set.union(u, v)
        
        A_set = common_set.clone()
        B_set = common_set.clone()

        # print common_set.look_up
        
        
        for t, u, v in edges:
            if t!=1: continue
            pu = A_set.find(u)
            pv = A_set.find(v)
            if pu==pv:
                ans += 1
            else:
                A_set.union(pu, pv)

        # print A_set.look_up
        
        if A_set.connected() != 1: return -1

        for t, u, v in edges:
            if t!=2: continue
            pu = B_set.find(u)
            pv = B_set.find(v)
            if pu==pv:
                ans += 1
            else:
                B_set.union(pu, pv)
        
        if B_set.connected() != 1: return -1

        print 'ans:', ans
        return ans



if __name__ == '__main__':
    n = 4
    edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
    # n = 4
    # edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]

    # n = 4
    # edges = [[3,2,3],[1,1,2],[2,3,4]]
    su = Solution()
    su.maxNumEdgesToRemove(n, edges)

# https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/solution/bing-cha-ji-zheng-ming-zui-zhong-di-san-chong-lei-/







