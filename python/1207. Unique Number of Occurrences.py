class Solution(object):
    def uniqueOccurrences(self, arr):
        """
        :type arr: List[int]
        :rtype: bool
        """
        count = collections.Counter(arr)
        st = set()
        for v in count.itervalues():
            if v in st:
                return False
            st.add(v)
        return True