class Solution(object):
    def maskPII(self, S):
        """
        :type S: str
        :rtype: str
        """

        def encodeEmail(s):
            s = s.lower()
            mid = s.index('@')
            name = s[:mid]
            return name[0] + '*' * 5 + name[-1] + s[mid:]

        def encodePhone(s):
            t = []
            for d in s:
                if d.isdigit():
                    t.append(d)
            s = ''.join(t)
            if len(s) == 10:
                return '***-***-' + s[-4:]
            else:
                return '+' + '*'*(len(s)-10) + '-***-***-' + s[-4:]

        if S[0].isalpha():
            return encodeEmail(S)
        else:
            return encodePhone(S)