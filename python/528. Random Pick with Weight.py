class Solution(object):
    
    def __init__(self, w):
        """
        :type w: List[int]
        """
        self.ws = []
        s = 0
        for x in w:
            s += x
            self.ws.append(s)
        self.s = s

    def pickIndex(self):
        """
        :rtype: int
        """
        w = random.randint(1, self.s)
        return bisect.bisect_left(self.ws, w)


# Your Solution object will be instantiated and called as such:
# obj = Solution(w)
# param_1 = obj.pickIndex()