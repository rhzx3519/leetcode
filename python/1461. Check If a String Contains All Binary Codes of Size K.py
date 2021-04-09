class Solution(object):
    def hasAllCodes(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: bool
        """
        st = set()
        n = len(s)
        for i in range(n-k+1):
            st.add(s[i:i+k])
        # print st
        return len(st)==2**k

