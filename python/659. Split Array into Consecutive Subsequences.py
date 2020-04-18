class Solution(object):
    def isPossible(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        d = collections.defaultdict(int)
        for i in nums:
            d[i] += 1

        for i in nums:
            subNum = 0
            next = i
            p = 1
            while d.get(next, 0) >= p:
                p = d.get(next)
                d[next] = p-1
                next += 1
                subNum += 1
            if subNum>0 and subNum<3:
                return False
        return True

            