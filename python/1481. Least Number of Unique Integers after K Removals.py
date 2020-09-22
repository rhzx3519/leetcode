class Solution(object):
    def findLeastNumOfUniqueInts(self, arr, k):
        """
        :type arr: List[int]
        :type k: int
        :rtype: int
        """
        cnt = collections.Counter(arr)
        keys = cnt.keys()
        keys.sort(key=lambda x: cnt[x])
        
        n = len(keys)
        idx = 0
        for i, key in enumerate(keys):
            num = cnt[key]
            k -= num
            if k == 0:
                return n-i - 1
            elif k < 0:
                return n-i
        return 0

            

