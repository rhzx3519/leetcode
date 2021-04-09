class Solution(object):
    def sortString(self, s):
        """
        :type s: str
        :rtype: str
        """
        ans = []
        s = list(s)
        while s:
            min_ch = None
            idx = -1
            for i, ch in enumerate(s):
                if not min_ch or min_ch > ch:
                    min_ch = ch
                    idx = i
            ans.append(s.pop(idx))
            if not s:
                break
            
            while s:
                min_ch = None
                idx = -1
                for i, ch in enumerate(s):
                    if ch > ans[-1] and (not min_ch or min_ch > ch):
                        min_ch = ch
                        idx = i
                if not min_ch:
                    break
                ans.append(s.pop(idx))

            if not s:
                break
            
            max_ch = None
            idx = -1
            for i, ch in enumerate(s):
                if not max_ch or max_ch < ch:
                    max_ch = ch
                    idx = i
            ans.append(s.pop(idx))
            if not s:
                break
            
            while s:
                max_ch = None
                idx = -1
                for i, ch in enumerate(s):
                    if ch < ans[-1] and (not max_ch or max_ch < ch):
                        max_ch = ch
                        idx = i
                if not max_ch:
                    break
                ans.append(s.pop(idx))
        return ''.join(ans)




