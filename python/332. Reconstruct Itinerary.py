#!usr/bin/env python  
#-*- coding:utf-8 _*-  

import collections

class Solution(object):
    def findItinerary(self, tickets):
        """
        # 思路：dfs，遍历的时候要删除之前走过的边
        # time：O(N!), space: O(1)     
        """
        edges = collections.defaultdict(list)
        for start, end in tickets:
            edges[start].append(end)
        
        for start in edges.iterkeys():
            edges[start].sort(reverse=True)
        
        path = []
        def dfs(start):
            while edges[start]:
                dfs(edges[start].pop())
            path.append(start)
        
        dfs('JFK')
        return path[::-1]



if __name__ == '__main__':
    # tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
    tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
    su = Solution()
    su.findItinerary2(tickets)


      


            
