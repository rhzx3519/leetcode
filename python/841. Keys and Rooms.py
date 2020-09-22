class Solution(object):
    def canVisitAllRooms(self, rooms):
        """
        :type rooms: List[List[int]]
        :rtype: bool
        """
        N = len(rooms)
        vis = [False] * N
        
        def dfs(idx):
            vis[idx] = True
            for nx in rooms[idx]:
                if not vis[nx]:
                    dfs(nx)
        dfs(0)
        return all(vis)