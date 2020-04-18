class Solution(object):
    def nextGreaterElement(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        st = []
        n1, n2 = len(nums1), len(nums2)

        nxt = [-1 for _ in range(n2)]
        for i, num in enumerate(nums2):
            while st and nums2[st[-1]] < num:
                nxt[st.pop()] = num
            st.append(i)
        
        res = []
        for i, num in enumerate(nums1):
            j = nums2.index(num)
            res.append(nxt[j])
        return res
        