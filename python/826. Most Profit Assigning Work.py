class Solution(object):
    def maxProfitAssignment(self, difficulty, profit, worker):
        """
        :type difficulty: List[int]
        :type profit: List[int]
        :type worker: List[int]
        :rtype: int
        """
        N = len(profit)
        arr = zip(profit, difficulty)
        arr.sort(key=lambda x: -x[0])
        ans = 0
        j = 0
        for w in sorted(worker, reverse=True):
            while arr:
                if w < arr[0][1]:
                    arr.pop(0)
                else:
                    ans += arr[j][0]
                    break
        return ans