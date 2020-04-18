class FenwickTree(object):
    def __init__(self, n):
        super(FenwickTree, self).__init__()
        self.sums = [0]*(n+1)
    
    def update(self, i, k):
        while i < len(self.sums):
            self.sums[i] += k
            i += self.lowbit(i)

    def getSum(self, i):
        res = 0
        while i>0:
            res += self.sums[i]
            i -= self.lowbit(i)
        return res

    def lowbit(self, x):
        return x&(-x)

class Solution(object):
    def countRangeSum(self, nums, lower, upper):
        """
        :type nums: List[int]
        :type lower: int
        :type upper: int
        :rtype: int
        """
        n = len(nums)
        tree = FenwickTree(len(nums))
        for i in range(len(nums)):
            tree.update(i+1, nums[i])

        res = 0
        for i in range(n):
            for j in range(i, n):
                s = tree.getSum(j+1) - tree.getSum(i)
                if s>= lower and s<=upper:
                    res += 1
                    
        return res
        

if __name__ == '__main__':
    nums = [-2,5,-1]
    lower = -2
    upper = 2

    su = Solution()
    print su.countRangeSum(nums, lower, upper)