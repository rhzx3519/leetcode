class Solution(object):
    def checkArithmeticSubarrays(self, nums, l, r):
        """
        :type nums: List[int]
        :type l: List[int]
        :type r: List[int]
        :rtype: List[bool]
        """
        def foo(a):
            if len(a) <= 1: return False
            if len(a) == 2: return True
            a.sort()
            for i in range(1, len(a) - 1):
                if a[i] - a[i-1] != a[i+1] - a[i]:
                    return False
            return True


        n = len(nums)
        m = len(l)
        return [foo(nums[l1:r1+1]) for l1, r1 in zip(l, r)]
            