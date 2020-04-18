class Solution(object):
    def findContinuousSequence(self, target):
        """
        :type target: int
        :rtype: List[List[int]]
        """
        res = []
        a = []
        s = 0
        for i in range(1, target):
            while s > target:
                s -= a[0]
                del a[0]
            if s==target:
                res.append(a[:])
            s += i
            a.append(i)
        return res

        
if __name__ == '__main__':
    target = 9
    su = Solution()
    print su.findContinuousSequence(9)                    