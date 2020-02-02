class Solution(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """

        def is_digit(ch):
            return ord(ch) >= ord('0') and ord(ch) <= ord('9')

        num_st = []
        word_st = []
        n = len(s)
        cur = ''
        num = 0
        for i in range(n):
            if is_digit(s[i]):
                num = num*10 + (ord(s[i]) - ord('0'))
            elif s[i]=='[':
                num_st.append(num)
                word_st.append(cur)
                num = 0
                cur = ''
            elif s[i]==']':
                k = num_st.pop()
                word = word_st.pop()
                for _ in range(k):
                    word += cur
                cur = word
            else: # letter
                cur += s[i]
        return cur
        

if __name__ == '__main__':
    s = "3[a]2[bc]"
    su = Solution()
    print su.decodeString(s)            
            
