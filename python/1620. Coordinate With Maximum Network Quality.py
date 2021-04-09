class Solution(object):
    def bestCoordinate(self, towers, radius):
        """
        :type towers: List[List[int]]
        :type radius: int
        :rtype: List[int]
        """
        def dis(x1, y1, x2, y2):
            return math.sqrt((x1-x2)**2 + (y1-y2)**2)

        max_val = float('-inf')
        ans = None
        N = 51
        for i in range(N):
            for j in range(N):
                total = 0
                for x, y, q in towers:
                    d = dis(i, j, x, y)
                    if d > radius: continue
                    total += int(q / (1 + d))
                if total > max_val:
                    max_val = total
                    ans = [i, j]
        return ans