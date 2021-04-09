class Solution(object):
    def getTriggerTime(self, increase, requirements):
        """
        :type increase: List[List[int]]
        :type requirements: List[List[int]]
        :rtype: List[int]
        """
        increase.insert(0, [0,0,0])
        for i in range(1, len(increase)):
            increase[i][0] += increase[i-1][0]
            increase[i][1] += increase[i-1][1]
            increase[i][2] += increase[i-1][2]
        
        N = len(increase)
        C, R, H = [x for x, _, _ in increase], [x for _, x, _ in increase], [x for _, _, x in increase]

        # print C, R, H
        ans = []
        for _, (c, r, h) in enumerate(requirements):
            i = bisect.bisect_left(C, c)
            j = bisect.bisect_left(R, r)
            k = bisect.bisect_left(H, h)
            # print i, j, k
            if i>=N or j>=N or k>=N:
                ans.append(-1)
            else:
                ans.append(max(i, j, k))

        return ans

