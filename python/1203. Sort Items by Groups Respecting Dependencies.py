import collections

class Solution(object):
    def sortItems(self, n, m, group, beforeItems):
        """
        :type n: int
        :type m: int
        :type group: List[int]
        :type beforeItems: List[List[int]]
        :rtype: List[int]
        """
        item_ind = [0]*n
        item_adj = [set() for _ in range(n)]



        for i in range(n):
            if group[i] == -1:
                group[i] = m
                m += 1
            for bi in beforeItems[i]:
                item_adj[bi].add(i)

        for i in range(n):
            for ni in item_adj[i]:
                item_ind[ni] += 1

        group_ind = [0]*m
        group_adj = [set() for _ in range(m)]

        for i in range(n):
            for bi in beforeItems[i]:
                if group[bi] != group[i]:
                    group_adj[group[bi]].add(group[i])

        for i in range(m):
            for ni in group_adj[i]:
                group_ind[ni] += 1

        print item_adj, item_ind 
        print group_adj, group_ind
        print group


        item_topo = self.topo_sort(item_adj, item_ind)
        group_topo = self.topo_sort(group_adj, group_ind)

        print item_topo, map(lambda i: group[i], item_topo)
        print group_topo

        mem = collections.defaultdict(list)
        for i in item_topo:
            mem[group[i]].append(i)

        ans = []
        for i in group_topo:
            if i in mem:
                ans.extend(mem[i])

        # print ans
        return ans if len(ans)==n else []


    def topo_sort(self, adj, ind):
        n = len(ind)
        que = []
        for i in xrange(n):
            if ind[i] == 0:
                que.append(i)

        ans = []
        while que:
            i = que.pop(0)
            ans.append(i)
            for ni in adj[i]:
                ind[ni] -= 1
                if ind[ni] == 0:
                    que.append(ni)
        return ans







if __name__ == '__main__':
    # n = 8
    # m = 2
    # group = [-1,-1,1,0,0,1,0,-1]
    # beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]

    n = 8
    m = 2
    group = [-1,-1,1,0,0,1,0,-1]
    beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
    su = Solution()
    print su.sortItems(n, m, group, beforeItems)

