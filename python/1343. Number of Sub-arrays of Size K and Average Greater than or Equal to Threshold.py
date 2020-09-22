class Solution(object):
    def numOfSubarrays(self, arr, k, threshold):
        """
        :type arr: List[int]
        :type k: int
        :type threshold: int
        :rtype: int
        """
        ans = 0
        s = 0
        l = 0
        n = len(arr)
        for r in range(n):
            s += arr[r]
            if r >= k-1:
                # print r-k+1, r, s, s/k
                if s/k >= threshold:
                    ans += 1
                s -= arr[r-k+1]
        return ans