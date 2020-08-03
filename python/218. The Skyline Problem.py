class Solution(object):
    def getSkyline(self, buildings):
        """
        :type buildings: List[List[int]]
        :rtype: List[List[int]]
        """
        
        a = []
        for b in buildings:
            a.append((b[0], -b[2]))
            a.append((b[1], b[2]))
        a.sort(key=lambda x: (x[0], x[1])) # 按照横坐标排序，相同情况下，纵坐标从小到大排序

        res = []
        que = [0]
        cur = prev = 0
        for x, y in a:
            if y < 0:
                que.append(-y)
            else:
                que.pop(que.index(y))
            cur = max(que)
            if cur != prev:
                res.append((x, cur))
                prev = cur
        return res

# 思路： 扫描线法, 使用扫描线，从左至右扫过。如果遇到左端点，将高度入堆，如果遇到右端点，则将高度从堆中删除。使用 cur 变量记录当前最大高度，prev记录上一个最大高度。
# time: O(N^2), space: O(N)
        