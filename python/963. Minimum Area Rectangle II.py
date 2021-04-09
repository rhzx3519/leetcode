import collections

class Solution(object):
    def minAreaFreeRect(self, points):
        """
        :type points: List[List[int]]
        :rtype: float
        """
        mem = collections.defaultdict(list)
        for i in range(len(points)):
            for j in range(i+1, len(points)):
                center = ((points[i][0] - points[j][0])*0.5, 
                            (points[i][1] - points[j][1])*0.5)
                r = ((points[i][0] - points[j][0])**2 + (points[i][1] - points[j][1])**2) * 0.5
                mem[center, r].append(points[i])

        
        print mem.values()
    

if __name__ == '__main__':
    points = [[1,2],[2,1],[1,0],[0,1]]
    su = Solution()
    su.minAreaFreeRect(points)