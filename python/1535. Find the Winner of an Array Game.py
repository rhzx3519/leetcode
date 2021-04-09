class Solution(object):
    def getWinner(self, arr, k):
        """
        :type arr: List[int]
        :type k: int
        :rtype: int
        """
        win = arr[0]
        count = 0
        n = len(arr)
        for i in range(1, n):
            if win > arr[i]:
                count += 1
            else:
                win = arr[i]
                count = 1
            if count >= k: break
        return win