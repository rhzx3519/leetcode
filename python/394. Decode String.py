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


class Solution(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        stack = []  # (str, int) 记录左括号之前的字符串和左括号外的上一个数字
        num = 0
        res = ""  # 实时记录当前可以提取出来的字符串
        for c in s:
            if c.isdigit():
                num = num * 10 + int(c)
            elif c == "[":
                stack.append((res, num))
                res, num = "", 0
            elif c == "]":
                top = stack.pop()
                res = top[0] + res * top[1]
            else:
                res += c
        return res
        

        

if __name__ == '__main__':
    s = "3[a]2[bc]"
    su = Solution()
    print su.decodeString(s)            
            
