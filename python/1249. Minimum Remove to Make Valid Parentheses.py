class Solution(object):
    def minRemoveToMakeValid(self, s):
        """
        :type s: str
        :rtype: str
        """
        st = []
        a = []
        for i, ch in enumerate(s):
            if ch not in ('(', ')'):
                a.append(i)
            elif ch == ')':
                if not st or s[st[-1]] != '(':
                    continue
                else:
                    st.pop()
                a.append(i)
            else:
                st.append(i)
                a.append(i)
        
        ans = []
        for i in a:
            if i in st:
                continue
            ans.append(s[i])
        return ''.join(ans)
            
