class Solution(object):
    def compress(self, chars):
        """
        :type chars: List[str]
        :rtype: int
        """
        n = len(chars)
        l = 0
        pre = chars[0]
        count = 1
        for i in range(1, n+1):
            if i==n or chars[i] != pre:
                chars[l] = pre
                l += 1
                if count > 1:
                    for d in str(count):
                        chars[l] = d
                        l += 1
                count = 0
            if i < n:
                pre = chars[i]
            count += 1
        return l
