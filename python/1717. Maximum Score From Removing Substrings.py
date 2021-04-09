class Solution(object):
    def maximumGain(self, s, x, y):
        """
        :type s: str
        :type x: int
        :type y: int
        :rtype: int
        """
        ans = [0]

        def count(s, target, score):
            st = []
            for i in range(len(s)):
                if st and st[-1]+s[i]==target:
                    st.pop()
                    ans[0] += score
                else:
                    st.append(s[i])
            return ''.join(st)

        if x>=y:
            s = count(s, 'ab', x)
            count(s, 'ba', y)
        else:
            s = count(s, 'ba', y)
            count(s, 'ab', x)

        return ans[0]



if __name__ == '__main__':
    s = "cdbcbbaaabab"
    x = 4
    y = 5
    # s = "aabbaaxybbaabb"
    # x = 5
    # y = 4
    su = Solution()
    print su.maximumGain(s, x, y)
