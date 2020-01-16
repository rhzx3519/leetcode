class Solution(object):
    
    def __init__(self, n_rows, n_cols):
        """
        :type n_rows: int
        :type n_cols: int
        """
        self.n_rows = n_rows
        self.n_cols = n_cols
        self.max_size = n_rows*n_cols
        self.mark = {}

    def flip(self):
        """
        :rtype: List[int]
        """
        i = 0
        j = 0
        while True:
            idx = random.randint(0, self.max_size-1)
            if self.mark.get(idx, 0) == 0:
                self.mark[idx] = 1
                i = idx / self.n_cols
                j = idx % self.n_cols
                break
            
        return [i, j]


    def reset(self):
        """
        :rtype: None
        """
        self.mark = {}

        


# Your Solution object will be instantiated and called as such:
# obj = Solution(n_rows, n_cols)
# param_1 = obj.flip()
# obj.reset()