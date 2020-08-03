class Solution(object):
    def intersect(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        n1 = len(nums1)
        n2 = len(nums2)
        if n1 > n2:
            return self.intersect(nums2, nums1)
        nums1.sort()
        nums2.sort()
        i = j = 0
        res = []
        while i<n1 and j < n2:
            if nums1[i] == nums2[j]:
                res.append(nums1[i])
                i += 1
                j += 1
            elif nums1[i] < nums2[j]:
                i += 1
            else:
                j += 1
        return res

# 思路: 归并思想
# time: O(NlogN), space: O(1)
