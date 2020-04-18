class Solution(object):
    def floodFill(self, image, sr, sc, newColor):
        """
        :type image: List[List[int]]
        :type sr: int
        :type sc: int
        :type newColor: int
        :rtype: List[List[int]]
        """
        m, n = len(image), len(image[0])

        def fill(x, y, originColor):
            if x<0 or x>=m or y<0 or y>=n: return
            
            if image[x][y]!=originColor: return
            image[x][y] = -1 # // choose：打标记，以免重复
            for dx, dy in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                fill(x+dx, y+dy, originColor)
            image[x][y] = newColor #  // unchoose：将标记替换为 newColor
        
        originColor = image[sr][sc]
        fill(sr, sc, originColor)
        return image

if __name__ == '__main__':
    image = [[0,0,0],[0,1,0]]
    sr = 1
    sc = 1
    newColor = 2
    su = Solution()
    su.floodFill(image, sr, sc, newColor)        