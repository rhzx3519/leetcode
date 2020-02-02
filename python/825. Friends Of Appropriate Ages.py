class Solution(object):
    def numFriendRequests(self, ages):
        """
        :type ages: List[int]
        :rtype: int
        """
        a = [0]*121
        for i in ages:
            a[i] += 1

        res = 0
        for i in range(15, len(a)):
            res += a[i]*(a[i] - 1)
            for j in range(int(i/2+7) + 1, i):
                res += a[i]*a[j]
        return res
        