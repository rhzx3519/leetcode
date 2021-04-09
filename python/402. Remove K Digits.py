class Solution(object):
    def removeKdigits(self, num, k):
        """
        :type num: str
        :type k: int
        :rtype: str
        """
        st = []
        for ch in num:
            while k and st and st[-1] > ch:
                st.pop()
                k -= 1
            st.append(ch)
        
        st = st if k<=0 else st[:-k]
        return ''.join(st).lstrip('0') or '0'
        
# 思路：尽量删除逆序的数字

if __name__ == '__main__':
    num = "1432219"
    k = 3
    su = Solution()
    print su.removeKdigits(num, k)        