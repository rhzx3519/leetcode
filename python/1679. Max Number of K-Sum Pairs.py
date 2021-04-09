class Solution(object):
    def maxOperations(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        ans = 0
        count = collections.Counter(nums)
        keys = count.keys()
        for x in keys:
            y = k - x
            if y not in count:
                continue
            if x == y:
                t = count[x]
                ans += t//2
                if t&1: count[x] = 1
                else: del count[x]
            else:
                z = min(count[x], count[y])
                count[x] -= z
                count[y] -= z
                ans += z
                if count[x]==0: del count[x]
                if count[y]==0: del count[y]
        
        return ans