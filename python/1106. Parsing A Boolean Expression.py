class Solution(object):
    def parseBoolExpr(self, expression):
        """
        :type expression: str
        :rtype: bool
        """
        st1 = []
        st2 = []
        for ch in expression:
            if ch in ('!', '|', '&'):
                st1.append(ch)
            elif ch == ')':
                f = t = 0
                while st2 and st2[-1]!='(':
                    if st2[-1] == 'f':
                        f += 1
                    elif st2[-1] == 't':
                        t += 1
                    st2.pop()
                st2.pop()
                x = st1.pop()
                if x == '!':
                    if t > 0:
                        st2.append('f')
                    else:
                        st2.append('t')
                elif x == '|':
                    if t > 0:
                        st2.append('t')
                    else:
                        st2.append('f')
                elif x == '&':
                    if f > 0:
                        st2.append('f')
                    else:
                        st2.append('t')
            else:
                st2.append(ch)
        
        return st2[-1]=='t'