class Solution(object):
    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """
        files = path.split('/')
        st = []
        for x in files:
            if x == '.' or x == '':
                pass
            elif x == '..':
                if st:
                    st.pop()
            else:
                st.append(x)
        if not st:
            return '/'
        res = ''
        i = 0
        while i < len(st):
            res += '/' + st[i]
            i += 1
        
        return res
        