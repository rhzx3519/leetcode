import bisect

class Solution(object):
    def medianSlidingWindow(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[float]
        """
        n = len(nums)
        if k%2==0:
            n1 = k//2
            n2 = k//2 - 1
        else:
            n1 = n2 = k//2

        t = sorted(nums[:k])
        res = [(t[n1] + t[n2]) / 2.0]
        for i in range(1, n-k+1):
            t.pop(bisect.bisect_left(t, nums[i-1]))
            bisect.insort(t, nums[i+k-1])
            res.append((t[n1] + t[n2]) / 2.0)
        return res



if __name__ == '__main__':
    nums = [1,4,2,3]
    k = 4
    su = Solution()
    print su.medianSlidingWindow(nums, k)            