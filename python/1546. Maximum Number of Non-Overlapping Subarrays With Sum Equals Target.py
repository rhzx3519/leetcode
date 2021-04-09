class Solution(object):
    def maxNonOverlapping(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        st = set([0])
        ans = 0
        pre = 0
        for num in nums:
            pre += num
            if (pre-target) in st: 
                pre = 0
                st = set([0])
                ans += 1
            else:
                st.add(pre)
        return ans

