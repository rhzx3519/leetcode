class Solution(object):
    def entityParser(self, text):
        """
        :type text: str
        :rtype: str
        """
        mem = {'&quot;' : '"',
                '&apos;' : '\'',
                '&amp;' : '&',
                '&gt;' : '>',
                '&lt;' : '<',
                '&frasl;' : '/'}
        a = []
        i = 0
        while i < len(text):
            if text[i]=='&':
                j = i + 1
                for k, v in mem.iteritems():
                    lk = len(k)
                    if i+lk <= len(text) and text[i: i+lk]==k:
                        a.append(v)
                        j = i + lk
                        break
                if j==i+1:
                    a.append(text[i])
                i = j
                continue
            a.append(text[i])
            i += 1
        return ''.join(a)