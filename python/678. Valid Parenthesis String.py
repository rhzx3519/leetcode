class Solution(object):
    def checkValidString(self, s):
        """
        :type s: str
        :rtype: bool
        """
        st1 = []
        st2 = []
        for i in range(len(s)):
            ch = s[i]
            if ch=='(':
                st1.append(i)
            elif ch=='*':
                st2.append(i)
            else:
                if not st1 and not st2:
                    return False
                if st1:
                    st1.pop()
                else:
                    st2.pop()

        while st1 and st2:
            if st1[-1] > st2[-1]:
                return False
            st1.pop()
            st2.pop()
        return not st1 and not st2

if __name__ == '__main__':
    su = Solution()
    print su.checkValidString('(*)')                        