class Solution(object):
    def reverseParentheses(self, s):
        """
        :type s: str
        :rtype: str
        """
        st = []
        tmp = []
        for ch in s:
            if ch==')':
                t = st.pop() if st else []
                tmp = t + tmp[::-1]
                # print ''.join(tmp)
            elif ch=='(':
                st.append(tmp[:])
                tmp = []
            else:
                tmp.append(ch)
        # print st
        return ''.join(tmp)

# a(b(cd)e)
# "((eqk((h))))"