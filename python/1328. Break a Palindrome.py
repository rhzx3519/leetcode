class Solution(object):
    def breakPalindrome(self, palindrome):
        """
        :type palindrome: str
        :rtype: str
        """
        n = len(palindrome)
        if n <= 1: return ''
        for i in range(n/2):
            if palindrome[i] != 'a':
                return palindrome[:i] + 'a' + palindrome[i+1:]
        return palindrome[:-1] + 'b'

# 1. 字符串少于等于1时说明无法做到返回空串
# 2. 从头检查到中间，如果有不为a的字符，将该字符替换为a即可返回
# 3. 单数串的中间字符不需要处理，即使不为a也不能替换该字符，因为替换后还是回文串（把单数串的中间字符认为a即可）
# 4. 若检查到中间了都没返回，即都是a，则只需要将最后一个字符替换为b即可返回，第3种情况和第4种情况一样处理