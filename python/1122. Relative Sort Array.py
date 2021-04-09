class Solution(object):
    def relativeSortArray(self, arr1, arr2):
        """
        :type arr1: List[int]
        :type arr2: List[int]
        :rtype: List[int]
        """
        d = {num: i for i, num in enumerate(arr2)}
        return sorted(arr1, key=lambda x: d.get(x, len(arr1) + x))