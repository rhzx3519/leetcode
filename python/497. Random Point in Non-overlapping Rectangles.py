from bisect import bisect_left
from random import randint

class Solution(object):

    def __init__(self, rects):
        """
        :type rects: List[List[int]]
        """
        self.rects = rects
        self.weight = []
        s = 0
        for rect in rects:
            s += (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
            self.weight.append(s)

    def pick(self):
        """
        :rtype: List[int]
        """
        index = bisect_left(self.weight, randint(1, self.weight[-1]))
        rect = self.rects[index]
        return [randint(rect[0], rect[2]), randint(rect[1], rect[3])]
        


# Your Solution object will be instantiated and called as such:
# obj = Solution(rects)
# param_1 = obj.pick()


## 根据面积和来随机