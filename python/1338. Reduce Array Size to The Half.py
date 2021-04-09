class Solution(object):
    def minSetSize(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        n = len(arr)
        cnt = collections.Counter(arr)
        freq = sorted(cnt.values())
        ans = t = 0
        for i in range(len(freq)-1, -1, -1):
            t += freq[i]
            ans += 1
            if t >= n/2:
                break
        return ans