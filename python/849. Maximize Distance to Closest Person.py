class Solution(object):
    def maxDistToClosest(self, seats):
        """
        :type seats: List[int]
        :rtype: int
        """
        max_len = 0
        s = 0
        l = 0
        i = 0
        n = len(seats)
        while i < n and seats[i] != 1:
            i += 1
        l = i

        while i < n:
            if seats[i]==1:
                s = 0
            else:
                s += 1
            max_len = max(max_len, s)
            i += 1
        if max_len&1:
            max_len = max_len//2 + 1
        else:
            max_len = max_len//2
        return max(max_len, s, l)

        