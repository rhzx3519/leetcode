class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        mp = {')':'(', ']':'[', '}':'{'}
        st = []
        for c in s:
            if c == '(' or c == '[' or c == '{':
                st.append(c)
                continue
            if not st or st[-1] != mp[c]:
                return False
            else:
                st.pop()
                
        return not st    