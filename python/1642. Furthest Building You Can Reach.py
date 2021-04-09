class Solution(object):
    def furthestBuilding(self, heights, bricks, ladders):
        """
        :type heights: List[int]
        :type bricks: int
        :type ladders: int
        :rtype: int
        """
        b = bricks
        l = ladders

        que = []
        n = len(heights)
        for i in range(1, n):
            if heights[i] <= heights[i-1]: continue
            h = heights[i] - heights[i-1]
            if ladders > 0:
                ladders -= 1
             while que and que[-1] > h:
                    que.pop()
                que.append(h)
            else:   
                

        return n - 1



if __name__ == '__main__':
    heights = [1,5,1,2,3,4,10000]
    bricks = 4
    ladders = 0

    # heights = [4,12,2,7,3,18,20,3,19]
    # bricks = 10
    # ladders = 2

    su = Solution()
    su.furthestBuilding(heights, bricks, ladders)