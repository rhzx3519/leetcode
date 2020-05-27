#!usr/bin/env python  
#-*- coding:utf-8 _*-  

import collections

class Solution(object):
    def findItinerary(self, tickets):
        """
        :type tickets: List[List[str]]
        :rtype: List[str]
        """
        n = len(tickets)
        path = ['JFK']
        def dfs(port, path, tickets, idx):
            if idx==n+1: return path
            nxt = []
            for i, ticket in enumerate(tickets):
                if ticket[0]==port:
                    nxt.append(i)
            for i in nxt:
                t = tickets[i][1],
                dfs(t, path+[t], tickets[:i]+tickets[i+1:], idx+1)

        return dfs('JFK', [], tickets, 1)


if __name__ == '__main__':
    # tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
    tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
    su = Solution()
    print su.findItinerary(tickets)


        


            
